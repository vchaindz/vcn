/*
 * Copyright (c) 2018-2019 vChain, Inc. All Rights Reserved.
 * This software is released under GPL3.
 * The full license information can be found under:
 * https://www.gnu.org/licenses/gpl-3.0.en.html
 *
 */

package cli

import (
	"fmt"
	"os"

	"github.com/vchain-us/vcn/pkg/api"
)

func AssertUserLogin() {
	// check for token
	token, _ := api.LoadToken()
	checkOk, _ := api.CheckToken(token)
	if !checkOk {
		fmt.Println("You need to be logged in.")
		fmt.Println("Proceed by authenticating yourself using <vcn login>")
		// errors.PrintErrorURLCustom("token", 428)
		os.Exit(1)
	}
}

func AssertUserKeystore() {
	hasKeystore, _ := api.HasKeystore()
	if hasKeystore == false {
		fmt.Println("You need a keystore to sign.")
		fmt.Println("Proceed by authenticating yourself using <vcn auth>")
		// errors.PrintErrorURLCustom("keystore", 428)
		os.Exit(1)
	}
}
