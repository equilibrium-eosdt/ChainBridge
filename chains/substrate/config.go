// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package substrate

import (
	"math/big"
	"strconv"

	"github.com/ChainSafe/chainbridge-utils/core"
)

func parseStartBlock(cfg *core.ChainConfig) uint64 {
	if blk, ok := cfg.Opts["startBlock"]; ok {
		res, err := strconv.ParseUint(blk, 10, 32)
		if err != nil {
			panic(err)
		}
		return res
	}
	return 0
}

const DefaultBlockDelay = 10

func parseBlockDelay(cfg *core.ChainConfig) *big.Int {
	if blockDelay, ok := cfg.Opts["blockDelay"]; ok {
		delay := big.NewInt(0)
		_, pass := delay.SetString(blockDelay, 10)
		if pass {
			return delay
		} else {
			panic("unable to parse block delay")
		}
	}
	return big.NewInt(DefaultBlockDelay)
}
