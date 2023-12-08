/*
Copyright © 2023 Alexandre Guerra <agmguerra@gmail.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var NumCards int

var rootCmd = &cobra.Command{
	Use:   "loterica",
	Short: "This CLI is responsible to generate lottery cards of Brazilian lottery games.",
	Long:  `This CLI is responsible to generate lottery cards of Brazilian lottery games.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().IntVarP(&NumCards, "numcards", "n", 1, "Number of cards generated")
	viper.BindPFlag("numcards", rootCmd.PersistentFlags().Lookup("numcards"))
}
