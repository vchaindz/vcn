/*
 * Copyright (c) 2018-2019 vChain, Inc. All Rights Reserved.
 * This software is released under GPL3.
 * The full license information can be found under:
 * https://www.gnu.org/licenses/gpl-3.0.en.html
 *
 */

package cmd

import (
	"fmt"
	"os"

	"github.com/inconshreveable/mousetrap"
	"github.com/vchain-us/vcn/internal/migrate"
	"github.com/vchain-us/vcn/pkg/cmd/list"
	"github.com/vchain-us/vcn/pkg/cmd/login"
	"github.com/vchain-us/vcn/pkg/cmd/logout"
	"github.com/vchain-us/vcn/pkg/cmd/sign"

	"github.com/vchain-us/vcn/pkg/cmd/dashboard"
	"github.com/vchain-us/vcn/pkg/cmd/verify"

	"github.com/vchain-us/vcn/pkg/meta"

	"github.com/spf13/cobra"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "vcn",
	Version: meta.Version(),
	Short:   "vChain CodeNotary - code signing in 1 simple step",
	Long:    ``,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		defer os.Exit(1)
	}
	preExitHook(rootCmd)
}

func init() {

	// Migrate old profile dirs, if any
	migrate.From03x()

	// Disable default behavior when started through explorer.exe
	cobra.MousetrapHelpText = ""

	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.vcn/config.json)")
	rootCmd.PersistentFlags().BoolP("quit", "q", true, "if false, ask for confirmation before quitting")
	rootCmd.PersistentFlags().MarkHidden("quit")

	// Verification group
	rootCmd.AddCommand(verify.NewCmdVerify())
	rootCmd.AddCommand(list.NewCmdList())

	// Signing group
	rootCmd.AddCommand(sign.NewCmdSign())
	rootCmd.AddCommand(sign.NewCmdUntrust())
	rootCmd.AddCommand(sign.NewCmdUnsupport())

	// User group
	rootCmd.AddCommand(login.NewCmdLogin())
	rootCmd.AddCommand(logout.NewCmdLogout())
	rootCmd.AddCommand(dashboard.NewCmdDashboard())

}

func preExitHook(cmd *cobra.Command) {
	if quit, _ := cmd.PersistentFlags().GetBool("quit"); !quit || mousetrap.StartedByExplorer() {
		fmt.Println()
		fmt.Println("Press 'Enter' to continue...")
		fmt.Scanln()
	}
}
