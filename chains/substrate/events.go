// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package substrate

import (
	"fmt"
	"github.com/ChainSafe/chainbridge-utils/core"
	"math/big"

	events "github.com/ChainSafe/chainbridge-substrate-events"
	"github.com/ChainSafe/chainbridge-utils/msg"
	"github.com/ChainSafe/log15"

	"github.com/ChainSafe/ChainBridge/shared/equilibrium"
)

type eventName string
type eventHandler func(interface{}, core.MessageContext, log15.Logger) (msg.Message, core.MessageContext, error)

const FungibleTransfer eventName = "FungibleTransfer"
const NonFungibleTransfer eventName = "NonFungibleTransfer"
const GenericTransfer eventName = "GenericTransfer"

var Subscriptions = []struct {
	name    eventName
	handler eventHandler
}{
	{FungibleTransfer, fungibleTransferHandler},
	{NonFungibleTransfer, nonFungibleTransferHandler},
	{GenericTransfer, genericTransferHandler},
}

func fungibleTransferHandler(evtI interface{}, messageContext core.MessageContext, log log15.Logger) (msg.Message, core.MessageContext, error) {
	evt, ok := evtI.(events.EventFungibleTransfer)
	if !ok {
		return msg.Message{}, make(core.MessageContext), fmt.Errorf("failed to cast EventFungibleTransfer type")
	}

	resourceId := msg.ResourceId(evt.ResourceId)
	log.Info("Got fungible transfer event!", "destination", evt.Destination, "resourceId", resourceId.Hex(), "amount", evt.Amount)

	equilibrium.EventFungibleTransfer("EventFound", fmt.Sprintf("(S->E) Handle FungibleTransferEvent"), evt, messageContext)

	factor := big.NewInt(1000000000)
	amount := new(big.Int).Mul(evt.Amount.Int, factor)

	return msg.NewFungibleTransfer(
		0, // Unset
		msg.ChainId(evt.Destination),
		msg.Nonce(evt.DepositNonce),
		amount,
		resourceId,
		evt.Recipient,
	), messageContext, nil
}

func nonFungibleTransferHandler(evtI interface{}, messageContext core.MessageContext, log log15.Logger) (msg.Message, core.MessageContext, error) {
	evt, ok := evtI.(events.EventNonFungibleTransfer)
	if !ok {
		return msg.Message{}, make(core.MessageContext), fmt.Errorf("failed to cast EventNonFungibleTransfer type")
	}

	log.Info("Got non-fungible transfer event!", "destination", evt.Destination, "resourceId", evt.ResourceId)

	return msg.NewNonFungibleTransfer(
		0, // Unset
		msg.ChainId(evt.Destination),
		msg.Nonce(evt.DepositNonce),
		msg.ResourceId(evt.ResourceId),
		big.NewInt(0).SetBytes(evt.TokenId[:]),
		evt.Recipient,
		evt.Metadata,
	), messageContext, nil
}

func genericTransferHandler(evtI interface{}, messageContext core.MessageContext, log log15.Logger) (msg.Message, core.MessageContext, error) {
	evt, ok := evtI.(events.EventGenericTransfer)
	if !ok {
		return msg.Message{}, make(core.MessageContext), fmt.Errorf("failed to cast EventGenericTransfer type")
	}

	log.Info("Got generic transfer event!", "destination", evt.Destination, "resourceId", evt.ResourceId)

	return msg.NewGenericTransfer(
		0, // Unset
		msg.ChainId(evt.Destination),
		msg.Nonce(evt.DepositNonce),
		msg.ResourceId(evt.ResourceId),
		evt.Metadata,
	), messageContext, nil
}
