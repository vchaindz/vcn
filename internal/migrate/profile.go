/*
 * Copyright (c) 2018-2019 vChain, Inc. All Rights Reserved.
 * This software is released under GPL3.
 * The full license information can be found under:
 * https://www.gnu.org/licenses/gpl-3.0.en.html
 *
 */
package migrate

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/dghubble/sling"
	"github.com/mitchellh/go-homedir"
	"github.com/vchain-us/vcn/pkg/api"
	"github.com/vchain-us/vcn/pkg/meta"
	"github.com/vchain-us/vcn/pkg/store"
)

const vcn03x = ".vcn"
const wallets03x = "wallets"
const token03x = "t"

// From03x silently migrates v0.3.x vcn's profile data to v0.4.x store.
func From03x() {
	home, err := homedir.Dir()
	if err != nil {
		return
	}
	dir := filepath.Join(home, vcn03x)
	tokenFile := filepath.Join(dir, token03x)
	walletsDir := filepath.Join(dir, wallets03x)

	defer os.Remove(tokenFile)
	t, err := ioutil.ReadFile(tokenFile)
	if err != nil {
		return
	}

	restError := new(api.Error)
	response := struct {
		Email string `json:"email"`
	}{}
	sling.New().Add("Authorization", "Bearer "+strings.TrimSpace(string(t))).
		Get(meta.PublisherEndpoint()).
		Receive(&response, restError)

	if response.Email != "" {
		// ensure v0.4.x user dir
		uDir := filepath.Join(dir, "u", response.Email)
		os.MkdirAll(uDir, store.DirPerm)

		// move wallets to user dir
		toDir := filepath.Join(uDir, "keystore")
		os.Rename(walletsDir, toDir)
	}
}
