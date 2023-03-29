package cmd

import (
	"fmt"
	"github.com/goutte/git-spend/gitime"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	rootCmd = &cobra.Command{
		Use:   "git-spend",
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

	// Might use a config file as well at some point for things like DaysInOneWeek
	//rootCmd.PersistentFlags().StringVar(&configFileFlag, "config", "", "config file (default is $HOME/.git-spend.yaml)")

	// Snippets for viper config
	//rootCmd.PersistentFlags().StringP("author", "a", "YOUR NAME", "author name for copyright attribution")
	//viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	//viper.SetDefault("author", "NAME HERE <EMAIL ADDRESS>")
}

func initConfig() {
	// Later on we'll want users to be able to override the config file
	// But first we need to figure out how to generate a "template" for that config file.
	//if configFileFlag != "" {
	//	viper.SetConfigFile(configFileFlag)
	//} else {
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)
	viper.AddConfigPath(home)
	viper.SetConfigType("yaml")
	viper.SetConfigName(".git-spend")
	//}

	viper.SetEnvPrefix("git_spend")
	viper.AutomaticEnv()

	_ = viper.ReadInConfig()

	gitime.UpdateTimeModuloConfiguration()
}

func fail(anything interface{}) {
	fmt.Println(anything)
	os.Exit(1)
}
