/*
 * Copyright (c) 2018-2019 vChain, Inc. All Rights Reserved.
 * This software is released under GPL3.
 * The full license information can be found under:
 * https://www.gnu.org/licenses/gpl-3.0.en.html
 *
 */
package errors

import (
	"fmt"
	"testing"

	"github.com/vchain-us/vcn/pkg/meta"
)

func TestFormatErrorURLCustom(t *testing.T) {
	expectedURL := meta.ErrorWikiURL() + "password-404"

	actualURL := formatErrorURLCustom("password", 404)

	if expectedURL != actualURL {
		t.Error(fmt.Sprintf("formatErrorURLCustom() does not match [%s != %s]", expectedURL, actualURL))
	}
}

func TestFormatErrorURLByEndpoint(t *testing.T) {
	expectedURL := meta.ErrorWikiURL() + "publisher-post-412"

	res := meta.PublisherEndpoint()
	verb := "pOsT" // should do lowercase
	status := 412

	actualURL := formatErrorURLByEndpoint(res, verb, status)

	if expectedURL != actualURL {
		t.Error(fmt.Sprintf("formatErrorURLByEndpoint() does not match [%s != %s]", expectedURL, actualURL))
	}
}
