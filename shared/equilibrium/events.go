package equilibrium

import (
	"github.com/centrifuge/go-substrate-rpc-client/types"
	ethereum "github.com/ethereum/go-ethereum/common"
)

type EventLockdropUnlock struct {
	Phase  types.Phase
	Who    types.AccountID
	Amount Balance
	Topics []types.Hash
}

type EventLockdropLock struct {
	Phase  types.Phase
	Who    types.AccountID
	Amount Balance
	Topics []types.Hash
}

type EventBailsmanUnregisteredBailsman struct {
	Phase  types.Phase
	Who    types.AccountID
	Topics []types.Hash
}

// EventBalancesTransfer is emitted when a Substrate client calls Currency::transfer.
type EventBalancesTransfer struct {
	Phase    types.Phase
	From     types.AccountID
	To       types.AccountID
	Currency Currency
	Value    Balance
	Reason   Reason
	Topics   []types.Hash
}

type EventBalancesDeleteAccount struct {
	Phase   types.Phase
	Account types.AccountID
	Topics  []types.Hash
}

type EventOracleNewPrice struct {
	Phase     types.Phase
	Currency  Currency
	MedPrice  Balance
	NewPrice  Balance
	AccountId types.AccountID
	Topics    []types.Hash
}

type EventSubaccountsSubaccountCreated struct {
	Phase      types.Phase
	AccountId1 types.AccountID
	AccountId2 types.AccountID
	SubAccType SubAccType
	Topics     []types.Hash
}

type EventSubaccountsRegisterBailsman struct {
	Phase      types.Phase
	AccountId1 types.AccountID
	AccountId2 types.AccountID
	Topics     []types.Hash
}

type EventVestingUpdated struct {
	Phase     types.Phase
	AccountId types.AccountID
	Balance   Balance
	Topics    []types.Hash
}

type EventVestingCompleted struct {
	Phase     types.Phase
	AccountId types.AccountID
	Topics    []types.Hash
}

type EventClaimed struct {
	Phase           types.Phase
	AccountId       types.AccountID
	EthereumAddress ethereum.Address
	Balance         Balance
	Topics          []types.Hash
}

type EventValidatorAdded struct {
	Phase     types.Phase
	Validator types.AccountID
	Topics    []types.Hash
}

type EventValidatorRemoved struct {
	Phase     types.Phase
	Validator types.AccountID
	Topics    []types.Hash
}

type EventAddedToWhitelist struct {
	Phase     types.Phase
	AccountId types.AccountID
	Topics    []types.Hash
}

type EventRemovedFromWhitelist struct {
	Phase     types.Phase
	AccountId types.AccountID
	Topics    []types.Hash
}

type EventMetricsRecalculated struct {
	Phase  types.Phase
	Topics []types.Hash
}

type EventAssetMetricsRecalculated struct {
	Phase  types.Phase
	Asset  AssetType
	Topics []types.Hash
}

type EventPortfolioMetricsRecalculated struct {
	Phase     types.Phase
	AccountId types.AccountID
	Topics    []types.Hash
}

type EventMaintenanceMarginCall struct {
	Phase     types.Phase
	AccountId types.AccountID
	Timestamp types.U64
}

type EventMarginCallExecuted struct {
	Phase     types.Phase
	AccountId types.AccountID
}
