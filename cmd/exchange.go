/*
Copyright © 2023 rghdrizzle
*/
package cmd

import (
	"fmt"

	"github.com/rghdrizzle/CurrenSee/pkg"
	"github.com/spf13/cobra"
)

// exchangeCmd represents the exchange command
var exchangeCmd = &cobra.Command{
	Use:   "exchange",
	Short: "A brief description of your command",
	Long: `The exchange command allows you to obtain the exchange rate between two currencies. It provides information such as the currency symbols, bid price, ask price, and mid price. This command is useful for quickly checking the current exchange rate when you need to convert one currency to another.

	Usage:
	
	exchange <base_currency><target_currency>
	
	Parameters:
	
		<base_currency>: The currency you want to convert from.
		<target_currency>: The currency you want to convert to.
	
	Example:
	
	$currensee exchange EURGBP
	
	requested time Fri, 16 Jun 2023 10:45:42 GMT timestamp 1686912342
	0
	symbol EUR GBP bid 0.85603 ask 0.85604 mid 0.85603
	
	Note:
	The example provided displays the exchange rate from Euro (EUR) to British Pound (GBP). The exchange command can be used with any valid currency symbols to retrieve the corresponding exchange rate information`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("exchange called")
		currensee.ExchangeRate(args[0])
	},
}

func init() {
	rootCmd.AddCommand(exchangeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// exchangeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// exchangeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
