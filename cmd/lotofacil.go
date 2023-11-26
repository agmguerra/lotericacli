/*
Copyright © 2023 Alexandre Guerra <agmguerra@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// lotofacilCmd represents the lotofacil command
var lotofacilCmd = &cobra.Command{
	Use:   "lotofacil",
	Short: "Generate randomic numbers like card for Brazillian Lotofácil lottery game",
	Long: `Lotofacil is a Brazilian lottry game. Each card has from 15 to 20 numbers. Players can win 
a money prize with 11, 12, 13, 14, 15 matching numbers.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("lotofacil called")
	},
}

func init() {
	rootCmd.AddCommand(lotofacilCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// lotofacilCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// lotofacilCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
