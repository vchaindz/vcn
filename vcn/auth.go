/*
 * Copyright (c) 2018-2019 vChain, Inc. All Rights Reserved.
 * This software is released under GPL3.
 * The full license information can be found under:
 * https://www.gnu.org/licenses/gpl-3.0.en.html
 */

package main

import (
	"fmt"

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

func CheckPublisherExists(email string) (success bool, err error) {
	response := new(PublisherExistsResponse)
	restError := new(Error)
	r, err := sling.New().
		Get(PublisherEndpoint()+"/exists").
		QueryStruct(&PublisherExistsParams{Email: email}).
		Receive(&response, restError)
	LOG.WithFields(logrus.Fields{
		"response":  response,
		"err":       err,
		"restError": restError,
	}).Trace("CheckPublisherExists")
	if err != nil {
		return false, err
	}
	if r.StatusCode != 200 {
		return false, fmt.Errorf("check publisher failed: %+v", restError)
	}
	return response.Exists, nil
}

func CheckToken(token string) (ret bool, err error) {
	restError := new(Error)
	response, err := NewSling(token).
		Get(TokenCheckEndpoint()).
		Receive(nil, restError)
	LOG.WithFields(logrus.Fields{
		"response":  response,
		"err":       err,
		"restError": restError,
	}).Trace("CheckToken")
	if err != nil {
		return false, err
	}
	if response.StatusCode == 200 {
		return true, nil
	} else {
		return false, fmt.Errorf("authentication failed: %+v", restError)
	}
}

func Authenticate(email string, password string) (err error) { // TODO: rework
	response := new(TokenResponse)
	restError := new(Error)
	r, err := sling.New().
		Post(PublisherEndpoint()+"/auth").
		BodyJSON(AuthRequest{Email: email, Password: password}).
		Receive(response, restError)
	LOG.WithFields(logrus.Fields{
		"email":     email,
		"response":  response,
		"err":       err,
		"restError": restError,
	}).Trace("Authenticate")
	if err != nil {
		return err
	}
	switch r.StatusCode {
	case 200:
		return WriteToken(response.Token)
	case 400:
		return fmt.Errorf("Your email address was not confirmed. " +
			"Please confirm it by clicking on the link we sent to " + email + ". " +
			"If you did not receive the email, please go to dashboard.codenotary.io and click on the link \"Resend email\"")
	case 401:
		return fmt.Errorf("invalid password")
	}
	return fmt.Errorf("unhandled authentication error: %+v", restError)
}
