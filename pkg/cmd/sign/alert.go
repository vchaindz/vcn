/*
 * Copyright (c) 2018-2020 vChain, Inc. All Rights Reserved.
 * This software is released under GPL3.
 * The full license information can be found under:
 * https://www.gnu.org/licenses/gpl-3.0.en.html
 *
 */

package sign

import (
	"fmt"

	"github.com/vchain-us/vcn/pkg/api"
	"github.com/vchain-us/vcn/pkg/cmd/internal/cli"
)

func handleAlert(alertConfigFile string, u api.User, a api.Artifact, v api.BlockchainVerification, output string) error {
	if alertConfigFile == "" {
		return nil
	}

	alertConfig, err := u.CreateAlert(a, v, api.Metadata{})
	if err != nil {
		return err
	}

	if output == "" {
		fmt.Printf("\nAlert %s has been created.\n", alertConfig.AlertUUID)
	}

	if err := cli.WriteYAML(alertConfig, alertConfigFile); err != nil {
		return err
	}
	if output == "" {
		fmt.Printf("Alert configuration saved to %s.\n", alertConfigFile)
	}
	return nil
}
