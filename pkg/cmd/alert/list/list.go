/*
 * Copyright (c) 2018-2020 vChain, Inc. All Rights Reserved.
 * This software is released under GPL3.
 * The full license information can be found under:
 * https://www.gnu.org/licenses/gpl-3.0.en.html
 *
 */

package list

import (
	"fmt"

	"github.com/vchain-us/vcn/pkg/cmd/internal/cli"

	"github.com/vchain-us/vcn/internal/assert"
	"github.com/vchain-us/vcn/pkg/api"
	"github.com/vchain-us/vcn/pkg/store"

	"github.com/spf13/cobra"
)

// NewCommand returns the cobra command for `vcn alerts list`
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List configured alert",
		Long:  ``,
		Args:  cobra.NoArgs,
		RunE:  runList,
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
	if hasAuth, _ := u.IsAuthenticated(); !hasAuth {
		return fmt.Errorf("you need to be logged in, please use <vcn login>")
	}

	if output == "" {
		fmt.Printf("Listing locally configured alerts for %s...\n\n", u.Email())
	}
	alerts, err := store.ReadAlerts(u.Email())
	if err != nil {
		return err
	}

	if output == "" && len(alerts) == 0 {
		fmt.Printf("No results.\n\n")
		return nil
	}

	var list []api.AlertResponse

	for _, a := range alerts {
		aConfig := api.AlertConfig{}
		if err := a.ExportConfig(&aConfig); err != nil || aConfig.AlertUUID == "" {
			cli.PrintWarning(
				output,
				fmt.Sprintf(`invalid configuration for: %#v`, a),
			)
			continue
		}

		alert, err := u.GetAlert(aConfig.AlertUUID)
		if err != nil {
			cli.PrintWarning(
				output,
				fmt.Sprintf(`errored alert %s: %s`, aConfig.AlertUUID, err),
			)
			continue
		}
		list = append(list, *alert)
	}

	return cli.PrintObjects(output, list)
}
