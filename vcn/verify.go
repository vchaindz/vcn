/*
 * Copyright (c) 2018-2019 vChain, Inc. All Rights Reserved.
 * This software is released under GPL3.
 * The full license information can be found under:
 * https://www.gnu.org/licenses/gpl-3.0.en.html
 *
 * Built on top of CLI (MIT license)
 * https://github.com/urfave/cli#overview
 */

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
)

type response struct {
	Message string `json:"message"`
}

func verifyAll(files []string) {
	// find . -type f -name "*.go" | xargs -I % vcn verify %
	//for --> verify
	//fmt.Println(files)
	for i := 0; i < len(files); i++ {
		//fmt.Println(i)
		verify(files[i])
	}
}

func verify(filename string) {

	hash := hash(filename)

	vcnClient := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
	}

	req, err := http.NewRequest(http.MethodGet, ApiEndpoint("files/"+hash), nil)
	if err != nil {
		log.Fatal(err)
	}

	res, getErr := vcnClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	verification := response{}
	jsonErr := json.Unmarshal(body, &verification)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	// TO DO: Trust level= ?? mb, also this should also go to STDOUT,
	// and also handle STDERR so it can be fully scripted

	trust := strings.TrimSpace(verification.Message)

	// Asset level
	fmt.Println(" File:", filename)

	fmt.Println()

	// Blockchain level
	fmt.Println(" Hash:", strings.TrimSpace(hash))
	fmt.Println("  Trx:", "0x6f34267ee323 (TODO)")
	fmt.Println(" Date:", "2018-08-15 12:29:34 UTC")

	fmt.Println()

	fmt.Println("Commit:", "Simon Tatham")
	fmt.Println("vChain:", "AK47")
	fmt.Println(" Trust:", "Strong verification (Level 3)")

	switch res.StatusCode {
	case 200:
		color.Set(color.FgHiWhite, color.BgCyan, color.Bold)
	case 404:
		color.Set(color.FgHiWhite, color.BgRed, color.Bold)
		defer os.Exit(1)
	}

	fmt.Print("Status: ", trust)
	color.Unset()

	fmt.Println()

	// how to use stderr and exit codes
	//fmt.Fprintln(os.Stderr, "err")
	//

}