/*
 * Copyright (c) 2018-2019 vChain, Inc. All Rights Reserved.
 * This software is released under GPL3.
 * The full license information can be found under:
 * https://www.gnu.org/licenses/gpl-3.0.en.html
 *
 */

package store

import (
	"fmt"
	"path/filepath"

	"github.com/ethereum/go-ethereum/accounts/keystore"
)

func ksInstance(keydir string) *keystore.KeyStore {
	return keystore.NewKeyStore(keydir, keystore.StandardScryptN, keystore.StandardScryptP)
}

// PubKeys return all User's public keys found
func (u User) PubKeys() []string {
	pubs := []string{}
	for _, ks := range u.Keystores {
		ksi := ksInstance(ks.Path)
		for _, a := range ksi.Accounts() {
			pubs = append(pubs, a.Address.Hex())
		}
	}
	return pubs
}

// AddKeystore adds a keystore dir to the User
func (u *User) AddKeystore(keydir string) (k Keystore, err error) {

	if u == nil || u.Email == "" {
		err = fmt.Errorf("user has not been initialized")
		return
	}

	path, err := filepath.Abs(keydir)
	if err != nil {
		return
	}

	k = Keystore{
		Path: path,
	}

	u.Keystores = append(u.Keystores, k)

	return
}

// AddDefaultKeystore adds the default keystore dir to the User
func (u *User) AddDefaultKeystore() (k Keystore, err error) {
	return u.AddKeystore(
		filepath.Join(dir, "u", u.Email, "keystore"),
	)
}

// CreateKey generates a new key and stores it into the Keystore directory,
// encrypting it with the passphrase.
func (k Keystore) CreateKey(passphrase string) (pubKey string, err error) {
	if passphrase == "" {
		err = fmt.Errorf("passphrase cannot be empty")
		return
	}

	if err = ensureDir(k.Path); err != nil {
		return
	}

	account, err := ksInstance(k.Path).NewAccount(passphrase)
	if err != nil {
		return
	}

	pubKey = account.Address.Hex()

	return pubKey, nil
}
