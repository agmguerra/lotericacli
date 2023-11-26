/*
Copyright © 2023 Alexandre Guerra <agmguerra@gmail.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

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

}
