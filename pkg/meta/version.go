/*
 * Copyright (c) 2018-2020 vChain, Inc. All Rights Reserved.
 * This software is released under GPL3.
 * The full license information can be found under:
 * https://www.gnu.org/licenses/gpl-3.0.en.html
 *
 */

package meta

var version = ""

var static = ""

var gitCommit = ""
var gitBranch = ""

// Version returns the current CodeNotary vcn version string
func Version() string {
	return version
}

// StaticBuild returns when the current vcn executable has been statically linked against libraries
func StaticBuild() bool {
	return static == "static"
}

// GitRevision returns the current CodeNotary vcn git revision string
func GitRevision() string {
	rev := gitCommit
	if gitBranch != "" {
		rev += " (" + gitBranch + ")"
	}
	return rev
}
