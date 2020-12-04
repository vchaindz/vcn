package wildcard

import (
	file2 "github.com/vchain-us/vcn/pkg/extractor/file"
	"log"

	"github.com/stretchr/testify/assert"
	"github.com/vchain-us/vcn/pkg/uri"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestWildcard(t *testing.T) {
	file, err := ioutil.TempFile("", "vcn-test-scheme-file")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(file.Name())
	err = ioutil.WriteFile(file.Name(), []byte("123\n"), 0644)
	if err != nil {
		log.Fatal(err)
	}

	// Default empty schema is wildcard
	u, _ := uri.Parse(file.Name())
	artifacts, err := Artifact(u)
	assert.NoError(t, err)
	assert.NotNil(t, artifacts[0])
	assert.Equal(t, file2.Scheme, artifacts[0].Kind)
	assert.Equal(t, filepath.Base(file.Name()), artifacts[0].Name)
	assert.Equal(t, "181210f8f9c779c26da1d9b2075bde0127302ee0e3fca38c9a83f5b1dd8e5d3b", artifacts[0].Hash)

	u, _ = uri.Parse("../../../docs/vcncheatsheet.pdf")
	artifacts, err = Artifact(u)
	assert.NoError(t, err)
	assert.NotNil(t, artifacts)
	assert.Equal(t, artifacts[0].ContentType, "application/pdf")
}
