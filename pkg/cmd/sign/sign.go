/*
 * Copyright (c) 2018-2019 vChain, Inc. All Rights Reserved.
 * This software is released under GPL3.
 * The full license information can be found under:
 * https://www.gnu.org/licenses/gpl-3.0.en.html
 *
 */

package sign

import (
	"fmt"
	"strings"

	"github.com/vchain-us/vcn/pkg/extractor"

	"github.com/caarlos0/spin"
	"github.com/spf13/cobra"
	"github.com/vchain-us/vcn/internal/cli"
	"github.com/vchain-us/vcn/pkg/api"
	"github.com/vchain-us/vcn/pkg/meta"
	"github.com/vchain-us/vcn/pkg/store"
)

// NewCmdSign returns the cobra command for `vcn sign`
func NewCmdSign() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "sign",
		Aliases: []string{"s"},
		Short:   "Sign digital assets' hashes onto the blockchain",
		Long:    ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runSignWithState(cmd, args, meta.StatusTrusted)
		},

		Args: cobra.ExactArgs(1),
	}

	cmd.Flags().VarP(make(mapOpts), "attr", "a", "add user defined attributes (format: --attr key=value)")
	cmd.Flags().StringP("key", "k", "", "specify which user's key to use for signing, if not set the last available is used")
	cmd.Flags().BoolP("public", "p", false, "when signed as public, the asset name and the signer's identity will be visible to everyone")

	cmd.SetUsageTemplate(
		strings.Replace(cmd.UsageTemplate(), "{{.UseLine}}", "{{.UseLine}} ARG", 1),
	)

	return cmd
}

func runSignWithState(cmd *cobra.Command, args []string, state meta.Status) error {

	public, err := cmd.Flags().GetBool("public")
	if err != nil {
		return err
	}

	pubKey, err := cmd.Flags().GetString("key")
	if err != nil {
		return err
	}

	metadata := cmd.Flags().Lookup("attr").Value.(mapOpts).StringToInterface()

	cmd.SilenceUsage = true
	return sign(args[0], pubKey, state, meta.VisibilityForFlag(public), metadata)
}

func sign(arg string, pubKey string, state meta.Status, visibility meta.Visibility, metadata api.Metadata) error {

	if err := cli.AssertUserLogin(); err != nil {
		return err
	}
	u := api.NewUser(store.Config().CurrentContext)

	if err := cli.AssertUserKeystore(); err != nil {
		return err
	}

	// Extract artifact from arg
	a, err := extractor.Extract(arg)
	if err != nil {
		return err
	}

	// Copy user provided custom attributes
	a.Metadata.SetValues(metadata)

	if a.Size < 0 {
		return fmt.Errorf("invalid size")
	}

	if pubKey == "" {
		pubKey = u.DefaultKey()
	}

	fmt.Println("Signer:\t" + u.Email())
	fmt.Println("Key:\t" + pubKey)
	passphrase, err := cli.ProvidePassphrase()
	if err != nil {
		return err
	}

	s := spin.New("%s Signing asset...")
	s.Set(spin.Spin1)
	s.Start()

	// TODO: return and display: block #, trx #
	verification, err := u.Sign(*a, pubKey, passphrase, state, visibility)

	s.Stop()
	if err != nil {
		return err
	}

	fmt.Println()
	print(a, verification)
	return nil
}
