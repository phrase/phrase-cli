package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Phrase Strings",
	Long:  `All software has versions. This is Phrase's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s\n", PHRASE_CLIENT_VERSION)
	},
}
