// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	"context"
	"errors"
	"fmt"
	"github.com/ChainSafe/ChainBridge/shared/equilibrium"
	"github.com/ChainSafe/chainbridge-utils/core"
	"math/big"
	"time"

	"github.com/ChainSafe/ChainBridge/shared/equilibrium/metrics"
	"github.com/ChainSafe/chainbridge-utils/blockstore"
	metricTypes "github.com/ChainSafe/chainbridge-utils/metrics/types"
	"github.com/ChainSafe/chainbridge-utils/msg"
	eth "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"

	"github.com/ChainSafe/ChainBridge/bindings/Bridge"
	"github.com/ChainSafe/ChainBridge/bindings/ERC20Handler"
	"github.com/ChainSafe/ChainBridge/bindings/ERC721Handler"
	"github.com/ChainSafe/ChainBridge/bindings/GenericHandler"
	"github.com/ChainSafe/ChainBridge/chains"
	utils "github.com/ChainSafe/ChainBridge/shared/ethereum"
)

var BlockRetryInterval = time.Second * 5
var BlockRetryLimit = 5
var ErrFatalPolling = errors.New("listener block polling failed")

type listener struct {
	cfg                    Config
	conn                   Connection
	router                 chains.Router
	bridgeContract         *Bridge.Bridge // instance of bound bridge contract
	erc20HandlerContract   *ERC20Handler.ERC20Handler
	erc721HandlerContract  *ERC721Handler.ERC721Handler
	genericHandlerContract *GenericHandler.GenericHandler
	log                    equilibrium.TransferLogger
	blockstore             blockstore.Blockstorer
	stop                   <-chan int
	sysErr                 chan<- error // Reports fatal error to core
	latestBlock            metricTypes.LatestBlock
	metrics                *metrics.ChainMetrics
}

// NewListener creates and returns a listener
func NewListener(conn Connection, cfg *Config, log equilibrium.TransferLogger, bs blockstore.Blockstorer, stop <-chan int, sysErr chan<- error, m *metrics.ChainMetrics) *listener {
	return &listener{
		cfg:         *cfg,
		conn:        conn,
		log:         log,
		blockstore:  bs,
		stop:        stop,
		sysErr:      sysErr,
		latestBlock: metricTypes.LatestBlock{LastUpdated: time.Now()},
		metrics:     m,
	}
}

// setContracts sets the listener with the appropriate contracts
func (l *listener) setContracts(bridge *Bridge.Bridge, erc20Handler *ERC20Handler.ERC20Handler, erc721Handler *ERC721Handler.ERC721Handler, genericHandler *GenericHandler.GenericHandler) {
	l.bridgeContract = bridge
	l.erc20HandlerContract = erc20Handler
	l.erc721HandlerContract = erc721Handler
	l.genericHandlerContract = genericHandler
}

// sets the router
func (l *listener) setRouter(r chains.Router) {
	l.router = r
}

// start registers all subscriptions provided by the config
func (l *listener) start() error {
	l.log.Debug("Starting listener...")

	go func() {
		err := l.pollBlocks()
		if err != nil {
			l.log.Error("Polling blocks failed", "err", err)
		}
	}()

	return nil
}

