package cmd

import (
	pull "github.com/phrase/phrase-cli/cmd/internal"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	initPull()
}

func initPull() {
	params := viper.New()
	var pullCmd = &cobra.Command{
		Use:   "pull",
		Short: "Pull translation changes",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			cmdPull := pull.PullCommand{
				Branch:             params.GetString("branch"),
				UseLocalBranchName: params.GetBool("use-local-branch-name"),
				Async:              params.GetBool("async"),
			}
			err := cmdPull.Run(Config)
			if err != nil {
				HandleError(err)
			}
		},
	}
	rootCmd.AddCommand(pullCmd)

	AddFlag(pullCmd, "string", "branch", "b", "branch", false)
	AddFlag(pullCmd, "bool", "use-local-branch-name", "", "use local branch name", false)
	AddFlag(pullCmd, "bool", "async", "a", "use asynchronous locale downloads (recommended for large number of keys)", false)
	params.BindPFlags(pullCmd.Flags())
}
