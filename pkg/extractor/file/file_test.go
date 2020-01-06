/*
 * Copyright (c) 2018-2020 vChain, Inc. All Rights Reserved.
 * This software is released under GPL3.
 * The full license information can be found under:
 * https://www.gnu.org/licenses/gpl-3.0.en.html
 *
 */

package file

import (
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/stretchr/testify/assert"

	"os"
	"testing"

	"github.com/vchain-us/vcn/pkg/uri"
)

func TestFile(t *testing.T) {
	file, err := ioutil.TempFile("", "vcn-test-scheme-file")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(file.Name())
	err = ioutil.WriteFile(file.Name(), []byte("123\n"), 0644)
	if err != nil {
		log.Fatal(err)
	}
	u, _ := uri.Parse("file://" + file.Name())

	a, err := Artifact(u)
	assert.NoError(t, err)
	assert.NotNil(t, a)
	assert.Equal(t, Scheme, a.Kind)
	assert.Equal(t, filepath.Base(file.Name()), a.Name)
	assert.Equal(t, "181210f8f9c779c26da1d9b2075bde0127302ee0e3fca38c9a83f5b1dd8e5d3b", a.Hash)

	u, _ = uri.Parse(file.Name())
	a, err = Artifact(u)
	assert.NoError(t, err)
	assert.NotNil(t, a)
	assert.Equal(t, Scheme, a.Kind)
	assert.Equal(t, filepath.Base(file.Name()), a.Name)
	assert.Equal(t, "181210f8f9c779c26da1d9b2075bde0127302ee0e3fca38c9a83f5b1dd8e5d3b", a.Hash)

	u, _ = uri.Parse("../../../docs/vcncheatsheet.pdf")
	a, err = Artifact(u)
	assert.NoError(t, err)
	assert.NotNil(t, a)
	assert.Equal(t, a.ContentType, "application/pdf")
}
