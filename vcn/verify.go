/*
 * Copyright (c) 2018-2019 vChain, Inc. All Rights Reserved.
 * This software is released under GPL3.
 * The full license information can be found under:
 * https://www.gnu.org/licenses/gpl-3.0.en.html
 *
 */

package main

import (
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
)

type BlockchainVerification struct {
	Owner     common.Address
	Level     Level
	Status    Status
	Timestamp time.Time
}

func BlockChainVerify(hash string) (verification *BlockchainVerification, err error) {
	LOG.WithFields(logrus.Fields{
		"hash": hash,
	}).Trace("BlockChainVerify")
	client, err := ethclient.Dial(MainNetEndpoint())
	if err != nil {
		return nil, err
	}
	contractAddress := common.HexToAddress(AssetsRelayContractAddress())
	instance, err := NewAssetsRelay(contractAddress, client)
	if err != nil {
		return nil, err
	}
	address, level, status, timestamp, err := instance.Verify(nil, hash)
	if err != nil {
		return nil, err
	}
	verification = new(BlockchainVerification)
	verification.Owner = address
	verification.Level = Level(level.Int64())
	verification.Status = Status(status.Int64())
	verification.Timestamp = time.Unix(timestamp.Int64(), 0)
	LOG.
		WithField("verification", verification).
		Trace("Received blockchain verification")
	return verification, nil
}

func BlockChainVerifyMatchingPublicKey(hash string, publicKey string) (verification *BlockchainVerification, err error) {
	LOG.WithFields(logrus.Fields{
		"hash": hash,
	}).Trace("BlockChainVerifyMatchingPublicKey")
	client, err := ethclient.Dial(MainNetEndpoint())
	if err != nil {
		return nil, err
	}
	contractAddress := common.HexToAddress(AssetsRelayContractAddress())
	instance, err := NewAssetsRelay(contractAddress, client)
	if err != nil {
		return nil, err
	}
	count, err := instance.GetAssetCountForHash(nil, hash)
	if err != nil {
		return nil, err
	}
	for i := count.Int64() - 1; i >= 0; i-- {
		address, level, status, timestamp, err := instance.VerifyByIndex(nil, hash, big.NewInt(i))
		if err != nil {
			return nil, err
		}
		if address.Hex() == publicKey {
			return &BlockchainVerification{
				Owner:     address,
				Level:     Level(level.Int64()),
				Status:    Status(status.Int64()),
				Timestamp: time.Unix(timestamp.Int64(), 0),
			}, nil
		}
	}
	return nil, fmt.Errorf("no matching asset for hash %s and publicKey %s", hash, publicKey)
}
