/*
Copyright © 2022 Rodrigo Tavares <rodrigo.actavares@gmail.com>

*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tavaresrodrigo/godicio/ledictionnaire"
)

// ptCmd represents the pt command
var frCmd = &cobra.Command{
	Use:   "fr",
	Short: "Get the word definition in English",
	Long: `This command will get the word definition in English from oxfordlearnersdictionaries.com/.`,
	Run: func(cmd *cobra.Command, args []string) {
		ledictionnaire.GetFRWord(args)
	},
}

func init() {
	rootCmd.AddCommand(frCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ptCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ptCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}



