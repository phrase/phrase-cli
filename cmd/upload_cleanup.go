package cmd

import (
	uploadCleanup "github.com/phrase/phrase-cli/cmd/internal"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	initUpoadCleanup()
}

func initUpoadCleanup() {
	params := viper.New()
	var upoadCleanupCmd = &cobra.Command{
		Use:   "cleanup",
		Short: "Delete unmentioned keys for given upload",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			cmduploadCleanup := uploadCleanup.UploadCleanupCommand{
				Config:  *Config,
				ID:      params.GetString("id"),
				Confirm: params.GetBool("confirm"),
			}
			err := cmduploadCleanup.Run()
			if err != nil {
				HandleError(err)
			}
		},
	}
	UploadsApiCmd.AddCommand(upoadCleanupCmd)
	AddFlag(upoadCleanupCmd, "bool", "confirm", "y", "Don’t ask for confirmation", false)
	AddFlag(upoadCleanupCmd, "string", "id", "", "Upload id", true)
	params.BindPFlags(upoadCleanupCmd.Flags())
}
