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

/// Represents balance value.
type Balance types.U64

/// Describes the reason for a transfer.
type Reason byte

/// Describes type of a subaccount.
type SubAccType byte

/// Represents asset type.
type AssetType types.U64

/// Represents Pool Id type.
type PoolId types.U32

/// Represents Pool Token Index type.
type PoolTokenIndex types.U32

/// Represents Asset Id Inner type.
type AssetIdInnerType types.U64

type EqDexOrderId types.U64

type OptionAccountId struct {
	HasValue bool
	AccountId types.AccountID
}