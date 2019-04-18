/*
 * Copyright (c) 2018-2019 vChain, Inc. All Rights Reserved.
 * This software is released under GPL3.
 * The full license information can be found under:
 * https://www.gnu.org/licenses/gpl-3.0.en.html
 *
 */

package cli

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"syscall"

	"github.com/fatih/color"
	"golang.org/x/crypto/ssh/terminal"
)

func readPassword(msg string) (string, error) {
	fmt.Print(msg)
	password, err := terminal.ReadPassword(int(syscall.Stdin))
	fmt.Println(".")
	if err != nil {
		return "", err
	}
	return string(password), nil
}

func hash(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	h := sha256.New()
	if _, err := io.Copy(h, file); err != nil {
		log.Fatal(err)
	}
	checksum := h.Sum(nil)
	return hex.EncodeToString(checksum)
}

func printColumn(field string, value string, fallback string, p ...color.Attribute) {
	var spaces string
	for i := len(field); i < 8; i++ {
		spaces += " "
	}
	fmt.Print(field + ":" + spaces)
	if p != nil {
		c := color.New(p...)
		c.Set()
	}
	if value != "" {
		fmt.Print(value)
	} else {
		fmt.Print(fallback)
	}
	if p != nil {
		color.Unset()
	}
	fmt.Println()
}