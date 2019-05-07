/*
 * Copyright (c) 2018-2019 vChain, Inc. All Rights Reserved.
 * This software is released under GPL3.
 * The full license information can be found under:
 * https://www.gnu.org/licenses/gpl-3.0.en.html
 *
 */

package verify

import (
	"encoding/json"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/ethereum/go-ethereum/common"
	"github.com/vchain-us/vcn/internal/cli"
	"github.com/vchain-us/vcn/pkg/api"
	"github.com/vchain-us/vcn/pkg/meta"
)

type result struct {
	Artifact     *api.ArtifactResponse       `json:"artifact"`
	Verification *api.BlockchainVerification `json:"verification"`
}

func print(output string, a *api.Artifact, artifact *api.ArtifactResponse, verification *api.BlockchainVerification) error {

	if output == "json" {

		r := result{
			Verification: verification,
		}
		if artifact != nil {
			r.Artifact = artifact
		} else if a != nil {
			r.Artifact = &api.ArtifactResponse{
				Name: a.Name,
				Kind: a.Kind,
				Hash: a.Hash,
				Size: a.Size,
			}
		}

		b, err := json.MarshalIndent(r, "", "  ")
		if err != nil {
			return err
		}
		fmt.Println(string(b))
		return nil
	}

	if output != "" {
		return fmt.Errorf("output format not supported: %s", output)
	}

	if artifact != nil {
		cli.PrintColumn("Name", artifact.Name, a.Name)
		cli.PrintColumn("Kind", artifact.Kind, "NA")
		cli.PrintColumn("Hash", a.Hash, "NA")
		cli.PrintColumn("Date", verification.Timestamp.String(), "NA")
		cli.PrintColumn("Signer", artifact.Publisher, "NA")
		cli.PrintColumn("Key", strings.ToLower(verification.Owner.Hex()), "NA")
		if artifact.Size > 0 {
			cli.PrintColumn("Size", humanize.Bytes(artifact.Size), "NA")
		} else {
			cli.PrintColumn("Size", "NA", "NA")
		}
		cli.PrintColumn("ContentType", artifact.ContentType, "NA")
		cli.PrintColumn("Url", artifact.Url, "NA")
		for k, v := range artifact.Metadata {
			if vv, err := json.Marshal(v); err == nil {
				cli.PrintColumn("Metadata", fmt.Sprintf("%s=%s\t", k, string(vv)), "")
			}
		}
		cli.PrintColumn("Company", artifact.PublisherCompany, "NA")
		cli.PrintColumn("Website", artifact.PublisherWebsiteUrl, "NA")
		cli.PrintColumn("Level", meta.LevelName(verification.Level), "NA")
	} else {
		cli.PrintColumn("Name", a.Name, "NA")
		cli.PrintColumn("Kind", a.Kind, "NA")
		cli.PrintColumn("Hash", a.Hash, "NA")
		if verification.Timestamp != time.Unix(0, 0) {
			cli.PrintColumn("Date", verification.Timestamp.String(), "NA")
		} else {
			cli.PrintColumn("Date", "NA", "NA")
		}
		cli.PrintColumn("Signer", "NA", "NA")
		if verification.Owner != common.BigToAddress(big.NewInt(0)) {
			cli.PrintColumn("Key", strings.ToLower(verification.Owner.Hex()), "NA")
		} else {
			cli.PrintColumn("Key", "NA", "NA")
		}
	}

	c, s := meta.StatusColor(verification.Status)
	cli.PrintColumn("Status", meta.StatusName(verification.Status), "NA", c, s)
	return nil
}
