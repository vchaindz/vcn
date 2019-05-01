/*
 * Copyright (c) 2018-2019 vChain, Inc. All Rights Reserved.
 * This software is released under GPL3.
 * The full license information can be found under:
 * https://www.gnu.org/licenses/gpl-3.0.en.html
 *
 */

package api

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/vchain-us/vcn/pkg/meta"
)

type Artifact struct {
	Kind     string
	Name     string
	Hash     string
	Size     uint64
	MimeType string
}

type ArtifactRequest struct {
	Kind     string `json:"kind"`
	Name     string `json:"name"`
	Hash     string `json:"hash"`
	MimeType string `json:"mimeType"`

	FileSize   uint64 `json:"fileSize"`
	Filename   string `json:"filename"`
	Url        string `json:"url"`
	License    string `json:"license"`
	Visibility string `json:"visibility"`
	Status     string `json:"status"`
	MetaHash   string `json:"metaHash"`
}

type PagedArtifactResponse struct {
	Content []ArtifactResponse `json:"content"`
}

type ArtifactResponse struct {
	Kind     string `json:"kind"`
	Name     string `json:"name"`
	Hash     string `json:"hash"`
	MimeType string `json:"mimeType"`

	Filename            string `json:"filename"`
	FileSize            uint64 `json:"fileSize"`
	Url                 string `json:"url"`
	License             string `json:"license"`
	Level               int64  `json:"level"`
	Visibility          string `json:"visibility"`
	Status              string `json:"status"`
	Publisher           string `json:"publisher"`
	CountVerifications  uint64 `json:"verificationCount"`
	CountConflicts      uint64 `json:"publisherCount"`
	CreatedAt           string `json:"createdAt"`
	PublisherCompany    string `json:"publisherCompany"`
	PublisherWebsiteUrl string `json:"publisherWebsiteUrl"`
}

func (a ArtifactResponse) String() string {
	return fmt.Sprintf("File:\t%s\nHash:\t%s\nStatus:\t%s\n\n",
		a.Name, a.Hash, a.Status)
}

func (u User) createArtifact(verification *BlockchainVerification, walletAddress string,
	artifact Artifact, visibility meta.Visibility, status meta.Status) error {

	hasAuth, err := u.IsAuthenticated()
	if err != nil {
		return err
	}
	if !hasAuth {
		return makeAuthRequiredError()
	}

	metaHash := verification.MetaHash()
	restError := new(Error)
	r, err := newSling(u.token()).
		Post(meta.ArtifactEndpointForWallet(walletAddress)).
		BodyJSON(ArtifactRequest{
			Kind:     artifact.Kind,
			Name:     artifact.Name,
			Hash:     artifact.Hash,
			MimeType: artifact.MimeType,

			Filename:   artifact.Name,
			FileSize:   artifact.Size,
			Visibility: meta.VisibilityName(visibility),
			Status:     meta.StatusName(status),
			MetaHash:   metaHash,
		}).Receive(nil, restError)
	if err != nil {
		return err
	}
	if r.StatusCode != 200 {
		return fmt.Errorf("request failed: %s (%d)", restError.Message,
			restError.Status)
	}
	return nil
}

func (u *User) LoadAllArtifacts() ([]ArtifactResponse, error) {
	ret := []ArtifactResponse{}
	for _, pubKey := range u.cfg.PubKeys() {
		chunk, err := u.LoadArtifacts(pubKey)
		if err != nil {
			return nil, err
		}
		ret = append(ret, chunk...)
	}
	return ret, nil
}

func (u *User) LoadArtifacts(walletAddress string) ([]ArtifactResponse, error) {
	response := new(PagedArtifactResponse)
	restError := new(Error)
	r, err := newSling(u.token()).
		Get(meta.ArtifactEndpointForWallet(walletAddress)).
		Receive(&response, restError)
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		return nil, fmt.Errorf("request failed: %s (%d)",
			restError.Message, restError.Status)
	}
	return response.Content, nil
}

func LoadArtifactForHash(user *User, hash string, metahash string) (*ArtifactResponse, error) {
	response := new(ArtifactResponse)
	restError := new(Error)
	r, err := newSling(user.token()).
		Get(meta.ArtifactEndpoint()+"/"+hash+"/"+metahash).
		Receive(&response, restError)
	logger().WithFields(logrus.Fields{
		"response":  response,
		"err":       err,
		"restError": restError,
	}).Trace("LoadArtifactForHash")
	if err != nil {
		return nil, err
	}
	switch r.StatusCode {
	case 200:
		return response, nil
	case 404:
		return nil, fmt.Errorf("no asset matching hash %s/%s found", hash, metahash)
	}
	return nil, fmt.Errorf("loading artifact for hash failed: %+v", restError)
}
