// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package core

import (
	"math/big"

	msg "github.com/ChainSafe/ChainBridge/message"
	metrics "github.com/ChainSafe/ChainBridge/metrics/types"
	"github.com/ChainSafe/ChainBridge/router"
)

type Chain interface {
	Start() error // Start chain
	SetRouter(*router.Router)
	Id() msg.ChainId
	Name() string
<<<<<<< HEAD
	GetLatestBlock() (*big.Int, error)
=======
	LatestBlock() metrics.LatestBlock
>>>>>>> origin/master
	Stop()
}

type ChainConfig struct {
	Name           string            // Human-readable chain name
	Id             msg.ChainId       // ChainID
	Endpoint       string            // url for rpc endpoint
	From           string            // address of key to use
	KeystorePath   string            // Location of key files
	Insecure       bool              // Indicated whether the test keyring should be used
	BlockstorePath string            // Location of blockstore
	FreshStart     bool              // If true, blockstore is ignored at start.
	LatestBlock    bool              // If true, overrides blockstore or latest block in config and starts from current block
	Opts           map[string]string // Per chain options
}