// pollBlocks will poll for the latest block and proceed to parse the associated events as it sees new blocks.
// Polling begins at the block defined in `l.cfg.startBlock`. Failed attempts to fetch the latest block or parse
// a block will be retried up to BlockRetryLimit times before continuing to the next block.
func (l *listener) pollBlocks() error {
	l.log.Info(fmt.Sprintf("Polling Blocks for %s with interval %d and sync interval %d...", l.cfg.name, l.cfg.pollingInterval, l.cfg.syncPollingInterval), "chain_id", l.cfg.id)

	var blockDelay = l.cfg.blockDelay
	l.log.Info(fmt.Sprintf("BlockDelay is %d", blockDelay), "chain_id", l.cfg.id)

	var currentBlock = l.cfg.startBlock
	var retry = BlockRetryLimit
	var isSync = true

	for {
		select {
		case <-l.stop:
			return errors.New("polling terminated")
		default:

			var pollingInterval time.Duration
			if isSync {
				pollingInterval = l.cfg.syncPollingInterval
			} else {
				pollingInterval = l.cfg.pollingInterval
			}

			if pollingInterval != 0 {
				time.Sleep(pollingInterval)
			}

			// No more retries, goto next block
			if retry == 0 {
				l.log.Error("Polling failed, retries exceeded", "chain_id", l.cfg.id)
				l.sysErr <- ErrFatalPolling
				return nil
			}

			latestBlock, err := l.conn.LatestBlock()
			if err != nil {
				l.log.Error("Unable to get latest block", "block", currentBlock, "err", err, "chain_id", l.cfg.id)
				if l.metrics != nil {
					l.metrics.LatestBlockFailedRequests.Inc()
				}
				retry--
				time.Sleep(BlockRetryInterval)
				continue
			}

			if l.metrics != nil {
				l.metrics.LatestKnownBlock.Set(float64(latestBlock.Int64()))
				expectedCurrentBlock := big.NewInt(0).Sub(latestBlock, blockDelay)
				l.metrics.CurrentBlockLag.Set(float64(big.NewInt(0).Sub(expectedCurrentBlock, currentBlock).Int64()))
				l.metrics.LatestBlocksRequested.Inc()
			}

			// Sleep if the difference is less than BlockDelay; (latest - current) < BlockDelay
			if big.NewInt(0).Sub(latestBlock, currentBlock).Cmp(blockDelay) == -1 {
				l.log.Debug("Block not ready, will retry", "target", currentBlock, "latest", latestBlock, "chain_id", l.cfg.id)
				time.Sleep(BlockRetryInterval)
				if isSync {
					l.log.Info(fmt.Sprintf("Sync ended. Latest block: %d, current block: %d", latestBlock, currentBlock))
				}
				isSync = false
				continue
			}

			// Parse out events
			err = l.getDepositEventsForBlock(currentBlock)
			if err != nil {
				l.log.Error("Failed to get events for block", "block", currentBlock, "err", err, "chain_id", l.cfg.id)
				retry--
				continue
			}

			// Write to block store. Not a critical operation, no need to retry
			err = l.blockstore.StoreBlock(currentBlock)
			if err != nil {
				l.log.Error("Failed to write latest block to blockstore", "block", currentBlock, "err", err, "chain_id", l.cfg.id)
			}

			if l.metrics != nil {
				l.metrics.BlocksProcessed.Inc()
				l.metrics.LatestProcessedBlock.Set(float64(latestBlock.Int64()))
			}

			l.latestBlock.Height = big.NewInt(0).Set(latestBlock)
			l.latestBlock.LastUpdated = time.Now()

			// Goto next block and reset retry counter
			currentBlock.Add(currentBlock, big.NewInt(1))
			retry = BlockRetryLimit
		}
	}
}

// getDepositEventsForBlock looks for the deposit event in the latest block
func (l *listener) getDepositEventsForBlock(latestBlock *big.Int) error {
	l.log.Debug("Querying block for deposit events", "block", latestBlock)
	query := buildQuery(l.cfg.bridgeContract, utils.Deposit, latestBlock, latestBlock)

	// querying for logs
	logs, err := l.conn.Client().FilterLogs(context.Background(), query)
	if err != nil {
		return fmt.Errorf("unable to Filter Logs: %w", err)
	}

	// read through the log events and handle their deposit event if handler is recognized
	for _, log := range logs {
		var m msg.Message
		destId := msg.ChainId(log.Topics[1].Big().Uint64())
		rId := msg.ResourceIdFromSlice(log.Topics[2].Bytes())
		nonce := msg.Nonce(log.Topics[3].Big().Uint64())

		addr, err := l.bridgeContract.ResourceIDToHandlerAddress(&bind.CallOpts{From: l.conn.Keypair().CommonAddress()}, rId)
		if err != nil {
			return fmt.Errorf("failed to get handler from resource ID %x", rId)
		}

		messageContext := core.MessageContext{
			"ethereum_tx_hash": log.TxHash.Hex(),
		}

		if addr == l.cfg.erc20HandlerContract {
			m, err = l.handleErc20DepositedEvent(destId, nonce, messageContext)
		} else if addr == l.cfg.erc721HandlerContract {
			m, err = l.handleErc721DepositedEvent(destId, nonce)
		} else if addr == l.cfg.genericHandlerContract {
			m, err = l.handleGenericDepositedEvent(destId, nonce)
		} else {
			l.log.ErrorTransfer("event has unrecognized handler", messageContext, "handler", addr.Hex())
			return nil
		}

		if err != nil {
			return err
		}

		err = l.router.Send(m, messageContext)
		if err != nil {
			l.log.ErrorTransfer("subscription error: failed to route message", messageContext, "err", err)
		}
	}

	return nil
}

// buildQuery constructs a query for the bridgeContract by hashing sig to get the event topic
func buildQuery(contract ethcommon.Address, sig utils.EventSig, startBlock *big.Int, endBlock *big.Int) eth.FilterQuery {
	query := eth.FilterQuery{
		FromBlock: startBlock,
		ToBlock:   endBlock,
		Addresses: []ethcommon.Address{contract},
		Topics: [][]ethcommon.Hash{
			{sig.GetTopic()},
		},
	}
	return query
}
