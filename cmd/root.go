package cmd

import (
	"fmt"
	"github.com/goutte/git-spend/gitime"
	"github.com/goutte/git-spend/locale"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	rootCmd = &cobra.Command{
		Use:               "git-spend",
		Short:             locale.T("CommandRootSummary"),
		Long:              locale.T("CommandRootDescription"),
		DisableAutoGenTag: true,
	}
)

// Execute the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// If we want the generated help to show correct defaults, we need this BEFORE cobra inits
	initConfig()

	// We will need to use a listener like so when we'll use a flag for the config file
	//cobra.OnInitialize(initCobra)

	// Might use a config file as well at some point for things like DaysInOneWeek
	//rootCmd.PersistentFlags().StringVar(&configFileFlag, "config", "", "config file (default is $HOME/.git-spend.yaml)")

	// Snippets for viper config
	//rootCmd.PersistentFlags().StringP("author", "a", "YOUR NAME", "author name for copyright attribution")
	//viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
}

func initConfig() {
	// Later on we'll want users to be able to override the config file
	// But first we need to figure out how to generate a "template" for that config file.
	//if configFile != "" {
	//	viper.SetConfigFile(configFile)
	//} else {
	home, err := os.UserHomeDir()
	if err != nil {
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".git-spend")
	}
	//}

	viper.SetEnvPrefix("git_spend")
	viper.AutomaticEnv()

	_ = viper.ReadInConfig()

	gitime.UpdateTimeModuloConfiguration()
}

func fail(anything interface{}, command *cobra.Command) {
	// Questions:
	// - stderr ?
	// - CLI colors ?
	fmt.Println(anything)

	_ = command.Help()
	os.Exit(1)
}
