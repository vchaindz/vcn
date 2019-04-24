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
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/keystore"
)

// AddKeystore adds a keystore dir to the User
func (u *User) AddKeystore(keydir string) (*Keystore, error) {

	if u == nil || u.Email == "" {
		return nil, fmt.Errorf("user has not been initialized")
	}

	path, err := filepath.Abs(keydir)
	if err != nil {
		return nil, err
	}

	// Check if already exists
	for _, k := range u.Keystores {
		if k.Path == path {
			return k, nil
		}
	}

	k := &Keystore{
		Path: path,
	}

	u.Keystores = append(u.Keystores, k)

	return k, nil
}

// DefaultKeystore returns the default keystore
func (u *User) DefaultKeystore() (*Keystore, error) {
	return u.AddKeystore(
		filepath.Join(dir, "u", u.Email, "keystore"),
	)
}

func ksInstance(keydir string) *keystore.KeyStore {
	return keystore.NewKeyStore(keydir, keystore.StandardScryptN, keystore.StandardScryptP)
}

// OpenKey opens the named pubKey for reading
func (u User) OpenKey(pubKey string) (io.Reader, error) {
	for _, ks := range u.Keystores {
		ksi := ksInstance(ks.Path)
		for _, a := range ksi.Accounts() {
			if a.Address.Hex() == pubKey {
				return os.Open(a.URL.Path)
			}
		}
	}
	return nil, fmt.Errorf("no key found matching %s", pubKey)
}

// HasKey returns true if the User has at least one public key
func (u User) HasKey() bool {
	for _, ks := range u.Keystores {
		ksi := ksInstance(ks.Path)
		if len(ksi.Accounts()) > 0 {
			return true
		}
	}
	return false
}

// PubKeys returns all User's public keys found
func (u User) PubKeys() []string {
	pubs := []string{}
	for _, ks := range u.Keystores {
		ksi := ksInstance(ks.Path)
		for _, a := range ksi.Accounts() {
			pubs = append(pubs, strings.ToLower(a.Address.Hex()))
		}
	}
	return pubs
}

// LastPubKey returns the last added User's public key
func (u User) LastPubKey() string {
	pubs := u.PubKeys()
	l := len(pubs)
	if l > 0 {
		return pubs[l-1]
	}
	return ""
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

	pubKey = strings.ToLower(account.Address.Hex())

	return pubKey, nil
}
