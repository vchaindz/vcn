/*
 * Copyright (c) 2018-2019 vChain, Inc. All Rights Reserved.
 * This software is released under GPL3.
 * The full license information can be found under:
 * https://www.gnu.org/licenses/gpl-3.0.en.html
 *
 */

package api

import (
	"github.com/dghubble/sling"
)

func contains(xs []string, x string) bool {
	for _, a := range xs {
		if a == x {
			return true
		}
	}
	return false
}

func newSling(token string) (s *sling.Sling) {
	s = sling.New()
	if token != "" {
		s = s.Add("Authorization", "Bearer "+token)
	}
	return s
}
