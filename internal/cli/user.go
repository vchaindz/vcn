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

	"github.com/vchain-us/vcn/pkg/api"
	"github.com/vchain-us/vcn/pkg/store"
)

const loginMsg = "Proceed by authenticating yourself using <vcn login>"

func AssertUserLogin() error {
	// check for token
	hasAuth, err := api.NewUser(store.Config().CurrentContext).IsAuthenticated()
	if err != nil {
		return err
	}

	if !hasAuth {
		return fmt.Errorf("You need to be logged in.\n%s", loginMsg)
	}

	return nil
}

func AssertUserKeystore() error {
	u := api.NewUser(store.Config().CurrentContext)
	if !u.HasKey() {
		return fmt.Errorf("You need a keystore to sign.\n%s", loginMsg)
	}
	return nil
}
