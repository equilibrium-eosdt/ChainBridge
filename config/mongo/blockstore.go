// Implementation of Blockstorer with MongoDB as backend.

package mongo

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ChainSafe/chainbridge-utils/msg"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/x/bsonx"

	mongodb "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	databaseName   = "chainbridge-relay"
	collectionName = "blockstore"
	timeout        = 10 * time.Second
)

// Blockstore implements Blockstorer.
type Blockstore struct {
	client *mongodb.Client
	id     string
}

// NewBlockstore creates new instance of Blockstore.
func NewBlockstore(url string, chain msg.ChainId, relayer string) (*Blockstore, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	client, err := mongodb.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}

	bs := Blockstore{
		client: client,
		id:     fmt.Sprintf("%s-%d", relayer, chain),
	}

	return &bs, nil
}

// StoreBlock writes the block number to disk.
func (b *Blockstore) StoreBlock(block *big.Int) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	coll := b.client.Database(databaseName).Collection(collectionName)

	filter := bson.D{{"relayer-id", bsonx.String(b.id)}}
	record := bson.D{{"relayer-id", bsonx.String(b.id)}, {"block", bsonx.String(block.String())}}
	var replacedDocument bson.M
	err := coll.FindOneAndReplace(ctx, filter, record).Decode(&replacedDocument)
	if err != nil {
		if err == mongodb.ErrNoDocuments {
			_, err = coll.InsertOne(ctx, record)
			return err
		}
		return err
	}

	return nil
}

// TryLoadLatestBlock will attempt to load the latest block for the chain/relayer pair,
// returning 0 if not found.
func (b *Blockstore) TryLoadLatestBlock() (*big.Int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	coll := b.client.Database(databaseName).Collection(collectionName)

	filter := bson.D{{"relayer-id", b.id}}
	var result bson.M
	err := coll.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if err == mongodb.ErrNoDocuments {
			return big.NewInt(0), nil
		} else {
			return nil, err
		}
	}

	s := result["block"].(string)
	block, ok := big.NewInt(0).SetString(s, 10)
	if !ok {
		return nil, fmt.Errorf("malformed integer %v", s)
	}
	return block, nil
}
