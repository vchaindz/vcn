package api

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/vchain-us/vcn/pkg/meta"
)

func LoadToken() (jwtToken string, err error) {
	logger().WithFields(logrus.Fields{
		"tokenFile": meta.TokenFile(),
	}).Trace("LoadToken")
	contents, err := ioutil.ReadFile(meta.TokenFile())
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(contents)), nil
}

func WriteToken(token string) (err error) {
	logger().WithFields(logrus.Fields{
		"tokenFile": meta.TokenFile(),
	}).Trace("WriteToken")
	return ioutil.WriteFile(meta.TokenFile(), []byte(token), os.FileMode(0600))
}
