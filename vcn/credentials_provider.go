package main

import (
	"fmt"
	"os"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

func ProvideKeystorePassword() (passphrase string, err error) {
	passphrase = os.Getenv("KEYSTORE_PASSWORD")
	if passphrase != "" {
		return passphrase, nil
	}
	fmt.Print("Keystore passphrase:")
	passphraseBytes, err := terminal.ReadPassword(int(syscall.Stdin))
	fmt.Println(".")
	if err != nil {
		return "", nil
	}
	return string(passphraseBytes), nil
}
