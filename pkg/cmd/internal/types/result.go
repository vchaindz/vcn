/*
 * Copyright (c) 2018-2019 vChain, Inc. All Rights Reserved.
 * This software is released under GPL3.
 * The full license information can be found under:
 * https://www.gnu.org/licenses/gpl-3.0.en.html
 *
 */

package types

import (
	"github.com/vchain-us/vcn/pkg/api"
)

type Result struct {
	Artifact     *api.ArtifactResponse       `json:"artifact" yaml:"artifact"`
	Verification *api.BlockchainVerification `json:"verification" yaml:"verification"`
	Errors       []error                     `json:"error,omitempty" yaml:"error,omitempty"`
}

func (r *Result) AddError(err error) {
	r.Errors = append(r.Errors, err)
}

func NewResult(a *api.Artifact, ar *api.ArtifactResponse, v *api.BlockchainVerification) *Result {
	r := Result{}
	if v != nil {
		vCopy := *v
		r.Verification = &vCopy
	}
	if ar != nil {
		arc := *ar
		r.Artifact = &arc
	} else if a != nil {
		r.Artifact = &api.ArtifactResponse{
			Name:        a.Name,
			Kind:        a.Kind,
			Hash:        a.Hash,
			Size:        a.Size,
			ContentType: a.ContentType,
			Metadata:    a.Metadata,
		}
	}
	// Do not show status and level from platform
	if r.Artifact != nil {
		r.Artifact.Status = ""
		r.Artifact.Level = 0
	}
	return &r
}
