// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	"errors"
	"fmt"
	"github.com/ChainSafe/chainbridge-utils/msg"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"math/big"

	"github.com/ChainSafe/ChainBridge/shared/equilibrium"

	erc20 "github.com/ChainSafe/ChainBridge/bindings/ERC20"
)

func (l *listener) handleErc20DepositedEvent(destId msg.ChainId, nonce msg.Nonce) (msg.Message, error) {
	l.log.Info("Handling fungible deposit event", "dest", destId, "nonce", nonce)

	record, err := l.erc20HandlerContract.GetDepositRecord(&bind.CallOpts{From: l.conn.Keypair().CommonAddress()}, uint64(nonce), uint8(destId))
	if err != nil {
		l.log.Error("Error Unpacking ERC20 Deposit Record", "err", err)
		return msg.Message{}, err
	}

	tokenAddress := record.TokenAddress
	erc20Contract, err := erc20.NewERC20(tokenAddress, l.conn.Client())
	if err != nil {
		l.log.Error("Error Creating ERC20 Contract From TokenAddress", "err", err)
		return msg.Message{}, err
	}

	tokenDecimals, err := erc20Contract.Decimals(&bind.CallOpts{From: l.conn.Keypair().CommonAddress()})
	if err != nil {
		l.log.Error("Error Getting Token Decimals")
		return msg.Message{}, err
	}

	if tokenDecimals > 36 {
		l.log.Error("Token Has Very Strange Decimals Count")
		return msg.Message{}, errors.New("token has very strange decimals count")
	}

	//// make amount always be 18 decimals
	amount := record.Amount
	if tokenDecimals <= 18 {
		multiplier := new(big.Int).Exp(big.NewInt(10), big.NewInt(18 - int64(tokenDecimals)), nil)
		amount = new(big.Int).Mul(amount, multiplier)
	} else {
		divisor := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(tokenDecimals) - 18), nil)
		amount = new(big.Int).Div(amount, divisor)
	}

	direction := "E->S"
	//equilibrium.Info(fmt.Sprintf("(%s) Scale value %s -> %s", direction, oldAmount, amount.String()))

	result := msg.NewFungibleTransfer(
		l.cfg.id,
		destId,
		nonce,
		amount,
		record.ResourceID,
		record.DestinationRecipientAddress,
	)

	equilibrium.Message("EventFound", fmt.Sprintf("(%s) Handle Erc20DepositedEvent", direction),
		result, nil, nil)

	return result, nil
}

func (l *listener) handleErc721DepositedEvent(destId msg.ChainId, nonce msg.Nonce) (msg.Message, error) {
	l.log.Info("Handling nonfungible deposit event")

	record, err := l.erc721HandlerContract.GetDepositRecord(&bind.CallOpts{From: l.conn.Keypair().CommonAddress()}, uint64(nonce), uint8(destId))
	if err != nil {
		l.log.Error("Error Unpacking ERC20 Deposit Record", "err", err)
		return msg.Message{}, err
	}

	return msg.NewNonFungibleTransfer(
		l.cfg.id,
		destId,
		nonce,
		record.ResourceID,
		record.TokenID,
		record.DestinationRecipientAddress,
		record.MetaData,
	), nil
}

func (l *listener) handleGenericDepositedEvent(destId msg.ChainId, nonce msg.Nonce) (msg.Message, error) {
	l.log.Info("Handling generic deposit event")

	record, err := l.genericHandlerContract.GetDepositRecord(&bind.CallOpts{From: l.conn.Keypair().CommonAddress()}, uint64(nonce), uint8(destId))
	if err != nil {
		l.log.Error("Error Unpacking Generic Deposit Record", "err", err)
		return msg.Message{}, nil
	}

	return msg.NewGenericTransfer(
		l.cfg.id,
		destId,
		nonce,
		record.ResourceID,
		record.MetaData[:],
	), nil
}
