/*
 * Copyright (c) 2018-2019 vChain, Inc. All Rights Reserved.
 * This software is released under GPL3.
 * The full license information can be found under:
 * https://www.gnu.org/licenses/gpl-3.0.en.html
 *
 */

package cli

import (
	"context"
	"log"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
	"github.com/vchain-us/vcn/internal/blockchain"
	"github.com/vchain-us/vcn/internal/errors"
	"github.com/vchain-us/vcn/internal/utils"
	"github.com/vchain-us/vcn/pkg/api"
	"github.com/vchain-us/vcn/pkg/logs"
	"github.com/vchain-us/vcn/pkg/meta"
)

func commitHash(hash string, passphrase string, filename string, fileSize int64, status meta.Status, visibility meta.Visibility) (ret bool, code int) {
	reader, err := utils.FirstFile(meta.WalletDirectory())
	if err != nil {
		logs.LOG.WithFields(logrus.Fields{
			"error": err,
		}).Fatal("Could not load keystore")
	}
	transactor, err := bind.NewTransactor(reader, passphrase)
	if err != nil {
		log.Fatal(err)
	}
	walletSynced, err := api.IsWalletSynced(transactor.From.Hex())
	if err != nil {
		logs.LOG.WithFields(logrus.Fields{
			"error": err,
		}).Error("Could not load wallets")
		errors.PrintErrorURLCustom("wallet", 400)
		os.Exit(1)
	}
	if !walletSynced {
		logs.LOG.Error("\n", filename, " cannot be signed with CodeNotary. We are "+
			"finalizing your account configuration.\nWe will complete the "+
			"configuration shortly and we will update you as soon as this "+
			"is done.\nWe are sorry for the inconvenience and would like "+
			"to thank you for your patience.")
		os.Exit(1)
	}
	transactor.GasLimit = meta.GasLimit()
	transactor.GasPrice = meta.GasPrice()
	client, err := ethclient.Dial(meta.MainNetEndpoint())
	if err != nil {
		logs.LOG.WithFields(logrus.Fields{
			"error":   err,
			"network": meta.MainNetEndpoint(),
		}).Fatal("Cannot connect to blockchain")
	}
	address := common.HexToAddress(meta.AssetsRelayContractAddress())
	instance, err := blockchain.NewAssetsRelay(address, client)
	if err != nil {
		logs.LOG.WithFields(logrus.Fields{
			"error":    err,
			"contract": meta.AssetsRelayContractAddress(),
		}).Fatal("Cannot instantiate contract")
	}
	tx, err := instance.Sign(transactor, hash, big.NewInt(int64(status)))
	if err != nil {
		logs.LOG.WithFields(logrus.Fields{
			"error": err,
			"hash":  hash,
		}).Fatal("method <Sign> failed")
	}
	timeout, err := waitForTx(tx.Hash(), meta.TxVerificationRounds(), meta.PollInterval())
	if err != nil {
		logs.LOG.WithFields(logrus.Fields{
			"error": err,
		}).Error("Could not write to blockchain")
		errors.PrintErrorURLCustom("blockchain-permission", 403)
		os.Exit(1)
	}
	if timeout {
		logs.LOG.WithFields(logrus.Fields{
			"error": err,
		}).Fatal("Writing to blockchain timed out")
	}
	publicKey, err := api.PublicKeyForLocalWallet()
	if err != nil {
		log.Fatal(err)
	}
	verification, err := api.BlockChainVerifyMatchingPublicKey(hash, transactor.From.Hex())
	if err != nil {
		log.Fatal(err)
	}
	err = api.CreateArtifact(verification, publicKey, filename, hash, fileSize, visibility, status)
	if err != nil {
		log.Fatal(err)
	}

	return true, 0
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
