package cmd

import (
	"github.com/phrase/phrase-cli/cmd/pull"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(pullCmd)
}

var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Pull transaltion chnages",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		cmdPull := pull.PullCommand{}
		err := cmdPull.Run(Config)
		if err != nil {
			HandleError(err)
		}
	},
}
