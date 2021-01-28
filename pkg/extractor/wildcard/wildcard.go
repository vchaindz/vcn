/*
 * Copyright (c) 2018-2020 vChain, Inc. All Rights Reserved.
 * This software is released under GPL3.
 * The full license information can be found under:
 * https://www.gnu.org/licenses/gpl-3.0.en.html
 *
 */

package wildcard

import (
	"errors"
	"github.com/vchain-us/vcn/pkg/extractor/dir"
	"github.com/vchain-us/vcn/pkg/extractor/file"
	"os"
	"path/filepath"
	"strings"

	"github.com/vchain-us/vcn/pkg/api"
	"github.com/vchain-us/vcn/pkg/extractor"
	"github.com/vchain-us/vcn/pkg/uri"
)

// Scheme for dir
const Scheme = "wildcard"

// ManifestKey is the metadata's key for storing the manifest
const ManifestKey = "manifest"

// PathKey is the metadata's key for the directory path
const PathKey = "path"

type opts struct {
	initIgnoreFile    bool
	skipIgnoreFileErr bool
	recursive         bool
}

// Artifact returns a file *api.Artifact from a given u
func Artifact(u *uri.URI, options ...extractor.Option) ([]*api.Artifact, error) {

	if u.Scheme != "" && u.Scheme != Scheme {
		return nil, nil
	}

	opts := &opts{}
	if err := extractor.Options(options).Apply(opts); err != nil {
		return nil, err
	}

	path := strings.TrimPrefix(u.Opaque, "//")
	wildcard := filepath.Base(path)
	p, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	// provided path is a file
	if fileInfo, err := os.Stat(p); err == nil {
		if !fileInfo.IsDir() {
			u, err := uri.Parse("file://" + p)
			if err != nil {
				return nil, err
			}
			return file.Artifact(u)
		}
		u, err := uri.Parse("dir://" + p)
		if err != nil {
			return nil, err
		}
		return dir.Artifact(u)
	}

	root := filepath.Dir(p)

	if opts.initIgnoreFile {
		if err := dir.InitIgnoreFile(path); err != nil {
			if !opts.skipIgnoreFileErr {
				return nil, err
			}
		}
	}

	// build a list of all files matching the wildcard provided. Method is based on filepath.Glob
	var filePaths []string
	if opts.recursive {
		err = filepath.Walk(root, buildFilePaths(wildcard, &filePaths))
		if err != nil {
			return nil, err
		}
	} else {
		i, err := os.Stat(root)
		if err != nil {
			return nil, err
		}
		err = buildFilePaths(wildcard, &filePaths)(root, i, nil)
		if err != nil {
			return nil, err
		}
	}

	if len(filePaths) == 0 {
		return nil, errors.New("no files matching from provided search terms")
	}

	arst := []*api.Artifact{}
	// convert files path list to artifacts
	for _, fp := range filePaths {
		u, err := uri.Parse("file://" + fp)
		if err != nil {
			return nil, err
		}
		ars, err := file.Artifact(u)
		if err != nil {
			return nil, err
		}
		arst = append(arst, ars...)
	}

	return arst, nil
}

func buildFilePaths(wildcard string, filePaths *[]string) func(ele string, info os.FileInfo, err error) error {
	return func(ele string, info os.FileInfo, err error) error {
		if info.IsDir() {
			fpd, err := filepath.Glob(filepath.Join(ele, wildcard))
			if err != nil {
				return err
			}
			if len(fpd) > 0 {
				for _, fp := range fpd {
					info, err = os.Stat(fp)
					if err != nil {
						return err
					}
					if !info.IsDir() {
						*filePaths = append(*filePaths, fp)
					}
				}
			}
		}
		return nil
	}
}

// WithRecursive wildcard usage will walk inside subdirectories of provided path
func WithRecursive() extractor.Option {
	return func(o interface{}) error {
		if o, ok := o.(*opts); ok {
			o.recursive = true
		}
		return nil
	}
}
