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
	Short: "list to view all the currencies and crypto",
	Long: `The "list" command is a powerful feature of our CLI tool that displays a comprehensive list of all the current currencies along with their corresponding short forms.This command provides users with a quick reference to easily identify and work with various currencies within the tool.To utilize the "list" command, simply enter "list" in the command line interface.The tool will promptly retrieve and display the currency data, presenting a clean and organized list that includes the currency's short form and its full name.
Example Usage:

$ currensee list

Currencies:
- USD: United States Dollar
- EUR: Euro
- GBP: British Pound Sterling
- JPY: Japanese Yen
- ...

You can also view all crypto currency by doing the following:

$currensee list crypto

Cryptocurrency List:
	- BTC: Bitcoin
	- ETH: Ethereum
	- XRP: Ripple
	- LTC: Litecoin
	- BCH: Bitcoin Cash
	- ADA: Cardano
	- XLM: Stellar
	- DOT: Polkadot
	- UNI: Uniswap
	- DOGE: Dogecoin
	
	Note:
	The provided example showcases a few common cryptocurrencies, but the crypto command will display a comprehensive list of various
`,
	
	 Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Currency List")
		fmt.Println("-------------------")
		currensee.Getlist()
		
		
		
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.AddCommand(listCryptoCmd)

}
