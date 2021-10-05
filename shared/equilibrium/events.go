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
	Phase  types.Phase
	From   types.AccountID
	To     types.AccountID
	Asset  AssetType
	Value  Balance
	Reason Reason
	Topics []types.Hash
}

type EventBalancesDeleteAccount struct {
	Phase   types.Phase
	Account types.AccountID
	Topics  []types.Hash
}

type EventBalancesExchange struct {
	Phase    types.Phase
	Account1 types.AccountID
	Asset1   AssetType
	Amount1  Balance
	Account2 types.AccountID
	Asset2   AssetType
	Amount2  Balance
	Topics   []types.Hash
}

type EventBalancesDeposit struct {
	Phase 	types.Phase
	Account types.AccountID
	Asset 	AssetType
	Amount 	Balance
	Topics  []types.Hash
}

type EventBalancesWithdraw struct {
	Phase 	types.Phase
	Account types.AccountID
	Asset 	AssetType
	Amount 	Balance
	Topics  []types.Hash
}

type EventOracleNewPrice struct {
	Phase     types.Phase
	Asset     AssetType
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
	Topics    []types.Hash
}

type EventMarginCallExecuted struct {
	Phase     types.Phase
	AccountId types.AccountID
	Topics    []types.Hash
}

type EventNewAsset struct {
	Phase     types.Phase
	Asset     types.U64
	AssetName []types.U8
	Topics    []types.Hash
}

type EventDeleteAsset struct {
	Phase     types.Phase
	Asset     types.U64
	AssetName []types.U8
	Topics    []types.Hash
}

type EventUpdateAsset struct {
	Phase     types.Phase
	Asset     types.U64
	AssetName []types.U8
	Topics    []types.Hash
}

type EventMultisigSudoInitialized struct {
	Phase  types.Phase
	Topics []types.Hash
}

type EventMultisigSudoKeyAdded struct {
	Phase     types.Phase
	AccountId types.AccountID
	Topics    []types.Hash
}

type EventMultisigSudoKeyRemoved struct {
	Phase     types.Phase
	AccountId types.AccountID
	Topics    []types.Hash
}

type EventMultisigSudoThresholdModified struct {
	Phase     types.Phase
	Threshold types.U32
	Topics    []types.Hash
}

type EventMultisigSudoNewProposal struct {
	Phase     types.Phase
	AccountId types.AccountID
	CallHash  [32]types.U8
	Topics    []types.Hash
}

type EventMultisigSudoProposalCancelled struct {
	Phase    types.Phase
	CallHash [32]types.U8
	Topics   []types.Hash
}

type EventMultisigSudoProposalApproved struct {
	Phase     types.Phase
	AccountId types.AccountID
	CallHash  [32]types.U8
	Topics    []types.Hash
}

type EventMultisigSudoSudid struct {
	Phase    types.Phase
	CallHash [32]types.U8
	Topics   []types.Hash
}

type EventMultisigSudoFailed struct {
	Phase    types.Phase
	CallHash [32]types.U8
	Topics   []types.Hash
}

type EventEqBridgeRemark struct {
	Phase  types.Phase
	Hash   [32]types.Hash
	Topics []types.Hash
}

type EventCurveDistributionCurveAdminFeesDistributed struct {
	Phase     types.Phase
	Asset     AssetIdInnerType
	AssetName []types.U8
	Amount    Balance
	Topics    []types.Hash
}

type EventCurveAmmCreatePool struct {
	Phase     types.Phase
	AccountId types.AccountID
	PoolId    PoolId
	Topics    []types.Hash
}

type EventCurveAmmAddLiquidity struct {
	Phase       types.Phase
	Provider    types.AccountID
	PoolId      PoolId
	TokenAmount []Balance
	Fees        []Balance
	Invariant   Balance
	TokenSupply Balance
	MintedAmount Balance
	Topics      []types.Hash
}

type EventCurveAmmTokenExchange struct {
	Phase                       types.Phase
	Provider                    types.AccountID
	PoolId                      PoolId
	SendCoinIndex               PoolTokenIndex
	SendCoinExchangedAmount     Balance
	ReceivedCoinIndex           PoolTokenIndex
	ReceivedCoinExchangedAmount Balance
	Fee 						Balance
	Topics                      []types.Hash
}

type EventCurveAmmRemoveLiquidity struct {
	Phase        types.Phase
	Provider     types.AccountID
	PoolId       PoolId
	TokenAmounts []Balance
	Fees         []Balance
	TokenSupply  Balance
	Topics       []types.Hash
}

type EventCurveAmmRemoveLiquidityImbalance struct {
	Phase        types.Phase
	Provider     types.AccountID
	PoolId       PoolId
	TokenAmounts []Balance
	Fees         []Balance
	Invariant    Balance
	TokenSupply  Balance
	BurnAmount   Balance
	Topics       []types.Hash
}

type EventCurveAmmRemoveLiquidityOne struct {
	Phase               types.Phase
	Provider            types.AccountID
	PoolId              PoolId
	BurnAmount          Balance
	BurnedTokenIndex    PoolTokenIndex
	CoinExchangedAmount Balance
	NewTokenSupply      Balance
	Fee			        Balance
	Topics              []types.Hash
}

type EventCurveAmmWithdrawAdminFees struct {
	Phase     types.Phase
	Who       types.AccountID
	PoolId    PoolId
	AdminFees []Balance
	Topics    []types.Hash
}

type EventGensCrowdloanClaim struct {
	Phase   types.Phase
	Who     types.AccountID
	Amount  Balance
	Penalty Balance
	Topics  []types.Hash
}

type EventEqDexOrderCreated struct {
	Phase          types.Phase
	SubaccountId   types.AccountID
	OrderId        EqDexOrderId
	Asset          AssetType
	Amount         types.U128
	Price          types.I64
	Side           types.U8
	CreatedAt      types.U64
	ExpirationTime types.U64
	Topics         []types.Hash
}

type EventEqDexOrderDeleted struct {
	Phase     types.Phase
	AccountId types.OptionH256
	OrderId   EqDexOrderId
	Asset     AssetType
	Topics    []types.Hash
}

type EventEqDexMatch struct {
	Phase        types.Phase
	Asset        AssetType
	TakerRest    types.U128
	MakerPrice   types.I64
	MakerOrderId EqDexOrderId
	Maker        types.AccountID
	Taker        types.AccountID
	Topics       []types.Hash
}

type EventEqBridgeFromBridgeTransfer struct {
	Phase     types.Phase
	AccountId types.AccountID
	Asset     AssetType
	Amount    Balance
	Topics    []types.Hash
}

type EventEqBridgeToBridgeTransfer struct {
	Phase     types.Phase
	AccountId types.AccountID
	Asset     AssetType
	Amount    Balance
	Topics    []types.Hash
}

type ChainBridgeFeeChanged struct {
	Phase     types.Phase
	ChainId   types.U8
	Fee       Balance
	Topics    []types.Hash
}
