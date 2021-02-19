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

/// Defines the currency of a transfer.
type Currency byte

/// Represents balance value.
type Balance types.U64

/// Describes the reason for a transfer.
type Reason byte

/// Describes type of a subaccount.
type SubAccType byte

/// Represents a validator ID.
type ValidatorID types.U64

/// Represents asset type from the Financial Pallet.
type AssetType byte
