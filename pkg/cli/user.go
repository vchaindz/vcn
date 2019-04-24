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
	"github.com/vchain-us/vcn/pkg/store"
)

func AssertUserLogin() {
	// check for token
	hasAuth, err := api.NewUser(store.Config().CurrentContext).IsAuthenticated()

	if err != nil {
		fmt.Println("Error: ", err)
	}

	if !hasAuth {
		fmt.Println("You need to be logged in.")
		fmt.Println("Proceed by authenticating yourself using <vcn login>")
		// errors.PrintErrorURLCustom("token", 428)
		os.Exit(1)
	}
}

func AssertUserKeystore() {
	u := api.NewUser(store.Config().CurrentContext)
	if !u.HasKey() {
		fmt.Println("You need a keystore to sign.")
		fmt.Println("Proceed by authenticating yourself using <vcn login>")
		// errors.PrintErrorURLCustom("keystore", 428)
		os.Exit(1)
	}
}
