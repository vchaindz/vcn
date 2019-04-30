/*
 * Copyright (c) 2018-2019 vChain, Inc. All Rights Reserved.
 * This software is released under GPL3.
 * The full license information can be found under:
 * https://www.gnu.org/licenses/gpl-3.0.en.html
 *
 */
package migrate

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/mitchellh/go-homedir"
	"github.com/vchain-us/vcn/pkg/api"
	"github.com/vchain-us/vcn/pkg/store"
)

const vcn03x = ".vcn"
const wallets03x = "wallets"
const token03x = "t"

func ksInstance(keydir string) *keystore.KeyStore {
	return keystore.NewKeyStore(keydir, keystore.StandardScryptN, keystore.StandardScryptP)
}

func contains(xs []string, x string) bool {
	for _, a := range xs {
		if a == x {
			return true
		}
	}
	return false
}

// From03x silently migrates v0.3.x vcn's profile data to v0.4.x store.
func From03x(user *api.User) {
	home, err := homedir.Dir()
	if err != nil {
		return
	}
	dir := filepath.Join(home, vcn03x)

	// Cleanup old token
	tokenFile := filepath.Join(dir, token03x)
	os.Remove(tokenFile)

	// Find keys
	walletsDir := filepath.Join(dir, wallets03x)
	accs := ksInstance(walletsDir).Accounts()
	if len(accs) < 1 {
		return // nothing to migrate
	}

	// new user keystore dir
	newKeyStore, err := user.DefaultKeystore()
	if err != nil {
		return
	}
	os.MkdirAll(newKeyStore.Path, store.DirPerm)

	if wallets, err := user.Wallets(); err == nil {
		for _, w := range wallets {
			for _, a := range accs {
				if strings.ToLower(a.Address.Hex()) == w {
					oldPath := a.URL.Path
					newPath := filepath.Join(newKeyStore.Path, filepath.Base(a.URL.Path))
					if err := os.Rename(oldPath, newPath); err == nil {
						fmt.Printf("Pub key '%s' has been migrated to new profile dir.", w)
					}
				}
			}
		}
	}
}
