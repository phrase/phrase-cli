package cmd

import (
	commands "github.com/phrase/phrase-cli/cmd/internal"
	"github.com/spf13/cobra"
)

func init() {
	initInit()
}

func initInit() {
	var pullCmd = &cobra.Command{
		Use:   "init",
		Short: "Configure your Phrase client",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			cmdInit := commands.InitCommand{}
			err := cmdInit.Run()
			if err != nil {
				HandleError(err)
			}
		},
	}
	rootCmd.AddCommand(pullCmd)
}
