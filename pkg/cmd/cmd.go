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

	"github.com/vchain-us/vcn/pkg/cmd/sign"

	"github.com/vchain-us/vcn/pkg/cmd/dashboard"
	"github.com/vchain-us/vcn/pkg/cmd/verify"

	"github.com/vchain-us/vcn/pkg/meta"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
		fmt.Println(err)
		defer os.Exit(1)
	}
	preExitHook(rootCmd)
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.vcn/config.yaml)")
	rootCmd.PersistentFlags().BoolP("quit", "q", true, "if false, ask for confirmation before quitting")

	//
	rootCmd.AddCommand(verify.NewCmdVerify())
	rootCmd.AddCommand(sign.NewCmdSign())
	rootCmd.AddCommand(dashboard.NewCmdDashboard())

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home + "/.vcn" directory with name "config.yaml".
		viper.AddConfigPath(home + "/.vcn")
		viper.SetConfigName("config.yaml")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func preExitHook(cmd *cobra.Command) {
	if quit, _ := cmd.PersistentFlags().GetBool("quit"); !quit {
		fmt.Println()
		fmt.Println("Press 'Enter' to continue")
		fmt.Scanln()
	}
}
