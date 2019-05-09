/*
 * Copyright (c) 2018-2019 vChain, Inc. All Rights Reserved.
 * This software is released under GPL3.
 * The full license information can be found under:
 * https://www.gnu.org/licenses/gpl-3.0.en.html
 *
 */

package verify

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/vchain-us/vcn/pkg/extractor"

	"github.com/vchain-us/vcn/pkg/store"

	"github.com/ethereum/go-ethereum/common"
	"github.com/vchain-us/vcn/pkg/api"
	"github.com/vchain-us/vcn/pkg/meta"

	"github.com/spf13/cobra"
)

// NewCmdVerify returns the cobra command for `vcn verify`
func NewCmdVerify() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "verify",
		Example: "  vcn verify /bin/vcn",
		Aliases: []string{"v"},
		Short:   "Verify digital artifact against blockchain",
		Long:    ``,
		RunE:    runVerify,
		Args: func(cmd *cobra.Command, args []string) error {
			if hash, _ := cmd.Flags().GetString("hash"); hash != "" {
				if len(args) > 0 {
					return fmt.Errorf("cannot use arg(s) with --hash")
				}
				return nil
			}
			return cobra.MinimumNArgs(1)(cmd, args)
		},
	}

	cmd.SetUsageTemplate(
		strings.Replace(cmd.UsageTemplate(), "{{.UseLine}}", "{{.UseLine}} ...ARG(s)", 1),
	)

	cmd.Flags().StringP("key", "k", "", "specify the public key <vcn> should use, if not set the last available is used")
	cmd.Flags().String("hash", "", "specify a hash to verify, if set no arg(s) can be used")
	cmd.Flags().StringP("output", "o", "", "output format, one of: --output=json|--output=yaml|--output=''")

	return cmd
}

func runVerify(cmd *cobra.Command, args []string) error {
	hash, err := cmd.Flags().GetString("hash")
	if err != nil {
		return err
	}
	pubKey, err := cmd.Flags().GetString("key")
	if err != nil {
		return err
	}
	output, err := cmd.Flags().GetString("output")
	if err != nil {
		return err
	}
	cmd.SilenceUsage = true

	user := api.NewUser(store.Config().CurrentContext)

	// by hash
	if hash != "" {
		a := &api.Artifact{
			Hash: hash,
		}
		if err := verify(cmd, a, pubKey, user, output); err != nil {
			return err
		}
		return nil
	}

	// else by args
	for _, arg := range args {
		a, err := extractor.Extract(arg)
		if err != nil {
			return err
		}
		if err := verify(cmd, a, pubKey, user, output); err != nil {
			return err
		}
	}

	return nil
}

func verify(cmd *cobra.Command, a *api.Artifact, pubKey string, user *api.User, output string) (err error) {
	var verification *api.BlockchainVerification
	if pubKey != "" {
		// if a key has been passed, check for a verification matching that key
		verification, err = api.BlockChainVerifyMatchingPublicKey(a.Hash, pubKey)
	} else {
		if pubKeys := user.Keys(); len(pubKeys) > 0 {
			// if we have an user, check for verification matching user's keys first
			verification, err = api.BlockChainVerifyMatchingPublicKeys(a.Hash, pubKeys)
		}
		// if no user nor verification matching the user has found,
		// fallback to the last with highest level avaiable verification
		if verification == nil {
			verification, err = api.BlockChainVerify(a.Hash)
		}
	}
	if err != nil {
		return fmt.Errorf("unable to verify hash: %s", err)
	}

	var artifact *api.ArtifactResponse
	if verification.Owner != common.BigToAddress(big.NewInt(0)) {
		artifact, _ = api.LoadArtifactForHash(user, a.Hash, verification.MetaHash())
	}

	if err = print(output, a, artifact, verification); err != nil {
		return err
	}

	// todo(ameingast): redundant tracking events?
	_ = api.TrackPublisher(user, meta.VcnVerifyEvent)
	_ = api.TrackVerify(user, a.Hash, a.Name)

	if verification.Status != meta.StatusTrusted {
		if pubKey != "" {
			err = fmt.Errorf("%s is not verified by %s", a.Hash, pubKey)
		} else if email := user.Email(); email != "" {
			err = fmt.Errorf("%s is not verified by %s", a.Hash, email)
		} else {
			err = fmt.Errorf("%s is not verified", a.Hash)
		}
	}

	if output != "" {
		cmd.SilenceErrors = true
	}

	return
}
