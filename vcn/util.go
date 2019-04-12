/*
 * Copyright (c) 2018-2019 vChain, Inc. All Rights Reserved.
 * This software is released under GPL3.
 * The full license information can be found under:
 * https://www.gnu.org/licenses/gpl-3.0.en.html
 *
 */

package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"syscall"

	"github.com/dghubble/sling"
	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh/terminal"
)

func firstFile(dir string) (io.Reader, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	for _, f := range files {
		return os.Open(dir + "/" + f.Name())
	}
	return nil, fmt.Errorf("empty directory: %s", dir)
}

func contains(xs []string, x string) bool {
	for _, a := range xs {
		if a == x {
			return true
		}
	}
	return false
}

func readPassword(msg string) (string, error) {
	fmt.Print(msg)
	password, err := terminal.ReadPassword(int(syscall.Stdin))
	fmt.Println(".")
	if err != nil {
		return "", err
	}
	return string(password), nil
}

func CreateVcnDirectories() {
	if err := os.MkdirAll(WalletDirectory(),
		os.FileMode(VcnDirectoryPermissions)); err != nil {
		log.Fatal(err)
	}
}

// PrintErrorURLCustom takes custom domain and status code
func PrintErrorURLCustom(domain string, code int) {

	fmt.Print("Get help for this error at:\n")

	color.Set(StyleError())
	fmt.Print(formatErrorURLCustom(domain, code))
	color.Unset()

	fmt.Println()
	return

}
func formatErrorURLCustom(domain string, status int) string {

	errorPage := ErrorWikiURL()

	return fmt.Sprintf("%s%s-%d", errorPage, domain, status)

}

func formatErrorURLByEndpoint(resource string, verb string, status int) string {

	errorPage := ErrorWikiURL()

	// get last part of endpoint
	x := strings.Split(resource, "/")
	resource = x[len(x)-1]

	return fmt.Sprintf("%s%s-%s-%d", errorPage, resource, strings.ToLower(verb), status)

}

func hashAsset(verification *BlockchainVerification) string {
	metadata := fmt.Sprintf("%s-%d-%d-%d",
		verification.Owner.Hex(),
		int64(verification.Level),
		int64(verification.Status),
		int64(verification.Timestamp.Unix()))
	metadataHashAsBytes := sha256.Sum256([]byte(metadata))
	LOG.WithFields(logrus.Fields{
		"metahash": metadata,
	}).Trace("Generated metahash")
	return fmt.Sprintf("%x", metadataHashAsBytes)
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

func NewSling(token string) (s *sling.Sling) {
	s = sling.New()
	if token != "" {
		s = s.Add("Authorization", "Bearer "+token)
	}
	return s
}
