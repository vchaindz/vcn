/*
 * Copyright (c) 2018-2020 vChain, Inc. All Rights Reserved.
 * This software is released under GPL3.
 * The full license information can be found under:
 * https://www.gnu.org/licenses/gpl-3.0.en.html
 *
 */

package api

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"github.com/vchain-us/vcn/pkg/meta"
)

func (a Artifact) toLcArtifact() *LcArtifact {
	aR := &LcArtifact{
		// root fields
		Kind:        a.Kind,
		Name:        a.Name,
		Hash:        a.Hash,
		Size:        a.Size,
		ContentType: a.ContentType,

		// custom metadata
		Metadata: a.Metadata,
	}

	return aR
}

type LcArtifact struct {
	// root fields
	Kind        string `json:"kind" yaml:"kind" vcn:"Kind"`
	Name        string `json:"name" yaml:"name" vcn:"Name"`
	Hash        string `json:"hash" yaml:"hash" vcn:"Hash"`
	Size        uint64 `json:"size" yaml:"size" vcn:"Size"`
	ContentType string `json:"contentType" yaml:"contentType" vcn:"ContentType"`

	// custom metadata
	Metadata Metadata `json:"metadata" yaml:"metadata" vcn:"Metadata"`

	Signer string      `json:"signer" yaml:"signer" vcn:"Signer"`
	Status meta.Status `json:"status,omitempty" yaml:"status,omitempty" vcn:"Status"`
}

func (u LcUser) createArtifact(
	artifact Artifact, status meta.Status) error {

	aR := artifact.toLcArtifact()
	aR.Status = status

	hasher := sha256.New()
	hasher.Write([]byte(u.LcApiKey()))
	aR.Signer = base64.URLEncoding.EncodeToString(hasher.Sum(nil))

	arJson, err := json.Marshal(aR)
	_, err = u.Client.SafeSet(context.TODO(), []byte(artifact.Hash), arJson)
	if err != nil {
		return err
	}
	return nil
}

// LoadArtifact fetches and returns an *lcArtifact for the given hash and current u, if any.
func (u *LcUser) LoadArtifact(hash string) (lc *LcArtifact, verified bool, err error) {
	jsonAr, err := u.Client.SafeGet(context.TODO(), []byte(hash))
	if err != nil {
		return nil, false, err
	}
	var lcArtifact LcArtifact

	err = json.Unmarshal(jsonAr.Value, &lcArtifact)
	if err != nil {
		return nil, false, err
	}
	return &lcArtifact, jsonAr.Verified, nil
}
