package cmd

import (
	"fmt"
	"github.com/goutte/git-spend/locale"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
	"log"
	"os"
)

// manSection is to be determined/discussed
const manSection = 8

// manPath is injected with manSection, and holds where the man pages are installed
const manPath = "/usr/local/share/man/man%d/"

var (
	FlagOutput  string
	FlagInstall bool
)

// manCmd generates manpage(s) for git-spend
var manCmd = &cobra.Command{
	Hidden: true,
	Use:    "man",
	Short:  locale.T("CommandManSummary"),
	Long:   locale.T("CommandManDescription"),
	Run: func(cmd *cobra.Command, args []string) {
		header := &doc.GenManHeader{
			Title:   "git-spend",
			Section: fmt.Sprintf("%d", manSection),
			Manual:  "⌛",
			Source:  "git-spend man",
		}

		outputDir := FlagOutput
		if FlagInstall && outputDir == "." {
			outputDir = fmt.Sprintf(manPath, manSection)
		}

		err := os.MkdirAll(outputDir, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}

		err = doc.GenManTree(rootCmd, header, outputDir)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(manCmd)
	manCmd.Flags().StringVar(
		&FlagOutput,
		"output",
		".",
		locale.T("CommandManFlagOutput"),
	)
	manCmd.Flags().BoolVar(
		&FlagInstall,
		"install",
		false,
		locale.T("CommandManFlagInstall"),
	)
}
