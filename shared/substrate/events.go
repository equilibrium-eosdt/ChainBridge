// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package utils

import (
	events "github.com/ChainSafe/chainbridge-substrate-events"
	"github.com/centrifuge/go-substrate-rpc-client/types"

	"github.com/ChainSafe/ChainBridge/shared/equilibrium"
)

type EventErc721Minted struct {
	Phase   types.Phase
	Owner   types.AccountID
	TokenId types.U256
	Topics  []types.Hash
}

type EventErc721Transferred struct {
	Phase   types.Phase
	From    types.AccountID
	To      types.AccountID
	TokenId types.U256
	Topics  []types.Hash
}

type EventErc721Burned struct {
	Phase   types.Phase
	TokenId types.AccountID
	Topics  []types.Hash
}

type EventExampleRemark struct {
	Phase  types.Phase
	Hash   types.Hash
	Topics []types.Hash
}

// EventNFTDeposited is emitted when NFT is ready to be deposited to other chain.
type EventNFTDeposited struct {
	Phase  types.Phase
	Asset  types.Hash
	Topics []types.Hash
}

// EventFeeChanged is emitted when a fee for a given key is changed.
type EventFeeChanged struct {
	Phase    types.Phase
	Key      types.Hash
	NewPrice types.U128
	Topics   []types.Hash
}

// EventNewMultiAccount is emitted when a multi account has been created.
// First param is the account that created it, second is the multisig account.
type EventNewMultiAccount struct {
	Phase   types.Phase
	Who, ID types.AccountID
	Topics  []types.Hash
}

// EventMultiAccountUpdated is emitted when a multi account has been updated. First param is the multisig account.
type EventMultiAccountUpdated struct {
	Phase  types.Phase
	Who    types.AccountID
	Topics []types.Hash
}

// EventMultiAccountRemoved is emitted when a multi account has been removed. First param is the multisig account.
type EventMultiAccountRemoved struct {
	Phase  types.Phase
	Who    types.AccountID
	Topics []types.Hash
}

// EventNewMultisig is emitted when a new multisig operation has begun.
// First param is the account that is approving, second is the multisig account.
type EventNewMultisig struct {
	Phase   types.Phase
	Who, ID types.AccountID
	Topics  []types.Hash
}

// TimePoint contains height and index
type TimePoint struct {
	Height types.U32
	Index  types.U32
}

// EventMultisigApproval is emitted when a multisig operation has been approved by someone.
// First param is the account that is approving, third is the multisig account.
type EventMultisigApproval struct {
	Phase     types.Phase
	Who       types.AccountID
	TimePoint TimePoint
	ID        types.AccountID
	Topics    []types.Hash
}

// EventMultisigExecuted is emitted when a multisig operation has been executed by someone.
// First param is the account that is approving, third is the multisig account.
type EventMultisigExecuted struct {
	Phase     types.Phase
	Who       types.AccountID
	TimePoint TimePoint
	ID        types.AccountID
	Result    types.DispatchResult
	Topics    []types.Hash
}

// EventMultisigCancelled is emitted when a multisig operation has been cancelled by someone.
// First param is the account that is approving, third is the multisig account.
type EventMultisigCancelled struct {
	Phase     types.Phase
	Who       types.AccountID
	TimePoint TimePoint
	ID        types.AccountID
	Topics    []types.Hash
}

type EventTreasuryMinting struct {
	Phase  types.Phase
	Who    types.AccountID
	Topics []types.Hash
}

