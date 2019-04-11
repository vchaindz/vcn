package main

import (
	"fmt"
	"os"
	"strings"
	"syscall"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh/terminal"
)

func ProvideKeystorePassword() (passphrase string, err error) {
	passphrase = os.Getenv("KEYSTORE_PASSWORD")
	if passphrase != "" {
		LOG.Trace("Keystore password provided (environment)")
		return passphrase, nil
	}
	fmt.Print("Keystore passphrase: ")
	passphraseBytes, err := terminal.ReadPassword(int(syscall.Stdin))
	fmt.Println(".")
	if err != nil {
		return "", nil
	}
	LOG.Trace("Keystore password provided (interactive)")
	return string(passphraseBytes), nil
}

func ProvidePlatformUsername() (user string, err error) {
	user = os.Getenv("VCN_USER")
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
	password = os.Getenv("VCN_PASSWORD")
	if password != "" {
		LOG.Trace("Platform password provided (environment)")
		return password, nil
	}
	fmt.Print("Password: ")
	passphraseBytes, err := terminal.ReadPassword(int(syscall.Stdin))
	fmt.Println(".")
	if err != nil {
		return "", nil
	}
	password = string(passphraseBytes)
	LOG.Trace("Platform password provided (interactive)")
	if password == "" {
		return "", fmt.Errorf("password must not be empty")
	}
	return password, nil
}
