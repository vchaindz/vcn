/*
 * Copyright (c) 2018-2019 vChain, Inc. All Rights Reserved.
 * This software is released under GPL3.
 * The full license information can be found under:
 * https://www.gnu.org/licenses/gpl-3.0.en.html
 *
 */

package dir

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/src-d/go-git.v4/plumbing/format/gitignore"
)

// IgnoreFilename is the name of the ignore file used by this package when processing directories.
//
// The ignore file supports all pattern formats as specified in the gitignore specification:
//  - https://git-scm.com/docs/gitignore
//  - https://github.com/src-d/go-git/blob/master/plumbing/format/gitignore/doc.go
//
// However, this package implementation:
//  - only uses the ignore file within the root directory (nested ignore files will be treated as normal files)
//  - always ignores the manifest file (it cannot be excluded by the ignore file)
//
const IgnoreFilename = ".vcnignore"

const (
	ignorefileCommentPrefix = "#"
	ignorefileEOL           = "\n"
)

type nullMatcher struct{}

func (n *nullMatcher) Match(path []string, isDir bool) bool {
	return false
}

// newIgnoreFileMatcher reads and parses the ignore file in path and return a gitignore.Matcher.
// If the ignore file does not exists, a matcher that will never match is returned instead.
func newIgnoreFileMatcher(path string) (m gitignore.Matcher, err error) {
	f, err := os.Open(filepath.Join(path, IgnoreFilename))
	if err != nil {
		if os.IsNotExist(err) {
			return &nullMatcher{}, nil
		}
		return
	}
	defer f.Close()
	ps := []gitignore.Pattern{}
	if data, err := ioutil.ReadAll(f); err == nil {
		for _, s := range strings.Split(string(data), ignorefileEOL) {
			if !strings.HasPrefix(s, ignorefileCommentPrefix) && len(strings.TrimSpace(s)) > 0 {
				ps = append(ps, gitignore.ParsePattern(s, nil))
			}
		}
	}
	m = gitignore.NewMatcher(ps)
	return
}
