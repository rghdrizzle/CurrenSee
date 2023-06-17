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

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println("OH NO , there seems to be an error while running the cli :(",err)
		os.Exit(1)
	}
}

func init() {
	
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


