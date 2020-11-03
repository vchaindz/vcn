/*
 * Copyright (c) 2018-2020 vChain, Inc. All Rights Reserved.
 * This software is released under GPL3.
 * The full license information can be found under:
 * https://www.gnu.org/licenses/gpl-3.0.en.html
 *
 */

package dir

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vchain-us/vcn/pkg/uri"
)

func TestArtifact(t *testing.T) {

	tmpDir, err := ioutil.TempDir("", "TempDir")
	if err != nil {
		t.Fatal(err)
	}

	tmpFile := filepath.Join(tmpDir, "file")
	err = ioutil.WriteFile(tmpFile, nil, 0644)
	if err != nil {
		t.Fatal(err)
	}
	// dir - OK
	u, _ := uri.Parse("dir://" + tmpDir)
	artifacts, err := Artifact(u)
	assert.NoError(t, err)
	assert.NotNil(t, artifacts)
	assert.Equal(t, "dir", artifacts[0].Kind)
	assert.Equal(t, filepath.Base(tmpDir), artifacts[0].Name)
	assert.NotEmpty(t, artifacts[0].Hash)

	// dir (no schema) - Suppressed this behaviour. With wildcard it's possible to specify a dir so this is replaced
	/*u, _ = uri.Parse(tmpDir)
	artifacts, err = Artifact(u)
	assert.NoError(t, err)
	assert.NotNil(t, artifacts[0])
	assert.Equal(t, "dir", artifacts[0].Kind)
	assert.Equal(t, filepath.Base(tmpDir), artifacts[0].Name)
	assert.NotEmpty(t, artifacts[0].Hash)*/

	// wrong schema - SKIP (no error)
	u, _ = uri.Parse("file://" + tmpDir)
	artifacts, err = Artifact(u)
	assert.NoError(t, err)
	assert.Nil(t, artifacts)

	// not artifacts dir - ERROR
	u, _ = uri.Parse("dir://" + tmpFile)
	artifacts, err = Artifact(u)
	assert.Error(t, err)
	assert.Nil(t, artifacts)

	// not existing dir - ERROR
	u, _ = uri.Parse("dir://" + tmpDir + "/not-existing")
	artifacts, err = Artifact(u)
	assert.Error(t, err)
	assert.Nil(t, artifacts)
}
