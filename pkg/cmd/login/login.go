/*
 * Copyright (c) 2018-2019 vChain, Inc. All Rights Reserved.
 * This software is released under GPL3.
 * The full license information can be found under:
 * https://www.gnu.org/licenses/gpl-3.0.en.html
 *
 */

package login

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/vchain-us/vcn/pkg/api"
	"github.com/vchain-us/vcn/pkg/cli"
	"github.com/vchain-us/vcn/pkg/logs"
	"github.com/vchain-us/vcn/pkg/meta"
)

// NewCmdLogin returns the cobra command for `vcn login`
func NewCmdLogin() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "login",
		Short: "Sign-in to vChain.us",
		Long:  ``,
		RunE:  runLogin,
		Args:  cobra.NoArgs,
	}

	return cmd
}

func runLogin(cmd *cobra.Command, args []string) error {
	return login()
}

func login() error {

	// Check if a token already exists
	token, _ := api.LoadToken()
	tokenValid, err := api.CheckToken(token)
	if err != nil {
		return err
	}

	if !tokenValid {
		email, err := cli.ProvidePlatformUsername()
		if err != nil {
			return err
		}
		publisherExists, err := api.CheckPublisherExists(email)
		if err != nil {
			return err
		}
		if publisherExists {
			password, err := cli.ProvidePlatformPassword()
			if err != nil {
				return err
			}
			err = api.AuthenticateUser(email, password)
			if err != nil {
				return err
			}
		} else {
			return fmt.Errorf("no such user. please create an account at: %s", meta.DashboardURL())
		}
	}

	_ = api.TrackPublisher(meta.VcnLoginEvent)

	hasKeystore, err := api.HasKeystore()
	if err != nil {
		logs.LOG.WithFields(logrus.Fields{
			"error": err,
		}).Error("Could not access keystore")
		return fmt.Errorf("Could not access keystore")
	}
	if hasKeystore == false {

		fmt.Println("You have no keystore set up yet.")
		fmt.Println("<vcn> will now do this for you and upload the public key to the platform.")

		keystorePassphrase, err := cli.PromptKeystorePassphrase()
		if err != nil {
			return err
		}

		pubKey, wallet, err := api.CreateKeystore(keystorePassphrase)
		if err != nil {
			return err
		}

		fmt.Println("Keystore successfully created. We are updating your user profile.\n" +
			"You will be able to sign your first asset in one minute")
		fmt.Println("Public key:\t", pubKey)
		fmt.Println("Keystore:\t", wallet)

	}

	api.SyncKeys()

	fmt.Println("Login successful.")
	return nil
}
