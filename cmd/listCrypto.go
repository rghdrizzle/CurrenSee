/*
Copyright © 2023 rghdrizzle 

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/rghdrizzle/CurrenSee/pkg"
)

// listCryptoCmd represents the listCrypto command
var listCryptoCmd = &cobra.Command{
	Use:   "crypto",
	Short: "View a comprehensive list of cryptocurrencies,",
	Long: `The listCrypto command allows you to retrieve a list of cryptocurrencies along with their short forms (symbols) and full names. This command is particularly useful for quickly referencing the common abbreviations used in the cryptocurrency domain.

	Usage:
	
	currensee crypto
	
	Example:
	
	$currensee crypto
	
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
	The provided example showcases a few common cryptocurrencies, but the listCrypto command will display a comprehensive list of various cryptocurrencies and their corresponding short forms and full names.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Crypto List")
		fmt.Println("-------------------")
		currensee.Getcrypto()
	},
}

func init() {

}
