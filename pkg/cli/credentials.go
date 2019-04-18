package cli

import (
	"fmt"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/vchain-us/vcn/pkg/logs"
	"github.com/vchain-us/vcn/pkg/meta"
)

func ProvideKeystorePassword() (passphrase string, err error) {
	passphrase = os.Getenv(meta.KeyStorePasswordEnv)
	if passphrase != "" {
		logs.LOG.Trace("Keystore password provided (environment)")
		return passphrase, nil
	}
	passphrase, err = readPassword("Keystore passphrase: ")
	if err != nil {
		return "", err
	}
	logs.LOG.Trace("Keystore password provided (interactive)")
	return passphrase, nil
}

func ProvidePlatformUsername() (user string, err error) {
	user = os.Getenv(meta.VcnUserEnv)
	if user != "" {
		logs.LOG.WithFields(logrus.Fields{
			"username": user,
		}).Trace("Platform user provided (environment)")
		return user, nil
	}
	fmt.Print("Email address: ")
	cnt, err := fmt.Scanln(&user)
	if err != nil {
		return "", err
	}
	if cnt <= 0 {
		return "", fmt.Errorf("username must not be empty")
	}
	user = strings.TrimSpace(user)
	logs.LOG.WithFields(logrus.Fields{
		"username": user,
	}).Trace("Platform user provided (interactive)")
	return user, nil
}

func ProvidePlatformPassword() (password string, err error) {
	password = os.Getenv(meta.VcnPasswordEnv)
	if password != "" {
		logs.LOG.Trace("Platform password provided (environment)")
		return password, nil
	}
	password, err = readPassword("Password: ")
	if err != nil {
		return "", err
	}
	if password == "" {
		return "", fmt.Errorf("password must not be empty")
	}
	logs.LOG.Trace("Platform password provided (interactive)")
	return password, nil
}
