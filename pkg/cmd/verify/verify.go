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
	"path/filepath"
	"strings"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/ethereum/go-ethereum/common"
	"github.com/fatih/color"
	"github.com/vchain-us/vcn/internal/docker"
	"github.com/vchain-us/vcn/internal/utils"
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
		Args:    cobra.MinimumNArgs(1),
	}

	cmd.SetUsageTemplate(
		strings.Replace(cmd.UsageTemplate(), "{{.UseLine}}", "{{.UseLine}} ...arg(s)", 1),
	)

	return cmd
}

func runVerify(cmd *cobra.Command, args []string) error {
	_ = api.TrackPublisher(meta.VcnVerifyEvent)
	for _, spec := range args {
		if ok, err := verify(spec); !ok {
			cmd.SilenceUsage = true
			cmd.SilenceErrors = true
			if err != nil {
				return err
			}
			return fmt.Errorf("%s is not verified", spec)
		}
	}

	return nil
}

func verify(filename string) (success bool, err error) {
	var artifactHash string
	// fixme(leogr): refactor to spec
	if strings.HasPrefix(filename, "docker:") {
		artifactHash, err = docker.GetHash(filename)
		if err != nil {
			return false, fmt.Errorf("failed to get hash for docker image: %s", err)
		}
	} else {
		hash, err := utils.HashFile(filename)
		if err != nil {
			return false, err
		}
		artifactHash = strings.TrimSpace(hash)
	}
	_ = api.TrackVerify(artifactHash, filepath.Base(filename))
	verification, err := api.BlockChainVerify(artifactHash)
	if err != nil {
		return false, fmt.Errorf("unable to verify hash: %s", err)
	}

	var artifact *api.ArtifactResponse
	if verification.Owner != common.BigToAddress(big.NewInt(0)) {
		artifact, _ = api.LoadArtifactForHash(artifactHash, verification.HashAsset())
	}
	if artifact != nil {
		printColumn("Asset", artifact.Filename, filepath.Base(filename))
		printColumn("Hash", artifactHash, "NA")
		printColumn("Date", verification.Timestamp.String(), "NA")
		printColumn("Signer", artifact.Publisher, verification.Owner.Hex())
		printColumn("Name", artifact.Name, "NA")
		if artifact.FileSize > 0 {
			printColumn("Size", humanize.Bytes(artifact.FileSize), "NA")
		} else {
			printColumn("Size", "NA", "NA")
		}
		printColumn("Company", artifact.PublisherCompany, "NA")
		printColumn("Website", artifact.PublisherWebsiteUrl, "NA")
		printColumn("Level", meta.LevelName(verification.Level), "NA")
	} else {
		printColumn("Asset", filepath.Base(filename), "NA")
		printColumn("Hash", artifactHash, "NA")
		if verification.Timestamp != time.Unix(0, 0) {
			printColumn("Date", verification.Timestamp.String(), "NA")
		} else {
			printColumn("Date", "NA", "NA")
		}
		if verification.Owner != common.BigToAddress(big.NewInt(0)) {
			printColumn("Signer", verification.Owner.Hex(), "NA")
		} else {
			printColumn("Signer", "NA", "NA")
		}
		printColumn("Name", "NA", "NA")
		printColumn("Company", "NA", "NA")
		printColumn("Website", "NA", "NA")
		printColumn("Size", "NA", "NA")
		printColumn("Level", "NA", "NA")
	}

	var c, s color.Attribute
	switch verification.Status {
	case meta.StatusTrusted:
		success = true
		c, s = meta.StyleSuccess()
	case meta.StatusUnknown:
		success = false
		c, s = meta.StyleWarning()
	default:
		success = false
		c, s = meta.StyleError()
	}
	printColumn("Status", meta.StatusName(verification.Status), "NA", c, s)
	return
}

func printColumn(field string, value string, fallback string, p ...color.Attribute) {
	var spaces string
	for i := len(field); i < 8; i++ {
		spaces += " "
	}
	fmt.Print(field + ":" + spaces)
	if p != nil {
		c := color.New(p...)
		c.Set()
	}
	if value != "" {
		fmt.Print(value)
	} else {
		fmt.Print(fallback)
	}
	if p != nil {
		color.Unset()
	}
	fmt.Println()
}
