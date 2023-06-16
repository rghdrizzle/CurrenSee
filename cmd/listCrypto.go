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
	Use:   "listCrypto",
	Short: "View a comprehensive list of cryptocurrencies,",
	Long: `The listCrypto command allows you to retrieve a list of cryptocurrencies along with their short forms (symbols) and full names. This command is particularly useful for quickly referencing the common abbreviations used in the cryptocurrency domain.

	Usage:
	
	currensee listCrypto
	
	Example:
	
	$currensee listCrypto
	
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
		fmt.Println("listCrypto called")
		fmt.Println("Crypto List")
		fmt.Println("-------------------")
		currensee.Getcrypto()
	},
}

func init() {
	rootCmd.AddCommand(listCryptoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCryptoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCryptoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
