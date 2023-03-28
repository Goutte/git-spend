package cmd

import (
	"fmt"
	"github.com/goutte/gitime/gitime"
	"github.com/goutte/gitime/gitime/reader"
	"github.com/spf13/cobra"
	"log"
)

var (
	FlagAuthors      []string
	FlagSince        string
	FlagUntil        string
	FlagMinutes      bool
	FlagHours        bool
	FlagDays         bool
	FlagWeeks        bool
	FlagMonths       bool
	FlagExcludeMerge bool
)

var sumCmd = &cobra.Command{
	Use:   "sum",
	Short: "Sum /spent time recorded in commit messages",
	Long: `The /spend and /spent directives will be parsed and summed
from the commit messages of the currently checked out branch
of the git repository of the current working directory.

You can also get a raw number in a specific unit:

    gitime sum --minutes

You can also restrict to some commit authors, by name or email:

    gitime sum --author=Alice --author=bob@pop.net --author=Eve

`,
	Run: func(cmd *cobra.Command, args []string) {
		ts := Sum(FlagAuthors, FlagExcludeMerge, FlagSince, FlagUntil).Normalize()
		fmt.Println(formatTimeSpent(ts))
	},
}

func formatTimeSpent(ts *gitime.TimeSpent) string {
	out := ""
	if FlagMinutes {
		out = fmt.Sprintf("%d", ts.ToMinutes())
	} else if FlagHours {
		out = fmt.Sprintf("%d", ts.ToHours())
	} else if FlagDays {
		out = fmt.Sprintf("%d", ts.ToDays())
	} else if FlagWeeks {
		out = fmt.Sprintf("%d", ts.ToWeeks())
	} else if FlagMonths {
		out = fmt.Sprintf("%d", ts.ToMonths())
	} else {
		out = ts.String()
	}
	if out == "" {
		out = "No time-tracking directives /spend or /spent found in commits."
	}
	return out
}

func Sum(onlyAuthors []string, excludeMerge bool, since string, until string) *gitime.TimeSpent {
	var gitLog string
	if reader.DoesStdinHaveData() {
		if len(onlyAuthors) > 0 {
			log.Fatalln(`Flag --author is not supported with stdin parsing.
Meanwhile, you can use --author on git log, like so:

    git log --author Bob > log.log && cat log.log | gitime sum`)
		}
		if excludeMerge {
			log.Fatalln(`Flag --no-merges is not supported with stdin parsing.
Meanwhile, you can use --no-merges on git log, like so:

    git log --no-merges > log.log && cat log.log | gitime sum`)
		}
		if FlagSince != "" {
			log.Fatalln(`Flag --since is not supported with stdin parsing.`)
		}
		if FlagUntil != "" {
			log.Fatalln(`Flag --until is not supported with stdin parsing.`)
		}
		gitLog = reader.ReadStdin()
	} else {
		gitLog = reader.ReadGitLog(onlyAuthors, excludeMerge, since, until, ".")
	}

	return gitime.CollectTimeSpent(gitLog)
}

func addFormatFlags(command *cobra.Command) {
	command.Flags().BoolVarP(
		&FlagMinutes,
		"minutes",
		"",
		false,
		"show sum in minutes",
	)
	command.Flags().BoolVarP(
		&FlagHours,
		"hours",
		"",
		false,
		"show sum in hours",
	)
	command.Flags().BoolVarP(
		&FlagDays,
		"days",
		"",
		false,
		"show sum in days",
	)
	command.Flags().BoolVarP(
		&FlagWeeks,
		"weeks",
		"",
		false,
		"show sum in weeks",
	)
	command.Flags().BoolVarP(
		&FlagMonths,
		"months",
		"",
		false,
		"show sum in months",
	)

	command.MarkFlagsMutuallyExclusive(
		"months",
		"weeks",
		"days",
		"hours",
		"minutes",
	)
}

func addFilterFlags(command *cobra.Command) {
	command.Flags().StringArrayVar(
		&FlagAuthors,
		"author",
		[]string{},
		"only use commits by these authors (can be repeated)",
	)
	command.Flags().BoolVar(
		&FlagExcludeMerge,
		"no-merges",
		false,
		"ignore merge commits",
	)
	command.Flags().StringVar(
		&FlagSince,
		"since",
		"",
		"only use commits after this ref (exclusive)",
	)
	command.Flags().StringVar(
		&FlagUntil,
		"until",
		"",
		"only use commits before this ref (inclusive)",
	)
}

func init() {
	rootCmd.AddCommand(sumCmd)
	addFormatFlags(sumCmd)
	addFilterFlags(sumCmd)
}
