package cmd

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Used for flags.
	cfgFile string

	Debug bool

	rootCmd = &cobra.Command{
		Use:   "phrase",
		Short: "Phrase is a translation management platform for software projects.",
		Long:  `You can collaborate on language file translation with your team or order translations through our platform. The API allows you to import locale files, download locale files, tag keys or interact in other ways with the localization data stored in Phrase for your account.`,
	}
)

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().BoolVarP(&Debug, "debug", "d", false, "show debug messages")
	viper.BindPFlag("debug", rootCmd.PersistentFlags().Lookup("debug"))
	viper.SetDefault("debug", false)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./.phrase.yaml fallback to $HOME/.phrase.yaml)")
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		HandleError(err)
	}
}

func AddFlag(cmd *cobra.Command, flagType string, name string, short string, description string) {
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

	cmd.MarkFlagRequired(name)
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
		viper.AutomaticEnv()

		if err := viper.ReadInConfig(); err == nil {
			if Debug {
				fmt.Println("Using config file:", viper.ConfigFileUsed())
			}
		}
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			HandleError(err)
		}

		viper.AddConfigPath(".")
		viper.AddConfigPath(home)
		viper.SetConfigName(".phrase")
		viper.AutomaticEnv()

		if err := viper.ReadInConfig(); err == nil {
			if Debug {
				fmt.Println("Using config file:", viper.ConfigFileUsed())
			}
		}

		viper.SetConfigName(".phraseapp")
		if err := viper.MergeInConfig(); err == nil {
			if Debug {
				fmt.Println("Using config file:", viper.ConfigFileUsed())
			}
		}
	}
}

func HandleError(msg interface{}) {
	fmt.Println("Error:", msg)
	os.Exit(1)
}
