/*
Copyright © 2023 rghdrizzle 

*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/rghdrizzle/CurrenSee/pkg"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list to view all the currencies",
	Long: `The "list" command is a powerful feature of our CLI tool that displays a comprehensive list of all the current currencies along with their corresponding short forms.This command provides users with a quick reference to easily identify and work with various currencies within the tool.To utilize the "list" command, simply enter "list" in the command line interface.The tool will promptly retrieve and display the currency data, presenting a clean and organized list that includes the currency's short form and its full name.
Example Usage:

$ currensee list

Currencies:
- USD: United States Dollar
- EUR: Euro
- GBP: British Pound Sterling
- JPY: Japanese Yen
- ...
`,
	
	 Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Currency List")
		fmt.Println("-------------------")
		currensee.Getlist()
		
		
		
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
