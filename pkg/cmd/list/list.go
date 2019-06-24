/*
 * Copyright (c) 2018-2019 vChain, Inc. All Rights Reserved.
 * This software is released under GPL3.
 * The full license information can be found under:
 * https://www.gnu.org/licenses/gpl-3.0.en.html
 *
 */

package list

import (
	"github.com/vchain-us/vcn/pkg/cmd/internal/cli"

	"github.com/spf13/cobra"
	"github.com/vchain-us/vcn/internal/assert"
	"github.com/vchain-us/vcn/pkg/api"
	"github.com/vchain-us/vcn/pkg/store"
)

// NewCmdList returns the cobra command for `vcn list`
func NewCmdList() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"l"},
		Short:   "List your signed assets",
		Long:    ``,
		RunE:    runList,
		Args:    cobra.NoArgs,
	}

	return cmd
}

func runList(cmd *cobra.Command, args []string) error {
	output, err := cmd.Flags().GetString("output")
	if err != nil {
		return err
	}
	cmd.SilenceUsage = true
	if err := assert.UserLogin(); err != nil {
		return err
	}
	u := api.NewUser(store.Config().CurrentContext)

	artifacts, err := u.LoadAllArtifacts()
	if err != nil {
		return err
	}
	return cli.PrintList(output, artifacts)
}
