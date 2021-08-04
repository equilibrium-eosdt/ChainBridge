// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	"github.com/ChainSafe/ChainBridge/bindings/Bridge"
	"github.com/ChainSafe/ChainBridge/bindings/ERC20Handler"
	"github.com/ChainSafe/ChainBridge/shared/equilibrium"
	"github.com/ChainSafe/chainbridge-utils/core"
	"github.com/ChainSafe/ChainBridge/shared/equilibrium/metrics"
	"github.com/ChainSafe/chainbridge-utils/msg"
)

var _ core.Writer = &writer{}

// https://github.com/ChainSafe/chainbridge-solidity/blob/b5ed13d9798feb7c340e737a726dd415b8815366/contracts/Bridge.sol#L20
var PassedStatus uint8 = 2
var TransferredStatus uint8 = 3
var CancelledStatus uint8 = 4

type writer struct {
	cfg            Config
	conn           Connection
	bridgeContract *Bridge.Bridge // instance of bound receiver bridgeContract
	erc20HandlerContract   *ERC20Handler.ERC20Handler
	log            equilibrium.TransferLogger
	stop           <-chan int
	sysErr         chan<- error // Reports fatal error to core
	metrics        *metrics.ChainMetrics
}

// NewWriter creates and returns writer
func NewWriter(conn Connection, cfg *Config, log equilibrium.TransferLogger, stop <-chan int, sysErr chan<- error, m *metrics.ChainMetrics) *writer {
	return &writer{
		cfg:     *cfg,
		conn:    conn,
		log:     log,
		stop:    stop,
		sysErr:  sysErr,
		metrics: m,
	}
}

func (w *writer) start() error {
	w.log.Debug("Starting ethereum writer...")
	return nil
}

func (w *writer) setContracts(bridge *Bridge.Bridge, erc20Handler *ERC20Handler.ERC20Handler) {
	w.bridgeContract = bridge
	w.erc20HandlerContract = erc20Handler
}

// ResolveMessage handles any given message based on type
// A bool is returned to indicate failure/success, this should be ignored except for within tests.
func (w *writer) ResolveMessage(m msg.Message, context core.MessageContext) bool {
	w.log.Info("Attempting to resolve message", "type", m.Type, "src", m.Source, "dst", m.Destination, "nonce", m.DepositNonce, "rId", m.ResourceId.Hex())

	switch m.Type {
	case msg.FungibleTransfer:
		return w.createErc20Proposal(m, context)
	case msg.NonFungibleTransfer:
		return w.createErc721Proposal(m)
	case msg.GenericTransfer:
		return w.createGenericDepositProposal(m)
	default:
		w.log.ErrorTransfer("Unknown message type received", context, "type", m.Type)
		return false
	}
}
