package cmd

import (
	push "github.com/phrase/phrase-cli/cmd/internal"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	initPush()
}

func initPush() {
	params := viper.New()
	var pushCmd = &cobra.Command{
		Use:   "push",
		Short: "Push translation changes",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			cmdPush := push.PushCommand{
				Config:             *Config,
				Wait:               params.GetBool("wait"),
				Branch:             params.GetString("branch"),
				UseLocalBranchName: params.GetBool("use-local-branch-name"),
			}
			err := cmdPush.Run()
			if err != nil {
				HandleError(err)
			}
		},
	}
	rootCmd.AddCommand(pushCmd)

	AddFlag(pushCmd, "bool", "wait", "w", "Wait for files to be processed", false)
	AddFlag(pushCmd, "string", "branch", "b", "branch", false)
	AddFlag(pushCmd, "bool", "use-local-branch-name", "", "push from the branch with the name of your currently checked out branch (git or mercurial)", false)
	params.BindPFlags(pushCmd.Flags())
}
