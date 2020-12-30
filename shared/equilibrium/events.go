package equilibrium

import (
	"github.com/centrifuge/go-substrate-rpc-client/types"
)

/// Defines the currency of a transfer.
type Currency byte
const (
	Unknown Currency = 0
	Usd     Currency = 1
	Eq      Currency = 2
	Eth     Currency = 3
	Btc     Currency = 4
	Eos     Currency = 5
)

/// Describes the reason for a transfer.
type Reason byte
const (
	Common                 Reason = 0
	InterestFee            Reason = 1
	MarginCall             Reason = 2
	BailsmenRedistribution Reason = 3
	TreasuryEqBuyout       Reason = 4
	TreasuryBuyEq          Reason = 5
)

// EventBalancesTransfer is emitted when a Substrate client calls Currency::transfer.
type EventBalancesTransfer struct {
	Phase    types.Phase
	From     types.AccountID
	To       types.AccountID
	Currency Currency
	Value    types.U64
	Reason   Reason
	Topics   []types.Hash
}

type EventOracleNewPrice struct {
	Phase     types.Phase
	Currency  Currency
	MedPrice  types.U64
	NewPrice  types.U64
	AccountId types.AccountID
	Topics    []types.Hash
}

type SubAccType byte

type EventSubaccountsSubaccountCreated struct {
	Phase      types.Phase
	AccountId1 types.AccountID
	AccountId2 types.AccountID
	SubAccType SubAccType
	Topics     []types.Hash
}

type EventSubaccountsRegisterBailsman struct {
	Phase     types.Phase
	AccountId1 types.AccountID
	AccountId2 types.AccountID
	Topics    []types.Hash
}
