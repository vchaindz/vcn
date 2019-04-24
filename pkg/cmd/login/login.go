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

	"github.com/vchain-us/vcn/pkg/store"

	"github.com/spf13/cobra"
	"github.com/vchain-us/vcn/pkg/api"
	"github.com/vchain-us/vcn/pkg/cli"
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
	cmd.SilenceUsage = true
	return login()
}

func login() error {

	email, err := cli.ProvidePlatformUsername()
	if err != nil {
		return err
	}

	user := api.NewUser(email)
	isExist, err := user.IsExist()
	if err != nil {
		return err
	}
	if !isExist {
		return fmt.Errorf("no such user, please create an account at: %s", meta.DashboardURL())
	}

	password, err := cli.ProvidePlatformPassword()
	if err != nil {
		return err
	}

	if err := user.Authenticate(password); err != nil {
		return err
	}

	_ = api.TrackPublisher(user, meta.VcnLoginEvent)

	store.Config().CurrentContext = user.Email()
	if err := store.SaveConfig(); err != nil {
		return err
	}

	if !user.HasKey() {

		fmt.Println("You have no keys set up yet.")
		fmt.Println("<vcn> will now do this for you and upload the public key to the platform.")

		keyPassphrase, err := cli.PromptPassphrase()
		if err != nil {
			return err
		}

		keystore, err := user.DefaultKeystore()
		if err != nil {
			return err
		}
		if err := store.SaveConfig(); err != nil {
			return err
		}

		pubKey, err := keystore.CreateKey(keyPassphrase)
		if err != nil {
			return err
		}

		fmt.Println("Keystore successfully created. We are updating your user profile.\n" +
			"You will be able to sign your first asset in one minute")
		fmt.Println("Keystore:\t", keystore.Path)
		fmt.Println("Public key:\t", pubKey)

		_ = api.TrackPublisher(user, meta.KeyCreatedEvent)
	}

	err = user.SyncKeys()
	if err != nil {
		return err
	}

	fmt.Println("Login successful.")
	return nil
}
