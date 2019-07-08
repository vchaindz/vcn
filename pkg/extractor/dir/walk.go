/*
 * Copyright (c) 2018-2019 vChain, Inc. All Rights Reserved.
 * This software is released under GPL3.
 * The full license information can be found under:
 * https://www.gnu.org/licenses/gpl-3.0.en.html
 *
 */

package dir

import (
	"os"
	"path/filepath"

	"github.com/vchain-us/vcn/pkg/bundle"
)

func walk(root string) (files []bundle.Descriptor, err error) {
	files = make([]bundle.Descriptor, 0)
	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		// ignore dirs
		if info.IsDir() {
			return nil
		}

		relPath, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}

		// ignore  manifest file
		if relPath == bundle.ManifestFilename {
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		d, err := bundle.NewDescriptor(filepath.ToSlash(relPath), file)
		file.Close()
		if err != nil {
			return err
		}
		files = append(files, *d)

		return nil
	})
	return
}
