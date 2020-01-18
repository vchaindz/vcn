/*
 * Copyright (c) 2018-2020 vChain, Inc. All Rights Reserved.
 * This software is released under GPL3.
 * The full license information can be found under:
 * https://www.gnu.org/licenses/gpl-3.0.en.html
 *
 */

package store

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// WriteYAML writes _in_ data to a YAML file named by _filename_.
func WriteYAML(in interface{}, filename string) error {
	out, err := yaml.Marshal(in)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, out, 0644)
}

// ReadYAML reads the file named by _filename and assigns the decoded YAML content into the _out_ value.
func ReadYAML(out interface{}, filename string) error {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(b, out)
}
