/*
 * Copyright (c) 2018-2019 vChain, Inc. All Rights Reserved.
 * This software is released under GPL3.
 * The full license information can be found under:
 * https://www.gnu.org/licenses/gpl-3.0.en.html
 *
 */

package api

import (
	"fmt"
	"strings"

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

func (u User) isWalletSynced(address string) (result bool, err error) {
	authError := new(Error)
	pagedWalletResponse := new(PagedWalletResponse)
	r, err := newSling(u.token()).
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

func (u User) loadPublicKeys() (addresses []string, err error) {
	authError := new(Error)
	pagedWalletResponse := new(PagedWalletResponse)
	r, err := newSling(u.token()).
		Get(meta.WalletEndpoint()).
		Receive(pagedWalletResponse, authError)
	if err != nil {
		return
	}
	if r.StatusCode != 200 {
		err = makeError(
			fmt.Sprintf("request failed: %s (%d)", authError.Message, authError.Status),
			nil,
		)
		return
	}
	var result []string
	for _, wallet := range (*pagedWalletResponse).Content {
		result = append(result, wallet.Address)
	}
	return result, nil
}

func (u User) SyncKeys() error {

	hasAuth, err := u.IsAuthenticated()
	if err != nil {
		return err
	}
	if !hasAuth {
		return makeError("user not authenticated, please login", nil)
	}
	addresses, err := u.loadPublicKeys()
	if err != nil {
		return err
	}
	for _, localAddress := range u.cfg.PubKeys() {
		if contains(addresses, localAddress) {
			continue
		}

		authError := new(Error)
		r, err := newSling(u.token()).
			Post(meta.WalletEndpoint()).
			BodyJSON(Wallet{Address: localAddress}).
			Receive(nil, authError)
		if err != nil {
			return err
		}

		// If a wallet is already synced, just skip it
		if authError.Status == 409 {
			continue
		}
		if r.StatusCode != 200 {
			return makeFatal(
				fmt.Sprintf("request failed: %s (%d)", authError.Message, authError.Status),
				nil,
			)
		}

		_ = TrackPublisher(&u, meta.KeyUploadedEvent)
	}

	return nil
}
