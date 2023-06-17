/*
Copyright © 2023 rghdrizzle
*/
package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
)
var version = "0.0.1"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "currensee",
	Version: version,
	Short: "A cli tool to view Currency exchange",
	Long: `CuurenSee - A beautiful cli crafted by rghdrizzle which can be used for listing the currency of every type of currency and check their histories.
	Dw it gets updated every 2min`,
	Run: func(cmd *cobra.Command, args []string) {
		banner()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println("OH NO , there seems to be an error while running the cli :(",err)
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.CurrenSee.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