type Events struct {
	types.EventRecords
	events.Events
	Erc721_Minted                                  []EventErc721Minted                                              // nolint:stylecheck,golint
	Erc721_Transferred                             []EventErc721Transferred                                         // nolint:stylecheck,golint
	Erc721_Burned                                  []EventErc721Burned                                              // nolint:stylecheck,golint
	Example_Remark                                 []EventExampleRemark                                             // nolint:stylecheck,golint
	Nfts_DepositAsset                              []EventNFTDeposited                                              // nolint:stylecheck,golint
	Council_Proposed                               []types.EventCollectiveProposed                                  // nolint:stylecheck,golint
	Council_Voted                                  []types.EventCollectiveVoted                                     // nolint:stylecheck,golint
	Council_Approved                               []types.EventCollectiveApproved                                  // nolint:stylecheck,golint
	Council_Disapproved                            []types.EventCollectiveDisapproved                               // nolint:stylecheck,golint
	Council_Executed                               []types.EventCollectiveExecuted                                  // nolint:stylecheck,golint
	Council_MemberExecuted                         []types.EventCollectiveMemberExecuted                            // nolint:stylecheck,golint
	Council_Closed                                 []types.EventCollectiveClosed                                    // nolint:stylecheck,golint
	Fees_FeeChanged                                []EventFeeChanged                                                // nolint:stylecheck,golint
	MultiAccount_NewMultiAccount                   []EventNewMultiAccount                                           // nolint:stylecheck,golint
	MultiAccount_MultiAccountUpdated               []EventMultiAccountUpdated                                       // nolint:stylecheck,golint
	MultiAccount_MultiAccountRemoved               []EventMultiAccountRemoved                                       // nolint:stylecheck,golint
	MultiAccount_NewMultisig                       []EventNewMultisig                                               // nolint:stylecheck,golint
	MultiAccount_MultisigApproval                  []EventMultisigApproval                                          // nolint:stylecheck,golint
	MultiAccount_MultisigExecuted                  []EventMultisigExecuted                                          // nolint:stylecheck,golint
	MultiAccount_MultisigCancelled                 []EventMultisigCancelled                                         // nolint:stylecheck,golint
	TreasuryReward_TreasuryMinting                 []EventTreasuryMinting                                           // nolint:stylecheck,golint
	Bailsman_UnregisteredBailsman                  []equilibrium.EventBailsmanUnregisteredBailsman                  // nolint:stylecheck,golint
	EqBalances_Transfer                            []equilibrium.EventBalancesTransfer                              // nolint:stylecheck,golint
	EqBalances_DeleteAccount                       []equilibrium.EventBalancesDeleteAccount                         // nolint:stylecheck,golint
	EqBalances_Exchange                            []equilibrium.EventBalancesExchange                              // nolint:stylecheck,golint
	Oracle_NewPrice                                []equilibrium.EventOracleNewPrice                                // nolint:stylecheck,golint
	Subaccounts_SubaccountCreated                  []equilibrium.EventSubaccountsSubaccountCreated                  // nolint:stylecheck,golint
	Subaccounts_RegisterBailsman                   []equilibrium.EventSubaccountsRegisterBailsman                   // nolint:stylecheck,golint
	Vesting_VestingUpdated                         []equilibrium.EventVestingUpdated                                // nolint:stylecheck,golint
	Vesting_VestingCompleted                       []equilibrium.EventVestingCompleted                              // nolint:stylecheck,golint
	Claims_Claimed                                 []equilibrium.EventClaimed                                       // nolint:stylecheck,golint
	EqSessionManager_ValidatorAdded                []equilibrium.EventValidatorAdded                                // nolint:stylecheck,golint
	EqSessionManager_ValidatorRemoved              []equilibrium.EventValidatorRemoved                              // nolint:stylecheck,golint
	Whitelists_AddedToWhitelist                    []equilibrium.EventAddedToWhitelist                              // nolint:stylecheck,golint
	Whitelists_RemovedFromWhitelist                []equilibrium.EventRemovedFromWhitelist                          // nolint:stylecheck,golint
	EqLockdrop_Lock                                []equilibrium.EventLockdropLock                                  // nolint:stylecheck,golint
	EqLockdrop_Unlock                              []equilibrium.EventLockdropLock                                  // nolint:stylecheck,golint
	Financial_MetricsRecalculated                  []equilibrium.EventMetricsRecalculated                           // nolint:stylecheck,golint
	Financial_AssetMetricsRecalculated             []equilibrium.EventAssetMetricsRecalculated                      // nolint:stylecheck,golint
	Financial_PortfolioMetricsRecalculated         []equilibrium.EventPortfolioMetricsRecalculated                  // nolint:stylecheck,golint
	EqMarginCall_MaintenanceMarginCall             []equilibrium.EventMaintenanceMarginCall                         // nolint:stylecheck,golint
	EqMarginCall_MarginCallExecuted                []equilibrium.EventMarginCallExecuted                            // nolint:stylecheck,golint
	EqAssets_NewAsset                              []equilibrium.EventNewAsset                                      // nolint:stylecheck,golint
	EqAssets_DeleteAsset                           []equilibrium.EventDeleteAsset                                   // nolint:stylecheck,golint
	EqAssets_UpdateAsset                           []equilibrium.EventUpdateAsset                                   // nolint:stylecheck,golint
	EqMultisigSudo_Initialized                     []equilibrium.EventMultisigSudoInitialized                       // nolint:stylecheck,golint
	EqMultisigSudo_KeyAdded                        []equilibrium.EventMultisigSudoKeyAdded                          // nolint:stylecheck,golint
	EqMultisigSudo_KeyRemoved                      []equilibrium.EventMultisigSudoKeyRemoved                        // nolint:stylecheck,golint
	EqMultisigSudo_ThresholdModified               []equilibrium.EventMultisigSudoThresholdModified                 // nolint:stylecheck,golint
	EqMultisigSudo_NewProposal                     []equilibrium.EventMultisigSudoNewProposal                       // nolint:stylecheck,golint
	EqMultisigSudo_ProposalCancelled               []equilibrium.EventMultisigSudoProposalCancelled                 // nolint:stylecheck,golint
	EqMultisigSudo_ProposalApproved                []equilibrium.EventMultisigSudoProposalApproved                  // nolint:stylecheck,golint
	EqMultisigSudo_MultisigSudid                   []equilibrium.EventMultisigSudoSudid                             // nolint:stylecheck,golint
	EqMultisigSudo_SudoFailed                      []equilibrium.EventMultisigSudoFailed                            // nolint:stylecheck,golint
	EqBridge_Remark                                []equilibrium.EventEqBridgeRemark                                // nolint:stylecheck,golint
	CurveDistribution_CurveAdminFeesDistributed    []equilibrium.EventCurveDistributionCurveAdminFeesDistributed    // nolint:stylecheck,golint
	CurveAmm_CreatePool                            []equilibrium.EventCurveAmmCreatePool                            // nolint:stylecheck,golint
	CurveAmm_AddLiquidity                          []equilibrium.EventCurveAmmAddLiquidity                          // nolint:stylecheck,golint
	CurveAmm_TokenExchange                         []equilibrium.EventCurveAmmTokenExchange                         // nolint:stylecheck,golint
	CurveAmm_RemoveLiquidity                       []equilibrium.EventCurveAmmRemoveLiquidity                       // nolint:stylecheck,golint
	CurveAmm_RemoveLiquidityImbalance              []equilibrium.EventCurveAmmRemoveLiquidityImbalance              // nolint:stylecheck,golint
	CurveAmm_RemoveLiquidityOne                    []equilibrium.EventCurveAmmRemoveLiquidityOne                    // nolint:stylecheck,golint
	CurveAmm_WithdrawAdminFees                     []equilibrium.EventCurveAmmWithdrawAdminFees                     // nolint:stylecheck,golint
	GensCrowdloan_Claim                            []equilibrium.EventGensCrowdloanClaim                            // nolint:stylecheck,golint
	EqDex_OrderCreated                             []equilibrium.EventEqDexOrderCreated                             // nolint:stylecheck,golint
	EqDex_OrderDeleted                             []equilibrium.EventEqDexOrderDeleted                             // nolint:stylecheck,golint
	EqDex_Match                                    []equilibrium.EventEqDexMatch                                    // nolint:stylecheck,golint
	EqBridge_FromBridgeTransfer                    []equilibrium.EventEqBridgeFromBridgeTransfer                    // nolint:stylecheck,golint
	EqBridge_ToBridgeTransfer                      []equilibrium.EventEqBridgeToBridgeTransfer                      // nolint:stylecheck,golint
	ChainBridge_FeeChanged                         []equilibrium.ChainBridgeFeeChanged                              // nolint:stylecheck,golint
}
