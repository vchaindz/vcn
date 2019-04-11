/*
 * Copyright (c) 2018-2019 vChain, Inc. All Rights Reserved.
 * This software is released under GPL3.
 * The full license information can be found under:
 * https://www.gnu.org/licenses/gpl-3.0.en.html
 */

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/dghubble/sling"
	"github.com/sirupsen/logrus"
)

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type TokenResponse struct {
	Token string `token:"token"`
}

type Error struct {
	Status    int    `json:"status"`
	Message   string `json:"message"`
	Path      string `json:"path"`
	Timestamp string `json:"timestamp"`
	Error     string `json:"error"`
}

type PublisherExistsResponse struct {
	Exists bool `json:"exists"`
}
type PublisherExistsParams struct {
	Email string `url:"email"`
}
type PublisherResponse struct {
	Authorities []string `json:"authorities"`
	Email       string   `json:"email"`
	FirstName   string   `json:"firstName"`
	LastName    string   `json:"lastName"`
}

func CheckPublisherExists(email string) (ret bool) {

	email = strings.TrimSpace(email)

	params := &PublisherExistsParams{Email: email}
	response := new(PublisherExistsResponse)
	restError := new(Error)

	r, err := sling.New().
		Get(PublisherEndpoint()+"/exists").
		QueryStruct(params).
		Receive(&response, restError)

	if err != nil {
		fmt.Printf(err.Error())
		return false
	}
	if r.StatusCode != 200 {

		fmt.Printf(fmt.Sprintf("request failed: %s (%d)",
			restError.Message, restError.Status))
		return false
	}

	return response.Exists
}

func CheckToken(token string) (ret bool, err error) {
	restError := new(Error)
	r, err := NewSling(token).
		Get(TokenCheckEndpoint()).
		Receive(nil, restError)
	LOG.WithFields(logrus.Fields{
		"err":       err,
		"restError": restError,
	}).Trace("CheckToken")
	if err != nil {
		return false, err
	}
	if r.StatusCode == 200 {
		return true, nil
	} else {
		return false, fmt.Errorf("authentication failed: %+v", restError)
	}
}

func Authenticate(email string, password string) (ret bool, code int) { // TODO: rework

	if password == "" {
		return false, 401
	}
	token := new(TokenResponse)
	authError := new(Error)

	r, err := sling.New().
		Post(PublisherEndpoint()+"/auth").
		BodyJSON(AuthRequest{Email: email, Password: password}).
		Receive(token, authError)
	if err != nil {

		log.Fatal(err)
	}
	if r.StatusCode != 200 {

		LOG.WithFields(logrus.Fields{
			"code":  r.StatusCode,
			"error": authError.Message,
		}).Trace("API request failed")

		return false, authError.Status

	}
	err = ioutil.WriteFile(TokenFile(), []byte(token.Token),
		os.FileMode(0600))
	if err != nil {

		log.Fatal(err)
	}

	return true, 0

}

func LoadToken() (jwtToken string, err error) {

	LOG.WithFields(logrus.Fields{
		"tokenFile": TokenFile(),
	}).Trace("Access local token")

	contents, err := ioutil.ReadFile(TokenFile())
	if err != nil {
		return "", err
	}
	return string(contents), nil
}
