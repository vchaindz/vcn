/*
 * Copyright (c) 2018-2019 vChain, Inc. All Rights Reserved.
 * This software is released under GPL3.
 * The full license information can be found under:
 * https://www.gnu.org/licenses/gpl-3.0.en.html
 *
 * Built on top of CLI (MIT license)
 * https://github.com/urfave/cli#overview
 */

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"github.com/vchain-us/vcn/pkg/api"
	vcn "github.com/vchain-us/vcn/pkg/cli"
	"github.com/vchain-us/vcn/pkg/logs"
	"github.com/vchain-us/vcn/pkg/meta"
)

func main() {
	var publicSigning bool
	var quit bool
	var acknowledge bool
	vcn.CreateVcnDirectories()
	app := cli.NewApp()
	app.Name = "CodeNotary vcn"
	app.Usage = "code signing in 1 simple step"
	app.Version = meta.VcnVersion
	app.Commands = []cli.Command{
		{
			Category: "Artifact actions",
			Name:     "verify",
			Aliases:  []string{"v"},
			Usage:    "Verify digital artifact against blockchain",
			Action: func(c *cli.Context) error {
				if c.NArg() == 0 {
					return fmt.Errorf("assets required")
				}
				vcn.VerifyAll(c.Args(), quit)
				return nil
			},
			Flags: []cli.Flag{
				cli.BoolFlag{Name: "hash"},
				cli.BoolTFlag{Name: "quit, q", Destination: &quit},
			},
		},
		{
			Category: "Artifact actions",
			Name:     "sign",
			Aliases:  []string{"s"},
			Usage:    "Sign digital assets' hashes onto the blockchain",
			Action: func(c *cli.Context) error {
				if c.NArg() == 0 {
					return fmt.Errorf("filename or type:reference required")
				}
				vcn.Sign(c.Args().First(), meta.StatusTrusted, meta.VisibilityForFlag(publicSigning), quit, acknowledge)
				return nil
			},
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:        "public, p",
					Usage:       "when signed as public, the asset name and the signer's identity will be visible to everyone",
					Destination: &publicSigning},
				cli.BoolTFlag{
					Name:        "quit, q",
					Destination: &quit},
				cli.BoolFlag{
					Name:        "yes, y",
					Usage:       "when used, you automatically confirm the ownership of this asset",
					Destination: &acknowledge},
			},
		},
		{
			Category: "Artifact actions",
			Name:     "untrust",
			Aliases:  []string{"ut"},
			Usage:    "Untrust a digital asset.",
			Action: func(c *cli.Context) error {
				if c.NArg() == 0 {
					return fmt.Errorf("filename or type:reference required")
				}
				vcn.Sign(c.Args().First(), meta.StatusUntrusted, meta.VisibilityForFlag(publicSigning), quit, acknowledge)
				return nil
			},
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:        "public, p",
					Usage:       "when signed as public, the asset name and the signer's identity will be visible to everyone",
					Destination: &publicSigning},
				cli.BoolTFlag{Name: "quit, q", Destination: &quit},
				cli.BoolFlag{
					Name:        "yes, y",
					Usage:       "when used, you automatically confirm the ownership of this asset",
					Destination: &acknowledge},
			},
		},
		{
			Category: "Artifact actions",
			Name:     "unsupport",
			Aliases:  []string{"us"},
			Usage:    "Unsupport a digital asset.",
			Action: func(c *cli.Context) error {
				if c.NArg() == 0 {
					return fmt.Errorf("filename or type:reference required")
				}
				vcn.Sign(c.Args().First(), meta.StatusUnsupported, meta.VisibilityForFlag(publicSigning), quit, acknowledge)
				return nil
			},
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:        "public, p",
					Usage:       "when signed as public, the asset name and the signer's identity will be visible to everyone",
					Destination: &publicSigning},
				cli.BoolTFlag{Name: "quit, q", Destination: &quit},
				cli.BoolFlag{
					Name:        "yes, y",
					Usage:       "when used, you automatically confirm the ownership of this asset",
					Destination: &acknowledge},
			},
		},
		{
			Category: "Artifact actions",
			Name:     "list",
			Aliases:  []string{"l"},
			Usage:    "List your signed artifacts",
			Action: func(c *cli.Context) error {
				artifacts, err := api.LoadArtifactsForCurrentWallet()
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println("Artifacts:\n", artifacts)
				return nil
			},
		},
		{
			Category: "User actions",
			Name:     "login",
			Usage:    "Sign-in to vChain.us",
			Action: func(c *cli.Context) error {

				vcn.Login()
				return nil
			},
		},
		{
			Category: "User actions",
			Name:     "dashboard",
			Aliases:  []string{"d"},
			Usage:    "Open dashboard at vChain.us in browser",
			Action: func(c *cli.Context) error {

				vcn.Dashboard()
				return nil
			},
		},
	}
	app.CommandNotFound = func(c *cli.Context, command string) {
		fmt.Println("No such command:", command)
		_ = cli.ShowAppHelp(c)
	}
	logs.LOG.WithFields(logrus.Fields{
		"version": meta.VcnVersion,
		"stage":   meta.StageName(meta.StageEnvironment()),
	}).Trace("VCN")
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
