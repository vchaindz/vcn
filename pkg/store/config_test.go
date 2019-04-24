/*
 * Copyright (c) 2018-2019 vChain, Inc. All Rights Reserved.
 * This software is released under GPL3.
 * The full license information can be found under:
 * https://www.gnu.org/licenses/gpl-3.0.en.html
 *
 */

package store

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func mkTmpForConfig(t *testing.T) string {
	tdir, err := ioutil.TempDir("", "vcn-test-store-config")
	if err != nil {
		t.Fatal(err)
	}
	return tdir
}

func TestLoadConfig(t *testing.T) {
	tdir := mkTmpForConfig(t)
	SetDir(tdir + "/" + DefaultDirName)

	assert.Nil(t, Config())

	err := LoadConfig()
	assert.NoError(t, err)
	assert.NotNil(t, Config())
	assert.Equal(t, filepath.Join(tdir, DefaultDirName, configFilename), ConfigFile())
	assert.FileExists(t, ConfigFile())
	assert.NotNil(t, Config())
}

func TestSaveConfig(t *testing.T) {
	tdir := mkTmpForConfig(t)
	SetDir(tdir + "/" + DefaultDirName)

	email := "example@example.net"

	cfg = &config{
		CurrentContext: email,
		Users: []*User{
			&User{
				Email: email,
				Keystores: []*Keystore{
					&Keystore{
						Path: filepath.Join(tdir, "u", email, "keystore"),
					},
				},
			},
		},
	}

	err := SaveConfig()
	assert.NoError(t, err)

	LoadConfig()
	assert.Equal(t, email, Config().CurrentContext)
}
