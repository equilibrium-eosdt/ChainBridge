package equilibrium

import (
	"github.com/centrifuge/go-substrate-rpc-client/types"
)

// AccountInfo contains information of an account.
// This is simplified version of type AccountInfo
// from go-substrate-rpc-client.
// Equilibrium does not use extra information from original
// AccountInfo which caused deserialization error while
// querying the storage.
type AccountInfo struct {
	Nonce    types.U32
	Refcount types.U32
}
