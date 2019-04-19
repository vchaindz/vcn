/*
 * Copyright (c) 2018-2019 vChain, Inc. All Rights Reserved.
 * This software is released under GPL3.
 * The full license information can be found under:
 * https://www.gnu.org/licenses/gpl-3.0.en.html
 *
 */

package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/dghubble/sling"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/sirupsen/logrus"
	"github.com/vchain-us/vcn/internal/utils"
	"github.com/vchain-us/vcn/pkg/meta"
)

type Wallet struct {
	Address             string `json:"address"`
	CreatedAt           string `json:"createdAt"`
	Name                string `json:"name"`
	PermissionSyncState string `json:"permissionSyncState"`
	LevelSyncState      string `json:"levelSyncState"`
}

type PagedWalletResponse struct {
	Content []Wallet `json:"content"`
}

func CreateKeystore(password string) (pubKey string, wallet string, err error) {
	if password == "" {
		err = makeError("Keystore passphrase cannot be empty", nil)
		return
	}
	ks := keystore.NewKeyStore(meta.WalletDirectory(), keystore.StandardScryptN, keystore.StandardScryptP)
	account, err := ks.NewAccount(password)
	if err != nil {
		return
	}

	pubKey = account.Address.Hex()
	wallet = meta.WalletDirectory()

	_ = TrackPublisher(meta.KeyStoreCreatedEvent)

	return pubKey, wallet, nil
}

func IsWalletSynced(address string) (result bool, err error) {
	authError := new(Error)
	pagedWalletResponse := new(PagedWalletResponse)
	token, err := LoadToken()
	if err != nil {
		log.Fatal(err)
	}
	r, err := sling.New().
		Add("Authorization", "Bearer "+token).
		Get(meta.WalletEndpoint()).
		Receive(pagedWalletResponse, authError)
	if err != nil {
		return false, err
	}
	if r.StatusCode != 200 {
		return false, fmt.Errorf(
			"request failed: %s (%d)", authError.Message,
			authError.Status)
	}
	for _, wallet := range (*pagedWalletResponse).Content {
		if wallet.Address == strings.ToLower(address) {
			return wallet.PermissionSyncState == "SYNCED" &&
				wallet.LevelSyncState == "SYNCED", nil
		}
	}
	return false, fmt.Errorf("no such wallet: %s", address)
}

func HasKeystore() (bool, error) {

	logger().WithFields(logrus.Fields{
		"keystore": meta.WalletDirectory(),
	}).Trace("HasKeystore()")

	files, err := ioutil.ReadDir(meta.WalletDirectory())
	if err != nil {
		logger().WithFields(logrus.Fields{
			"error": err,
		}).Error("ReadDir() failed")
		return false, err
	}
	return len(files) > 0, nil
}

func LoadPublicKeys() (addresses []string, err error) {
	authError := new(Error)
	pagedWalletResponse := new(PagedWalletResponse)
	token, err := LoadToken()
	if err != nil {
		log.Fatal(err)
	}
	r, err := sling.New().
		Add("Authorization", "Bearer "+token).
		Get(meta.WalletEndpoint()).
		Receive(pagedWalletResponse, authError)
	if err != nil {
		log.Fatal(err)
	}
	if r.StatusCode != 200 {
		log.Fatalf("request failed: %s (%d)", authError.Message,
			authError.Status)
	}
	var result []string
	for _, wallet := range (*pagedWalletResponse).Content {
		result = append(result, wallet.Address)
	}
	return result, nil
}

func SyncKeys() {
	authError := new(Error)
	token, err := LoadToken()
	if err != nil {
		log.Fatal(err)
	}
	addresses, err := LoadPublicKeys()
	if err != nil {
		log.Fatal(err)
	}
	localAddress, err := PublicKeyForLocalWallet()
	if err != nil {
		log.Fatal(err)
	}
	if contains(addresses, localAddress) {
		return
	}
	r, err := sling.New().
		Add("Authorization", "Bearer "+token).
		Post(meta.WalletEndpoint()).
		BodyJSON(Wallet{Address: localAddress}).
		Receive(nil, authError)
	if err != nil {
		log.Fatal(err)
	}
	if r.StatusCode != 200 {
		log.Fatalf("request failed: %s (%d)", authError.Message,
			authError.Status)
	}

	_ = TrackPublisher(meta.KeyStoreUploadedEvent)
}

func PublicKeyForLocalWallet() (string, error) {
	reader, err := utils.FirstFile(meta.WalletDirectory())
	if err != nil {
		return "", err
	}
	contents, err := ioutil.ReadAll(reader)
	if err != nil {
		return "", err
	}
	var keyfile map[string]*json.RawMessage
	err = json.Unmarshal(contents, &keyfile)
	if err != nil {
		return "", err
	}
	var localAddress string
	err = json.Unmarshal(*keyfile["address"], &localAddress)
	if err != nil {
		return "", err
	}
	return "0x" + localAddress, nil
}
