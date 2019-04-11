package main

import (
	"io/ioutil"
	"os"

	"github.com/sirupsen/logrus"
)

func LoadToken() (jwtToken string, err error) {
	LOG.WithFields(logrus.Fields{
		"tokenFile": TokenFile(),
	}).Trace("LoadToken")
	contents, err := ioutil.ReadFile(TokenFile())
	if err != nil {
		return "", err
	}
	return string(contents), nil
}

func WriteToken(token string) (err error) {
	LOG.WithFields(logrus.Fields{
		"tokenFile": TokenFile(),
	}).Trace("WriteToken")
	return ioutil.WriteFile(TokenFile(), []byte(token), os.FileMode(0600))
}
