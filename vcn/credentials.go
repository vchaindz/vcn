package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

func ProvideKeystorePassword() (passphrase string, err error) {
	passphrase = os.Getenv(KeyStorePasswordEnv)
	if passphrase != "" {
		LOG.Trace("Keystore password provided (environment)")
		return passphrase, nil
	}
	passphrase, err = readPassword("Keystore passphrase: ")
	if err != nil {
		return "", err
	}
	LOG.Trace("Keystore password provided (interactive)")
	return passphrase, nil
}

func ProvidePlatformUsername() (user string, err error) {
	user = os.Getenv(VcnUserEnv)
	if user != "" {
		LOG.WithFields(logrus.Fields{
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
	LOG.WithFields(logrus.Fields{
		"username": user,
	}).Trace("Platform user provided (interactive)")
	return user, nil
}

func ProvidePlatformPassword() (password string, err error) {
	password = os.Getenv(VcnPasswordEnv)
	if password != "" {
		LOG.Trace("Platform password provided (environment)")
		return password, nil
	}
	password, err = readPassword("Password: ")
	if err != nil {
		return "", err
	}
	if password == "" {
		return "", fmt.Errorf("password must not be empty")
	}
	LOG.Trace("Platform password provided (interactive)")
	return password, nil
}
