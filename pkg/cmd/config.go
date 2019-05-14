/*
 * Copyright (c) 2018-2019 vChain, Inc. All Rights Reserved.
 * This software is released under GPL3.
 * The full license information can be found under:
 * https://www.gnu.org/licenses/gpl-3.0.en.html
 *
 */

package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/vchain-us/vcn/pkg/extractor"
	"github.com/vchain-us/vcn/pkg/extractor/docker"
	"github.com/vchain-us/vcn/pkg/extractor/file"

	"github.com/vchain-us/vcn/pkg/store"

	"github.com/mitchellh/go-homedir"
)

// initConfig reads in config file and ENV variables if set.
func initConfig() {

	// Register metadata extractors
	extractor.Register("", file.Artifact)
	extractor.Register(file.Scheme, file.Artifact)
	extractor.Register(docker.Scheme, docker.Artifact)

	// Find home directory
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	store.SetDir(filepath.Join(home, store.DefaultDirName))

	// Load config
	if cfgFile != "" {
		store.SetConfigFile(cfgFile)
		// fixme(leogr): if this enabled it will interfere with --output=json --output=yaml
		// fmt.Println("Using config file: ", store.ConfigFile())
	}
	if err := store.LoadConfig(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
