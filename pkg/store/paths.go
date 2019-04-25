/*
 * Copyright (c) 2018-2019 vChain, Inc. All Rights Reserved.
 * This software is released under GPL3.
 * The full license information can be found under:
 * https://www.gnu.org/licenses/gpl-3.0.en.html
 *
 */

package store

import (
	"os"
	"path/filepath"
)

var dir = DefaultDirName
var configFilepath string

func ensureDir(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.MkdirAll(path, DirPerm); err != nil {
			return err
		}
	}
	return nil
}

func defaultConfigFilepath() string {
	return filepath.Join(dir, configFilename)
}

// SetDir sets the store working directory (e.g. ~/.vcn)
func SetDir(p string) {
	dir = p
}

// ConfigFile returns the config file path
func ConfigFile() string {
	if configFilepath == "" {
		return defaultConfigFilepath()
	}
	return configFilepath
}

// SetConfigFile sets the config file path (e.g. ~/.vcn/config.json)
func SetConfigFile(filepath string) {
	configFilepath = filepath
}
