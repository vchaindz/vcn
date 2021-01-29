/*
 * Copyright (c) 2018-2020 vChain, Inc. All Rights Reserved.
 * This software is released under GPL3.
 * The full license information can be found under:
 * https://www.gnu.org/licenses/gpl-3.0.en.html
 *
 */

package sign

import (
	"path/filepath"

	"github.com/vchain-us/vcn/pkg/store"

	"github.com/vchain-us/vcn/pkg/api"
	"github.com/vchain-us/vcn/pkg/bundle"
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

func (h *hook) finalize(v *api.BlockchainVerification, readOnly bool) error {
	if h != nil && !v.Unknown() {
		manifest, path := dir.Metadata(h.a)
		if manifest != nil && path != "" {
			// manifest is optional, we can ignore errors
			store.SaveManifest(h.a.Kind, path, *manifest)
			if !readOnly {
				bundle.WriteManifest(*manifest, filepath.Join(path, bundle.ManifestFilename))
			}
		}
	}
	return nil
}

func (h *hook) finalizeWithoutVerification(readOnly bool) error {
	if h != nil {
		manifest, path := dir.Metadata(h.a)
		if manifest != nil && path != "" {
			// manifest is optional, we can ignore errors
			store.SaveManifest(h.a.Kind, path, *manifest)
			if !readOnly {
				bundle.WriteManifest(*manifest, filepath.Join(path, bundle.ManifestFilename))
			}
		}
	}
	return nil
}
