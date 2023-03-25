package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	configFileFlag string

	rootCmd = &cobra.Command{
		Use:   "gitime",
		Short: "Sum up your /spent time on commits",
		Long:  `Gather information about /spent time from commit messages.`,
	}
)

// Execute the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	// Might use some config at some point for things like DaysInOneWeek
	//rootCmd.PersistentFlags().StringVar(&configFileFlag, "config", "", "config file (default is $HOME/.gitime.yaml)")

	// Snippets
	//rootCmd.PersistentFlags().StringP("author", "a", "YOUR NAME", "author name for copyright attribution")
	//rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "name of license for the project")
	//rootCmd.PersistentFlags().Bool("minutes", true, "show result in minutes")
	//viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	//viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
	//viper.SetDefault("author", "NAME HERE <EMAIL ADDRESS>")
	//viper.SetDefault("license", "apache")
}

func initConfig() {
	if configFileFlag != "" {
		viper.SetConfigFile(configFileFlag)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".gitime")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		//fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
