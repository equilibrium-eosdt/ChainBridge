// Wrapper to support two blockstores: file and mongodb.

package config

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/ChainSafe/chainbridge-utils/blockstore"
	"github.com/ChainSafe/chainbridge-utils/msg"

	"github.com/ChainSafe/ChainBridge/config/mongo"
)

const MONGODB = "mongodb://"

// Blockstore implements Blockstorer.
type Blockstore struct {
	dir   *blockstore.Blockstore
	mongo *mongo.Blockstore
}

// InitBlockstore creates new blockstore.
func InitBlockstore(dirbs *blockstore.Blockstore, mongobs *mongo.Blockstore) *Blockstore {
	return &Blockstore{dirbs, mongobs}
}

// NewBlockstore creates new blockstore.
func NewBlockstore(url string, chain msg.ChainId, relayer string) (*Blockstore, error) {
	var err error
	var dirbs *blockstore.Blockstore
	var mongobs *mongo.Blockstore
	if strings.HasPrefix(url, MONGODB) {
		mongobs, err = mongo.NewBlockstore(url, chain, relayer)
		if err != nil {
			return nil, err
		}
	} else {
		dirbs, err = blockstore.NewBlockstore(url, chain, relayer)
		if err != nil {
			return nil, err
		}
	}
	return &Blockstore{dirbs, mongobs}, nil
}

// CheckBlockstore queries blockstore for the latest known block. If the latest block is
// greater than startBlock, then the latest block is returned, otherwise startBlock is.
func CheckBlockstore(b *Blockstore, startBlock uint64) (uint64, error) {
	if b.dir != nil {
		return checkDirBlockstore(b.dir, startBlock)
	}
	if b.mongo != nil {
		return checkMongoBlockstore(b.mongo, startBlock)
	}
	return 0, fmt.Errorf("no blockstore")
}

// StoreBlock writes the block number to disk.
func (b *Blockstore) StoreBlock(block *big.Int) error {
	if b.dir != nil {
		return b.dir.StoreBlock(block)
	}
	if b.mongo != nil {
		return b.mongo.StoreBlock(block)
	}
	return fmt.Errorf("no blockstore")
}

// checkDirBlockstore queries blockstore for the latest known block. If the latest block is
// greater than startBlock, then the latest block is returned, otherwise startBlock is.
func checkDirBlockstore(bs *blockstore.Blockstore, startBlock uint64) (uint64, error) {
	latestBlock, err := bs.TryLoadLatestBlock()
	if err != nil {
		return 0, err
	}

	if latestBlock.Uint64() > startBlock {
		return latestBlock.Uint64(), nil
	} else {
		return startBlock, nil
	}
}

// checkMongoBlockstore queries blockstore for the latest known block. If the latest block is
// greater than startBlock, then the latest block is returned, otherwise startBlock is.
func checkMongoBlockstore(bs *mongo.Blockstore, startBlock uint64) (uint64, error) {
	latestBlock, err := bs.TryLoadLatestBlock()
	if err != nil {
		return 0, err
	}

	if latestBlock.Uint64() > startBlock {
		return latestBlock.Uint64(), nil
	} else {
		return startBlock, nil
	}
}
