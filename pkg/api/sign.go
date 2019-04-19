/*
 * Copyright (c) 2018-2019 vChain, Inc. All Rights Reserved.
 * This software is released under GPL3.
 * The full license information can be found under:
 * https://www.gnu.org/licenses/gpl-3.0.en.html
 *
 */

package api

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
	"github.com/vchain-us/vcn/internal/blockchain"
	"github.com/vchain-us/vcn/internal/errors"
	"github.com/vchain-us/vcn/internal/utils"
	"github.com/vchain-us/vcn/pkg/meta"
)

func (a *Artifact) Sign(passphrase string, state meta.Status, visibility meta.Visibility) error {
	if a == nil {
		return makeFatal("nil artifact", nil)
	}
	if a.Hash == "" {
		return makeError("asset's hash is missing", nil)
	}
	if a.Name == "" {
		return makeError("asset's name is missing", nil)
	}
	_ = TrackPublisher(meta.VcnSignEvent)
	_ = TrackSign(a.Hash, a.Name, state)
	return commitHash(passphrase, a.Hash, a.Name, a.Size, state, visibility)
}

// todo(leogr): refactor
func commitHash(
	passphrase string,
	hash string,
	name string,
	fileSize uint64,
	status meta.Status,
	visibility meta.Visibility,
) (err error) {
	reader, err := utils.FirstFile(meta.WalletDirectory())
	if err != nil {
		err = makeFatal(
			"Could not load keystore",
			logrus.Fields{
				"error": err,
			},
		)
		return
	}
	transactor, err := bind.NewTransactor(reader, passphrase)
	if err != nil {
		return
	}
	walletSynced, err := IsWalletSynced(transactor.From.Hex())
	if err != nil {
		// fixme(leogr): logging, and avoid to output directly
		errors.PrintErrorURLCustom("wallet", 400)
		err = makeError(
			"Could not load wallets",
			logrus.Fields{
				"error": err,
			},
		)
		return
	}
	if !walletSynced {
		logger().Error("\n", name, " cannot be signed with CodeNotary. We are "+
			"finalizing your account configuration.\nWe will complete the "+
			"configuration shortly and we will update you as soon as this "+
			"is done.\nWe are sorry for the inconvenience and would like "+
			"to thank you for your patience.")
		err = makeError("Wallet not yet synced", nil)
		return
	}
	transactor.GasLimit = meta.GasLimit()
	transactor.GasPrice = meta.GasPrice()
	client, err := ethclient.Dial(meta.MainNetEndpoint())
	if err != nil {
		logger().WithFields(logrus.Fields{
			"error":   err,
			"network": meta.MainNetEndpoint(),
		}).Fatal("Cannot connect to blockchain")
		err = makeError("Cannot connect to blockchain", nil)
		return
	}
	address := common.HexToAddress(meta.AssetsRelayContractAddress())
	instance, err := blockchain.NewAssetsRelay(address, client)
	if err != nil {
		err = makeFatal(
			"Cannot instantiate contract",
			logrus.Fields{
				"error":    err,
				"contract": meta.AssetsRelayContractAddress(),
			},
		)
		return
	}
	tx, err := instance.Sign(transactor, hash, big.NewInt(int64(status)))
	if err != nil {
		err = makeFatal(
			"method <sign> failed",
			logrus.Fields{
				"error": err,
				"hash":  hash,
			},
		)
		return
	}
	timeout, err := waitForTx(tx.Hash(), meta.TxVerificationRounds(), meta.PollInterval())
	if err != nil {
		// fixme(leogr): logging, and avoid to output directly
		errors.PrintErrorURLCustom("blockchain-permission", 403)
		err = makeFatal(
			"Could not write to blockchain",
			logrus.Fields{
				"error": err,
			},
		)
		return
	}
	if timeout {
		err = makeFatal(
			"Writing to blockchain timed out",
			logrus.Fields{
				"error": err,
			},
		)
		return
	}
	publicKey, err := PublicKeyForLocalWallet()
	if err != nil {
		return
	}
	verification, err := BlockChainVerifyMatchingPublicKey(hash, transactor.From.Hex())
	if err != nil {
		return
	}
	err = CreateArtifact(verification, publicKey, name, hash, fileSize, visibility, status)
	if err != nil {
		return
	}

	return
}

func waitForTx(tx common.Hash, maxRounds uint64, pollInterval time.Duration) (timeout bool, err error) {
	client, err := ethclient.Dial(meta.MainNetEndpoint())
	if err != nil {
		return false, err
	}
	for i := uint64(0); i < maxRounds; i++ {
		_, pending, err := client.TransactionByHash(context.Background(), tx)
		if err != nil {
			return false, err
		}
		if !pending {
			return false, nil
		}
		time.Sleep(pollInterval)
	}
	return true, nil
}
