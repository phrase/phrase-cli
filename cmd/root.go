package cmd

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/bgentry/speakeasy"
	"github.com/phrase/phrase-cli/cmd/internal/updatechecker"
	"github.com/phrase/phrase-go/v2"
	api "github.com/phrase/phrase-go/v2"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Used for flags.
	cfgFile string
	Config  *phrase.Config

	rootCmd = &cobra.Command{
		Use:   "phrase",
		Short: "Phrase Strings is a translation management platform for software projects.",
		Long:  `You can collaborate on language file translation with your team or order translations through our platform. The API allows you to import locale files, download locale files, tag keys or interact in other ways with the localization data stored in Phrase Strings for your account.`,
	}
)

func init() {
	Config = &phrase.Config{
		Debug:     false,
		UserAgent: "Phrase Strings CLI (" + PHRASE_CLIENT_VERSION + ")",
	}

	cobra.OnInitialize(initConfig)
	cobra.OnInitialize(checkUpdate)

	rootCmd.PersistentFlags().BoolVarP(&Config.Debug, "verbose", "v", false, "show more messages")
	viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))
	viper.SetDefault("verbose", false)

	rootCmd.PersistentFlags().StringVarP(&Config.Credentials.Host, "host", "", "", "Host to send Request to")
	viper.BindPFlag("host", rootCmd.PersistentFlags().Lookup("host"))
	viper.SetDefault("host", false)

	rootCmd.PersistentFlags().StringVarP(&Config.Credentials.Token, "access_token", "t", "", "access token used for authentication")
	viper.BindPFlag("access_token", rootCmd.PersistentFlags().Lookup("access_token"))
	viper.SetDefault("access_token", false)

	rootCmd.PersistentFlags().StringVarP(&Config.Credentials.Username, "username", "u", "", "username used for authentication")
	viper.BindPFlag("username", rootCmd.PersistentFlags().Lookup("username"))
	viper.SetDefault("username", false)

	rootCmd.PersistentFlags().BoolVarP(&Config.Credentials.TFA, "tfa", "", false, "use Two-Factor Authentication")
	viper.BindPFlag("tfa", rootCmd.PersistentFlags().Lookup("tfa"))
	viper.SetDefault("tfa", false)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./.phrase.yml fallback to $HOME/.phrase.yml)")
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		HandleError(err)
	}
}

func AddFlag(cmd *cobra.Command, flagType string, name string, short string, description string, required ...bool) {
	switch flagType {
	case "bool":
		cmd.Flags().BoolP(name, short, false, description)
	case "int":
		cmd.Flags().Int64P(name, short, 0, description)
	case "int32":
		cmd.Flags().Int64P(name, short, 0, description)
	case "int64":
		cmd.Flags().Int64P(name, short, 0, description)
	case "float32":
		cmd.Flags().Float64P(name, short, 0, description)
	case "float64":
		cmd.Flags().Float64P(name, short, 0, description)
	default:
		cmd.Flags().StringP(name, short, "", description)
	}

	if len(required) == 1 {
		if required[0] {
			cmd.MarkFlagRequired(name)
		}
	} else {
		cmd.MarkFlagRequired(name)
	}

}

func initConfig() {
	config, err := phrase.ReadConfig(cfgFile)
	if err != nil {
		HandleError(err)
	}

	// flag overwrites debug option from file
	if Config.Debug {
		config.Debug = Config.Debug
	}
	config.UserAgent = Config.UserAgent

	if Config.Credentials.Host != "" {
		config.Credentials.Host = Config.Credentials.Host
	}

	if Config.Credentials.Token != "" {
		config.Credentials.Token = Config.Credentials.Token
	}

	if Config.Credentials.Username != "" {
		config.Credentials.Username = Config.Credentials.Username
	}

	phraseappAccessToken := os.Getenv("PHRASEAPP_ACCESS_TOKEN")
	if phraseappAccessToken != "" && config.Credentials.Token == "" && config.Credentials.Username == "" {
		config.Credentials.Token = phraseappAccessToken
	}

	phraseAccessToken := os.Getenv("PHRASE_ACCESS_TOKEN")
	if phraseAccessToken != "" && config.Credentials.Token == "" && config.Credentials.Username == "" {
		config.Credentials.Token = phraseAccessToken
	}

	if Config.Credentials.TFA {
		config.Credentials.TFA = Config.Credentials.TFA
	}

	if config.Debug {
		fmt.Printf("%+v\n", config)
	}

	Config = config
}

func Auth() context.Context {
	if Config.Credentials.Token != "" {
		return context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
			Key:    Config.Credentials.Token,
			Prefix: "token",
		})
	} else if Config.Credentials.Username != "" {
		pwd, err := speakeasy.Ask("Password: ")
		if err != nil {
			HandleError(err)
		}

		if Config.Credentials.TFA { // TFA only required for username+password based login.
			token, err := speakeasy.Ask("TFA-Token: ")
			if err != nil {
				HandleError(err)
			}
			Config.Credentials.TFAToken = token
		}
		return context.WithValue(context.Background(), api.ContextBasicAuth, api.BasicAuth{
			UserName: Config.Credentials.Username,
			Password: pwd,
		})
	} else {
		HandleError(fmt.Errorf("either username or token must be given"))
	}

	return nil
}

func checkUpdate() {
	var updateChecker = updatechecker.New(
		PHRASE_CLIENT_VERSION,
		filepath.Join(os.TempDir(), ".phrase.version"),
		"https://github.com/phrase/phrase-cli/releases/latest",
		os.Stderr,
	)

	updateChecker.Check()
}

func HandleError(msg interface{}) {
	fmt.Println("Error:", msg)
	os.Exit(1)
}
