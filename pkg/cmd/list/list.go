/*
 * Copyright (c) 2018-2019 vChain, Inc. All Rights Reserved.
 * This software is released under GPL3.
 * The full license information can be found under:
 * https://www.gnu.org/licenses/gpl-3.0.en.html
 *
 */

package list

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/vchain-us/vcn/pkg/api"
	"github.com/vchain-us/vcn/pkg/cli"
	"github.com/vchain-us/vcn/pkg/store"
)

// NewCmdList returns the cobra command for `vcn list`
func NewCmdList() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"l"},
		Short:   "List your signed artifacts",
		Long:    ``,
		RunE:    runList,
		Args:    cobra.NoArgs,
	}

	return cmd
}

func runList(cmd *cobra.Command, args []string) error {

	cli.AssertUserLogin()
	u := api.NewUser(store.Config().CurrentContext)

	artifacts, err := u.LoadAllArtifacts()
	if err != nil {
		cmd.SilenceUsage = true
		return err
	}

	for _, a := range artifacts {
		fmt.Print(a)
	}
	return nil
}
