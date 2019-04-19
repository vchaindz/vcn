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
	"log"
	"os"

	"github.com/dghubble/sling"
	"github.com/sirupsen/logrus"
	"github.com/vchain-us/vcn/internal/errors"
	"github.com/vchain-us/vcn/pkg/meta"
)

type Artifact struct {
	Name string
	Hash string
	Size uint64
}

type ArtifactRequest struct {
	Name       string `json:"name"`
	Hash       string `json:"hash"`
	Filename   string `json:"filename"`
	FileSize   uint64 `json:"fileSize"`
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
	Name                string `json:"name"`
	Hash                string `json:"hash"`
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

func CreateArtifact(verification *BlockchainVerification, walletAddress string,
	name string, hash string, fileSize uint64, visibility meta.Visibility, status meta.Status) error {
	restError := new(Error)
	token, err := LoadToken()
	if err != nil {
		fmt.Printf("\n%s\n", err.Error())
		errors.PrintErrorURLCustom("sign", 404)
		os.Exit(1)
	}
	metaHash := verification.HashAsset()
	r, err := sling.New().
		Post(meta.ArtifactEndpointForWallet(walletAddress)).
		Add("Authorization", "Bearer "+token).
		BodyJSON(ArtifactRequest{
			Name:       name,
			Hash:       hash,
			Filename:   name,
			FileSize:   fileSize,
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

func LoadArtifactsForCurrentWallet() ([]ArtifactResponse, error) {
	publicKey, err := PublicKeyForLocalWallet()
	if err != nil {
		return nil, err
	}
	return LoadArtifacts(publicKey)
}

func LoadArtifacts(walletAddress string) ([]ArtifactResponse, error) {
	response := new(PagedArtifactResponse)
	restError := new(Error)
	token, err := LoadToken()
	if err != nil {
		log.Fatal(err)
	}
	r, err := sling.New().
		Get(meta.ArtifactEndpointForWallet(walletAddress)).
		Add("Authorization", "Bearer "+token).
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

func LoadArtifactForHash(hash string, metahash string) (*ArtifactResponse, error) {
	response := new(ArtifactResponse)
	restError := new(Error)
	token, _ := LoadToken()
	r, err := newSling(token).
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
		return nil, fmt.Errorf("No asset matching hash %s/%s found", hash, metahash)
	}
	return nil, fmt.Errorf("Loading artifact for hash failed: %+v", restError)
}
