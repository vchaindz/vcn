/*
 * Copyright (c) 2018-2019 vChain, Inc. All Rights Reserved.
 * This software is released under GPL3.
 * The full license information can be found under:
 * https://www.gnu.org/licenses/gpl-3.0.en.html
 *
 */

package cli

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestHash(t *testing.T) {
	file, err := ioutil.TempFile("", "example")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(file.Name())
	err = ioutil.WriteFile(file.Name(), []byte("123\n"), 0644)
	if err != nil {
		log.Fatal(err)
	}
	hash := hash(file.Name())
	if hash != "181210f8f9c779c26da1d9b2075bde0127302ee0e3fca38c9a83f5b1dd8e5d3b" {
		t.Error("hash does not match:", hash)
	}
}
