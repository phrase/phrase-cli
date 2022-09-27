package cmd

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
)

var (
	LAST_CHANGE           = "LIVE"
	REVISION              = "DEV"
	LIBRARY_REVISION      = "DEV"
	PHRASE_CLIENT_VERSION = "DEV"
)

func init() {
	initInfo()
}

func initInfo() {
	var infoCmd = &cobra.Command{
		Use:   "info",
		Short: "Info about CLI client for phrase",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Print(GetInfo())
		},
	}
	rootCmd.AddCommand(infoCmd)
}

func GetInfo() string {
	info := []string{
		fmt.Sprintf("Phrase Strings client version:            %s", PHRASE_CLIENT_VERSION),
		fmt.Sprintf("Phrase Strings client revision:           %s", REVISION),
		fmt.Sprintf("Phrase Strings library revision:          %s", LIBRARY_REVISION),
		fmt.Sprintf("Last change at:                   %s", LAST_CHANGE),
		fmt.Sprintf("Go version:                       %s", runtime.Version()),
	}
	return fmt.Sprintf("%s\n", strings.Join(info, "\n"))
}
