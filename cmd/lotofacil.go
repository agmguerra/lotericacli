/*
Copyright © 2023 Alexandre Guerra <agmguerra@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/agmguerra/lotericacli/lotofacil"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// lotofacilCmd represents the lotofacil command
var lotofacilCmd = &cobra.Command{
	Use:   "lotofacil",
	Short: "Generate randomic numbers like card for Brazillian Lotofácil lottery game",
	Long: `Lotofacil is a Brazilian lottry game. Each card has from 15 to 20 numbers. Players can win 
a money prize with 11, 12, 13, 14, 15 matching numbers.`,
	Run: func(cmd *cobra.Command, args []string) {
		cards := lotofacil.GetLotofacilCards(viper.GetInt("numcards"))
		fmt.Println(cards)
	},
}

func init() {
	rootCmd.AddCommand(lotofacilCmd)
}
