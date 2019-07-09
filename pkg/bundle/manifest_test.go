/*
 * Copyright (c) 2018-2019 vChain, Inc. All Rights Reserved.
 * This software is released under GPL3.
 * The full license information can be found under:
 * https://www.gnu.org/licenses/gpl-3.0.en.html
 *
 */

package bundle

import (
	"encoding/json"
	"io"
	"strings"

	"github.com/stretchr/testify/assert"

	"testing"
)

func TestManifest(t *testing.T) {

	items := make([]Descriptor, 0)
	for path, src := range map[string]io.Reader{
		"digits.txt":     strings.NewReader("1234567890"),
		"dup-digits.txt": strings.NewReader("1234567890"),
		"letters.txt":    strings.NewReader("abcdef"),
	} {
		d, err := NewDescriptor(path, src)
		assert.NoError(t, err)
		items = append(items, *d)
	}

	m := NewManifest(items...)

	assert.NotNil(t, m)

	d, err := m.Digest()
	assert.NoError(t, err)
	assert.Equal(t, "sha256:2cc48ce16beff9987ad31f32a7623bda48317d9232dbb75ba9a83f6d85ed073e", d.String())

	j, err := json.Marshal(m)
	assert.NoError(t, err)
	assert.Equal(
		t,
		`{"schemaVersion":1,"items":[{"digest":"sha256:bef57ec7f53a6d40beb640a780a639c83bc29ac8a9816f1fc6c5c6dcd93c4721","size":6,"paths":["letters.txt"]},{"digest":"sha256:c775e7b757ede630cd0aa1113bd102661ab38829ca52a6422ab782862f268646","size":10,"paths":["digits.txt","dup-digits.txt"]}]}`,
		string(j),
	)

}
