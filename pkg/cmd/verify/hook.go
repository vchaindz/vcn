/*
 * Copyright (c) 2018-2019 vChain, Inc. All Rights Reserved.
 * This software is released under GPL3.
 * The full license information can be found under:
 * https://www.gnu.org/licenses/gpl-3.0.en.html
 *
 */

package verify

import (
	"fmt"
	"path/filepath"

	"github.com/vchain-us/vcn/pkg/bundle"

	"github.com/vchain-us/vcn/pkg/api"
	"github.com/vchain-us/vcn/pkg/extractor/dir"
)

type hook struct {
	a api.Artifact
}

func newHook(a *api.Artifact) *hook {
	if a != nil {
		h := hook{
			a: a.Copy(),
		}
		dir.RemoveMetadata(a)
		return &h
	}
	return nil
}

func (h *hook) finalize(v *api.BlockchainVerification, output string) error {
	if h != nil && output == "" {
		manifest, path := dir.Metadata(h.a)
		if manifest != nil && path != "" {
			oldManifest, err := bundle.ReadManifest(filepath.Join(path, bundle.ManifestFilename))
			if err != nil {
				return nil // continue if missing or bad manifest
			}
			report, equal, err := manifest.DiffByPath(*oldManifest)
			if err != nil {
				return err
			}
			if !equal {
				fmt.Printf("Diff\n%s\n", report)
			}

		}
	}
	return nil
}
