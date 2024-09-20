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
				Cleanup:            params.GetBool("cleanup"),
				Branch:             params.GetString("branch"),
				UseLocalBranchName: params.GetBool("use-local-branch-name"),
				Tag:                params.GetString("tag"),
			}
			err := cmdPush.Run()
			if err != nil {
				HandleError(err)
			}
		},
	}
	rootCmd.AddCommand(pushCmd)

	AddFlag(pushCmd, "bool", "wait", "w", "Wait for files to be processed", false)
	AddFlag(pushCmd, "bool", "cleanup", "c", "Delete keys not mentioned in any of the uploads", false)
	AddFlag(pushCmd, "string", "branch", "b", "branch", false)
	AddFlag(pushCmd, "bool", "use-local-branch-name", "", "push from the branch with the name of your currently checked out branch (git or mercurial)", false)
	AddFlag(pushCmd, "string", "tag", "", "Tag uploaded keys", false)
	AddFlag(pushCmd, "string", "translation-key-prefix", "", "Prefix for the names of the translation keys", false)
	params.BindPFlags(pushCmd.Flags())
}
