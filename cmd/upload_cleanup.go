package cmd

import (
	"errors"

	uploadCleanup "github.com/phrase/phrase-cli/cmd/internal"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	initUpoadCleanup()
}

func initUpoadCleanup() {
	params := viper.New()
	var uploadCleanupCmd = &cobra.Command{
		Use:   "cleanup",
		Short: "Delete unmentioned keys for given upload",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			projectId := params.GetString("project_id")
			if projectId == "" {
				projectId = Config.DefaultProjectID
			}
			if projectId == "" {
				HandleError(errors.New("required flag \"project_id\" not set"))
				return
			}

			cmduploadCleanup := uploadCleanup.UploadCleanupCommand{
				Config:    *Config,
				ID:        params.GetString("id"),
				ProjectID: projectId,
				Confirm:   params.GetBool("confirm"),
				Branch:    params.GetString("branch"),
			}
			err := cmduploadCleanup.Run()
			if err != nil {
				HandleError(err)
			}
		},
	}
	UploadsApiCmd.AddCommand(uploadCleanupCmd)
	AddFlag(uploadCleanupCmd, "bool", "confirm", "y", "Don’t ask for confirmation", false)
	AddFlag(uploadCleanupCmd, "string", "id", "", "Upload ID", true)
	AddFlag(uploadCleanupCmd, "string", "project_id", "", "Project ID - required if the current directory does not contain a config with a project_id", false)
	AddFlag(uploadCleanupCmd, "string", "branch", "", "Branch", false)
	params.BindPFlags(uploadCleanupCmd.Flags())
}
