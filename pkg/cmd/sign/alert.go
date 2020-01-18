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
	"os"
	"path/filepath"
	"strings"

	"github.com/vchain-us/vcn/pkg/store"

	"github.com/vchain-us/vcn/pkg/extractor/dir"
	"github.com/vchain-us/vcn/pkg/extractor/file"
	"github.com/vchain-us/vcn/pkg/extractor/git"
	"github.com/vchain-us/vcn/pkg/uri"

	"github.com/vchain-us/vcn/pkg/api"
)

func handleAlert(arg string, u api.User, name string, a api.Artifact, v api.BlockchainVerification, output string) error {

	// make path absolute
	aURI, err := uri.Parse(arg)
	if err != nil {
		return fmt.Errorf("invalid argument for alert: %s", arg)
	}

	switch a.Kind {
	case file.Scheme:
		fallthrough
	case dir.Scheme:
		fallthrough
	case git.Scheme:
		absPath, err := filepath.Abs(strings.TrimPrefix(aURI.Opaque, "//"))
		if err != nil {
			return err
		}
		aURI.Opaque = "//" + absPath
		arg = aURI.String()
	}

	m := api.Metadata{}

	hostname, _ := os.Hostname()
	if hostname != "" {
		m["hostname"] = hostname
	}

	if name == "" {
		name = hostname
	}

	alertConfig, err := u.CreateAlert(name, a, v, m)
	if err != nil {
		return fmt.Errorf("cannot create alert: %s", err)
	}

	if output == "" {
		fmt.Printf("\nAlert %s has been created.\n", alertConfig.AlertUUID)
	}

	if err := store.SaveAlert(u.Email(), alertConfig.AlertUUID, store.Alert{
		Name:   name,
		Arg:    arg,
		Config: alertConfig,
	}); err != nil {
		return fmt.Errorf("cannot save alert: %s", err)
	}

	if output == "" {
		configPath, _ := store.AlertFilepath(u.Email())
		fmt.Printf("\nThe alert configuration has been added to %s.\n", configPath)
	}

	return nil
}
