/*
Copyright © 2023 rghdrizzle
*/
package cmd

import (
	"github.com/rghdrizzle/CurrenSee/pkg"
	"github.com/spf13/cobra"
)
var from string
var to string
// convertCmd represents the convert command
var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Easily convert from one currency to another",
	Long: `The convert command allows you to convert an amount from one currency to another. Specify the amount, the original currency, and the target currency to perform the conversion.
	Usage
	
	$currensee convert <amount> -f <from-currency> -t <to-currency>
	
	Arguments
	
		<amount>: The numeric value to convert.
		-f, --from <from-currency>: The code of the original currency.
		-t, --to <to-currency>: The code of the target currency.
	
	Example
	
	To convert 10 USD to INR, use the following command:
	
	$currensee convert 10 -f USD -t INR
	
	This will calculate and display the converted amount based on the latest exchange rates.
	
	Please note that the exchange rates used for conversion are subject to market fluctuations and may vary. The Currency CLI Tool retrieves the most up-to-date rates available at the time of execution.
	Feel free to explore the other commands available in the CurrenSee CLI Tool to list currencies and display exchange rates. `,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		from,_= cmd.Flags().GetString("from")
		to,_ = cmd.Flags().GetString("to")
		currensee.Convert(args[0],from,to)
	},
}

func init() {
	rootCmd.AddCommand(convertCmd)
	convertCmd.Flags().StringP("from","f","USD","The currency you want to convert from")
	convertCmd.Flags().StringP("to","t","INR","The currency you want to convert to")

}
