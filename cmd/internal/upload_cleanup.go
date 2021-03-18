package internal

import (
	"fmt"
	"sort"
	"strings"

	"github.com/antihax/optional"
	prompt "github.com/phrase/phrase-cli/cmd/internal/prompt"
	"github.com/phrase/phrase-go/v2"
)

type UploadCleanupCommand struct {
	phrase.Config
	ID      string
	Confirm bool
	Branch  string
}

func (cmd *UploadCleanupCommand) Run() error {
	Config = &cmd.Config
	client := newClient()

	return UploadCleanup(client, cmd)
}

func UploadCleanup(client *phrase.APIClient, cmd *UploadCleanupCommand) error {
	q := "unmentioned_in_upload:" + cmd.ID
	keysListLocalVarOptionals := phrase.KeysListOpts{
		Page:    optional.NewInt32(1),
		PerPage: optional.NewInt32(100),
		Q:       optional.NewString(q),
		Branch:  optional.NewString(cmd.Branch),
	}
	keys, _, err := client.KeysApi.KeysList(Auth, cmd.Config.DefaultProjectID, &keysListLocalVarOptionals)
	if err != nil {
		return err
	}

	if len(keys) == 0 {
		fmt.Println("There were no keys unmentioned in that upload.")
		return nil
	}

	for len(keys) != 0 {
		ids := make([]string, len(keys))
		names := make([]string, len(keys))
		for i, key := range keys {
			ids[i] = key.Id
			names[i] = key.Name
		}

		if !cmd.Confirm {
			fmt.Println("You are about to delete the following key(s) from your project:")
			sort.Strings(names)
			fmt.Println(strings.Join(names, "\n"))

			confirmation := ""
			err := prompt.WithDefault("Are you sure you want to continue? (y/n)", &confirmation, "n")
			if err != nil {
				return err
			}

			if strings.ToLower(confirmation) != "y" {
				fmt.Println("Clean up aborted")
				return nil
			}
		}

		q := "ids:" + strings.Join(ids, ",")
		keysDeletelocalVarOptionals := phrase.KeysDeleteCollectionOpts{
			Q:      optional.NewString(q),
			Branch: optional.NewString(cmd.Branch),
		}
		affected, _, err := client.KeysApi.KeysDeleteCollection(Auth, cmd.Config.DefaultProjectID, &keysDeletelocalVarOptionals)

		if err != nil {
			return err
		}

		fmt.Printf("%d key(s) successfully deleted.\n", affected.RecordsAffected)

		keysListLocalVarOptionals.Page = optional.NewInt32(keysListLocalVarOptionals.Page.Value() + 1)

		keys, _, err = client.KeysApi.KeysList(Auth, cmd.Config.DefaultProjectID, &keysListLocalVarOptionals)
		if err != nil {
			return err
		}
	}

	return nil
}
