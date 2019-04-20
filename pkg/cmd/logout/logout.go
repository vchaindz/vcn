/*
 * Copyright (c) 2018-2019 vChain, Inc. All Rights Reserved.
 * This software is released under GPL3.
 * The full license information can be found under:
 * https://www.gnu.org/licenses/gpl-3.0.en.html
 *
 */

package logout

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/vchain-us/vcn/pkg/api"
)

// NewCmdLogout returns the cobra command for `vcn logout`
func NewCmdLogout() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "logout",
		Short: "Logout current user",
		Long:  ``,
		RunE:  runLogout,
		Args:  cobra.NoArgs,
	}

	return cmd
}

func runLogout(cmd *cobra.Command, args []string) error {

	if err := api.DeleteToken(); err != nil {
		cmd.SilenceUsage = true
		return err
	}

	fmt.Println("Logout successful.")
	return nil
}
