/*
 * Copyright (c) 2018-2020 vChain, Inc. All Rights Reserved.
 * This software is released under GPL3.
 * The full license information can be found under:
 * https://www.gnu.org/licenses/gpl-3.0.en.html
 *
 */

package sign

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapOpts(t *testing.T) {
	m := mapOpts{}

	err := m.Set("key=value")
	assert.NoError(t, err)

	assert.Equal(t, mapOpts{"key": "value"}, m)
	assert.Equal(t, `{"key":"value"}`, m.String())
	assert.Equal(t, map[string]interface{}{"key": "value"}, m.StringToInterface())
}
