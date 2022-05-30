package godicio

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "godicio",
	Short: "godicio is a CLI tool to search for words in dicio.com.br",
	Long: `godicio is a CLI tool to search for words in dicio.com.br.
	
	Example:
	
	godicio word
	`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
