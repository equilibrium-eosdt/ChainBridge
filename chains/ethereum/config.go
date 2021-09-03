// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"time"

	utils "github.com/ChainSafe/ChainBridge/shared/ethereum"
	"github.com/ChainSafe/chainbridge-utils/core"
	"github.com/ChainSafe/chainbridge-utils/msg"
	"github.com/ethereum/go-ethereum/common"
)

const DefaultGasLimit = 6721975
const DefaultGasPrice = 20000000000
const DefaultBlockDelay = 10

// Config encapsulates all necessary parameters in ethereum compatible forms
type Config struct {
	name                   string      // Human-readable chain name
	id                     msg.ChainId // ChainID
	endpoint               string      // url for rpc endpoint
	from                   string      // address of key to use
	keystorePath           string      // Location of keyfiles
	blockstorePath         string
	freshStart             bool // Disables loading from blockstore at start
	bridgeContract         common.Address
	erc20HandlerContract   common.Address
	erc721HandlerContract  common.Address
	genericHandlerContract common.Address
	gasLimit               *big.Int
	maxGasPrice            *big.Int
	http                   bool // Config for type of connection
	startBlock             *big.Int
	pollingInterval		   time.Duration
	syncPollingInterval    time.Duration
	blockDelay             *big.Int
}

// parseChainConfig uses a core.ChainConfig to construct a corresponding Config
func parseChainConfig(chainCfg *core.ChainConfig) (*Config, error) {

	config := &Config{
		name:                   chainCfg.Name,
		id:                     chainCfg.Id,
		endpoint:               chainCfg.Endpoint,
		from:                   chainCfg.From,
		keystorePath:           chainCfg.KeystorePath,
		blockstorePath:         chainCfg.BlockstorePath,
		freshStart:             chainCfg.FreshStart,
		bridgeContract:         utils.ZeroAddress,
		erc20HandlerContract:   utils.ZeroAddress,
		erc721HandlerContract:  utils.ZeroAddress,
		genericHandlerContract: utils.ZeroAddress,
		gasLimit:               big.NewInt(DefaultGasLimit),
		maxGasPrice:            big.NewInt(DefaultGasPrice),
		http:                   false,
		startBlock:             big.NewInt(0),
		pollingInterval:        time.Millisecond * 2000,
		syncPollingInterval:    time.Millisecond * 500,
		blockDelay:             big.NewInt(DefaultBlockDelay),
	}

	if contract, ok := chainCfg.Opts["bridge"]; ok && contract != "" {
		config.bridgeContract = common.HexToAddress(contract)
		delete(chainCfg.Opts, "bridge")
	} else {
		return nil, fmt.Errorf("must provide opts.bridge field for ethereum config")
	}

	config.erc20HandlerContract = common.HexToAddress(chainCfg.Opts["erc20Handler"])
	delete(chainCfg.Opts, "erc20Handler")

	config.erc721HandlerContract = common.HexToAddress(chainCfg.Opts["erc721Handler"])
	delete(chainCfg.Opts, "erc721Handler")

	config.genericHandlerContract = common.HexToAddress(chainCfg.Opts["genericHandler"])
	delete(chainCfg.Opts, "genericHandler")

	if gasPrice, ok := chainCfg.Opts["maxGasPrice"]; ok {
		price := big.NewInt(0)
		_, pass := price.SetString(gasPrice, 10)
		if pass {
			config.maxGasPrice = price
			delete(chainCfg.Opts, "maxGasPrice")
		} else {
			return nil, errors.New("unable to parse max gas price")
		}
	}

	if gasLimit, ok := chainCfg.Opts["gasLimit"]; ok {
		limit := big.NewInt(0)
		_, pass := limit.SetString(gasLimit, 10)
		if pass {
			config.gasLimit = limit
			delete(chainCfg.Opts, "gasLimit")
		} else {
			return nil, errors.New("unable to parse gas limit")
		}
	}

	if HTTP, ok := chainCfg.Opts["http"]; ok && HTTP == "true" {
		config.http = true
		delete(chainCfg.Opts, "http")
	} else if HTTP, ok := chainCfg.Opts["http"]; ok && HTTP == "false" {
		config.http = false
		delete(chainCfg.Opts, "http")
	}

	if startBlock, ok := chainCfg.Opts["startBlock"]; ok && startBlock != "" {
		block := big.NewInt(0)
		_, pass := block.SetString(startBlock, 10)
		if pass {
			config.startBlock = block
			delete(chainCfg.Opts, "startBlock")
		} else {
			return nil, errors.New("unable to parse start block")
		}
	}

	if pollingInterval, ok := chainCfg.Opts["pollingInterval"]; ok && pollingInterval != "" {
		interval, err := strconv.ParseInt(pollingInterval, 10, 64)
		if err == nil {
			config.pollingInterval = time.Millisecond * time.Duration(interval)
			delete(chainCfg.Opts, "pollingInterval")
		} else {
			return nil, errors.New("unable to parse polling interval")
		}
	}

	if syncPollingInterval, ok := chainCfg.Opts["syncPollingInterval"]; ok && syncPollingInterval != "" {
		interval, err := strconv.ParseInt(syncPollingInterval, 10, 64)
		if err == nil {
			config.syncPollingInterval = time.Millisecond * time.Duration(interval)
			delete(chainCfg.Opts, "syncPollingInterval")
		} else {
			return nil, errors.New("unable to parse sync polling interval")
		}
	}

	if blockDelay, ok := chainCfg.Opts["blockDelay"]; ok {
		delay := big.NewInt(0)
		_, pass := delay.SetString(blockDelay, 10)
		if pass {
			config.blockDelay = delay
			delete(chainCfg.Opts, "blockDelay")
		} else {
			return nil, errors.New("unable to parse block delay")
		}
	}

	if len(chainCfg.Opts) != 0 {
		return nil, fmt.Errorf("unknown Opts Encountered: %#v", chainCfg.Opts)
	}

	return config, nil
}
