package internal

import (
	"github.com/phrase/phrase-go/v4"
)

type UploadCleanupCommand struct {
	phrase.Config
	IDs       []string
	ProjectID string
	Confirm   bool
	Branch    string
}

func (cmd *UploadCleanupCommand) Run() error {
	Config = &cmd.Config
	client := newClient()

	return UploadCleanup(client, cmd.Confirm, cmd.IDs, cmd.Branch, cmd.ProjectID)
}
