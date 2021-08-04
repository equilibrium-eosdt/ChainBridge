// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only
/*
The ethereum package contains the logic for interacting with ethereum chains.

There are 3 major components: the connection, the listener, and the writer.
The currently supported transfer types are Fungible (ERC20), Non-Fungible (ERC721), and generic.

Connection

The connection contains the ethereum RPC client and can be accessed by both the writer and listener.

Listener

The listener polls for each new block and looks for deposit events in the bridge contract. If a deposit occurs, the listener will fetch additional information from the handler before constructing a message and forwarding it to the router.

Writer

The writer recieves the message and creates a proposals on-chain. Once a proposal is made, the writer then watches for a finalization event and will attempt to execute the proposal if a matching event occurs. The writer skips over any proposals it has already seen.
*/
package ethereum

import (
	"fmt"
	"github.com/ChainSafe/ChainBridge/shared/equilibrium"
	"math/big"
	"strings"

	bridge "github.com/ChainSafe/ChainBridge/bindings/Bridge"
	erc20Handler "github.com/ChainSafe/ChainBridge/bindings/ERC20Handler"
	erc721Handler "github.com/ChainSafe/ChainBridge/bindings/ERC721Handler"
	"github.com/ChainSafe/ChainBridge/bindings/GenericHandler"
	"github.com/ChainSafe/ChainBridge/config"
	"github.com/ChainSafe/ChainBridge/config/mongo"
	connection "github.com/ChainSafe/ChainBridge/connections/ethereum"
	"github.com/ChainSafe/ChainBridge/shared/equilibrium/metrics"
	"github.com/ChainSafe/chainbridge-utils/blockstore"
	"github.com/ChainSafe/chainbridge-utils/core"
	"github.com/ChainSafe/chainbridge-utils/crypto/secp256k1"
	"github.com/ChainSafe/chainbridge-utils/keystore"
	metricTypes "github.com/ChainSafe/chainbridge-utils/metrics/types"
	"github.com/ChainSafe/chainbridge-utils/msg"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var _ core.Chain = &Chain{}

var _ Connection = &connection.Connection{}

type Connection interface {
	Connect() error
	Keypair() *secp256k1.Keypair
	Opts() *bind.TransactOpts
	CallOpts() *bind.CallOpts
	LockAndUpdateOpts() error
	UnlockOpts()
	Client() *ethclient.Client
	EnsureHasBytecode(address common.Address) error
	LatestBlock() (*big.Int, error)
	WaitForBlock(block *big.Int) error
	Close()
}

type Chain struct {
	cfg      *core.ChainConfig // The config of the chain
	conn     Connection        // THe chains connection
	listener *listener         // The listener of this chain
	writer   *writer           // The writer of the chain
	stop     chan<- int
}

func isHttp(endpoint string) (bool, error) {
	if len(endpoint) == 0 {
		return false, fmt.Errorf("endpoint string is empty")
	}

	firstLetter := endpoint[0]

	if firstLetter != 'h' && firstLetter != 'w' {
		return false, fmt.Errorf("unknown protocol: %s", endpoint)
	}

	return firstLetter == 'h', nil
}

func InitializeChain(chainCfg *core.ChainConfig, logger equilibrium.TransferLogger, sysErr chan<- error, m *metrics.ChainMetrics) (*Chain, error) {
	cfg, err := parseChainConfig(chainCfg)
	if err != nil {
		return nil, err
	}

	kpI, err := keystore.KeypairFromAddress(cfg.from, keystore.EthChain, cfg.keystorePath, chainCfg.Insecure)
	if err != nil {
		return nil, err
	}
	kp, _ := kpI.(*secp256k1.Keypair)

	bs, err := setupBlockstore(cfg, kp)
	if err != nil {
		return nil, err
	}

	stop := make(chan int)

	http, err := isHttp(cfg.endpoint)
	if err != nil {
		return nil, err
	}

	conn := connection.NewConnection(cfg.endpoint, http, kp, logger, cfg.gasLimit, cfg.maxGasPrice)
	err = conn.Connect()
	if err != nil {
		return nil, err
	}
	err = conn.EnsureHasBytecode(cfg.bridgeContract)
	if err != nil {
		return nil, err
	}
	err = conn.EnsureHasBytecode(cfg.erc20HandlerContract)
	if err != nil {
		return nil, err
	}
	err = conn.EnsureHasBytecode(cfg.genericHandlerContract)
	if err != nil {
		return nil, err
	}

	bridgeContract, err := bridge.NewBridge(cfg.bridgeContract, conn.Client())
	if err != nil {
		return nil, err
	}

	chainId, err := bridgeContract.ChainID(conn.CallOpts())
	if err != nil {
		return nil, err
	}

	if chainId != uint8(chainCfg.Id) {
		return nil, fmt.Errorf("chainId (%d) and configuration chainId (%d) do not match", chainId, chainCfg.Id)
	}

	erc20HandlerContract, err := erc20Handler.NewERC20Handler(cfg.erc20HandlerContract, conn.Client())
	if err != nil {
		return nil, err
	}

	erc721HandlerContract, err := erc721Handler.NewERC721Handler(cfg.erc721HandlerContract, conn.Client())
	if err != nil {
		return nil, err
	}

	genericHandlerContract, err := GenericHandler.NewGenericHandler(cfg.genericHandlerContract, conn.Client())
	if err != nil {
		return nil, err
	}

	if chainCfg.LatestBlock {
		curr, err := conn.LatestBlock()
		if err != nil {
			return nil, err
		}
		cfg.startBlock = curr
	}

	listener := NewListener(conn, cfg, logger, bs, stop, sysErr, m)
	listener.setContracts(bridgeContract, erc20HandlerContract, erc721HandlerContract, genericHandlerContract)

	writer := NewWriter(conn, cfg, logger, stop, sysErr, m)
	writer.setContracts(bridgeContract, erc20HandlerContract)

	return &Chain{
		cfg:      chainCfg,
		conn:     conn,
		writer:   writer,
		listener: listener,
		stop:     stop,
	}, nil
}

func (c *Chain) SetRouter(r *core.Router) {
	r.Listen(c.cfg.Id, c.writer)
	c.listener.setRouter(r)
}

func (c *Chain) Start() error {
	err := c.listener.start()
	if err != nil {
		return err
	}

	err = c.writer.start()
	if err != nil {
		return err
	}

	c.writer.log.Debug("Successfully started chain")
	return nil
}

func (c *Chain) Id() msg.ChainId {
	return c.cfg.Id
}

func (c *Chain) Name() string {
	return c.cfg.Name
}

func (c *Chain) LatestBlock() metricTypes.LatestBlock {
	return c.listener.latestBlock
}

// Stop signals to any running routines to exit
func (c *Chain) Stop() {
	close(c.stop)
	if c.conn != nil {
		c.conn.Close()
	}
}

// setupBlockstore creates and inits new blockstore.
func setupBlockstore(cfg *Config, kp *secp256k1.Keypair) (*config.Blockstore, error) {
	var err error
	var dirbs *blockstore.Blockstore
	var mongobs *mongo.Blockstore
	if strings.HasPrefix(cfg.blockstorePath, config.MONGODB) {
		mongobs, err = setupMongoBlockstore(cfg, kp)
		if err != nil {
			return nil, err
		}
	} else {
		dirbs, err = setupDirBlockstore(cfg, kp)
		if err != nil {
			return nil, err
		}
	}
	return config.InitBlockstore(dirbs, mongobs), nil
}

// setupDirBlockstore creates new blockstore with directory.
func setupDirBlockstore(cfg *Config, kp *secp256k1.Keypair) (*blockstore.Blockstore, error) {
	bs, err := blockstore.NewBlockstore(cfg.blockstorePath, cfg.id, kp.Address())
	if err != nil {
		return nil, err
	}

	if !cfg.freshStart {
		latestBlock, err := bs.TryLoadLatestBlock()
		if err != nil {
			return nil, err
		}

		if latestBlock.Cmp(cfg.startBlock) == 1 {
			cfg.startBlock = latestBlock
		}
	}

	return bs, nil
}

// setupMongoBlockstore creates new blockstore with MongoDB.
func setupMongoBlockstore(cfg *Config, kp *secp256k1.Keypair) (*mongo.Blockstore, error) {
	bs, err := mongo.NewBlockstore(cfg.blockstorePath, cfg.id, kp.Address())
	if err != nil {
		return nil, err
	}

	if !cfg.freshStart {
		latestBlock, err := bs.TryLoadLatestBlock()
		if err != nil {
			return nil, err
		}

		if latestBlock.Cmp(cfg.startBlock) == 1 {
			cfg.startBlock = latestBlock
		}
	}

	return bs, nil
}
