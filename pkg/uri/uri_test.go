/*
 * Copyright (c) 2018-2020 vChain, Inc. All Rights Reserved.
 * This software is released under GPL3.
 * The full license information can be found under:
 * https://www.gnu.org/licenses/gpl-3.0.en.html
 *
 */

package uri

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestURI(t *testing.T) {
	u, err := Parse("scheme://opaque")
	assert.NoError(t, err)
	assert.NotNil(t, u)
	assert.Equal(t, URI{Scheme: "scheme", Opaque: "//opaque"}, *u)
	assert.Equal(t, "scheme://opaque", u.String())

	u, err = Parse("file.txt")
	assert.NoError(t, err)
	assert.NotNil(t, u)
	assert.Equal(t, URI{Scheme: "", Opaque: "file.txt"}, *u)
	assert.Equal(t, "file.txt", u.String())
}
