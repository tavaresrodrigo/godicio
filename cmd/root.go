package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "godicio",
	Short: "Godicio is a CLI tool to search for words definitions in online dictionaries",
	Long: `Godicio helps CLI people to use a dictionary from the Terminal in different
	lanaguages as English, Spanish and Portuguese.
	
	Example:
	
	Searching for the Portuguese word "amigo".
	
	$godicio pt amigo

	Searching for the English word "friend".
	
	$gdicio en friend
	
	Searching for the Spanish word "amigo".
	$godicio es amigo

	`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}
	

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.godicio.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


