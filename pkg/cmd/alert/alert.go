/*
 * Copyright (c) 2018-2020 vChain, Inc. All Rights Reserved.
 * This software is released under GPL3.
 * The full license information can be found under:
 * https://www.gnu.org/licenses/gpl-3.0.en.html
 *
 */

package alert

import (
	"github.com/vchain-us/vcn/internal/assert"
	"github.com/vchain-us/vcn/pkg/cmd/alert/list"

	"github.com/spf13/cobra"
)

// NewCommand returns the cobra command for `vcn alerts`
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "alerts",
		Aliases: []string{"alert"},
		Short:   "Manage alerts",
		Long:    ``,
		Args:    cobra.NoArgs,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			if err := assert.UserLogin(); err != nil {
				return err
			}
			return nil
		},
	}

	cmd.AddCommand(list.NewCommand())

	return cmd
}
