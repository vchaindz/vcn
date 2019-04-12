/*
 * Copyright (c) 2018-2019 vChain, Inc. All Rights Reserved.
 * This software is released under GPL3.
 * The full license information can be found under:
 * https://www.gnu.org/licenses/gpl-3.0.en.html
 *
 */

package main

import (
	"fmt"
	"testing"
)

func TestErrorURLComposition(t *testing.T) {

	expectedURL := "https://github.com/vchain-us/vcn/wiki/Errors#publisher-post-412"

	res := PublisherEndpoint()
	verb := "pOsT" // should do lowercase
	status := 412

	actualURL := formatErrorURLByEndpoint(res, verb, status)

	if expectedURL != actualURL {
		t.Error(fmt.Sprintf("formatErrorURLByEndpoint() does not match [%s != %s]", expectedURL, actualURL))
	}

}

func TestHash(t *testing.T) {
	if hash("../resources/testHash.example") != "181210f8f9c779c26da1d9b2075bde0127302ee0e3fca38c9a83f5b1dd8e5d3b" {
		t.Error(`hash("../resources/testHash.example") does not match`)
	}
}
