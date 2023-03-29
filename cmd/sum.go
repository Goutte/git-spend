package cmd

import (
	"fmt"
	"github.com/goutte/git-spend/gitime"
	"github.com/goutte/git-spend/gitime/reader"
	"github.com/spf13/cobra"
)

const (
	FlagTargetDefault = "."
)

var (
	FlagAuthors      []string
	FlagTarget       string
	FlagStdin        bool
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
from the commit messages in the current directory's git repository.

The default target is the current working directory, '.',
but you may specify another target using the --target flag:

	git-spend sum --target <some versioned dir with commits> 

You can also get a raw number in a specific unit:

    git spend sum --minutes

You can also restrict to some commit authors, by name or email:

    git spend sum --author=Alice --author=bob@pop.net --author=Eve

You can restrict to a range of commits, using a commit hash, a tag, or even HEAD~N.

	git spend sum --since <ref> --until <ref>

For example, to get the time spent on the last 15 commits :

	git spend sum --since HEAD~15

Or the time spent on a tag since the previous tag :

	git spend sum --since 0.1.0 --until 0.1.1

You can also use dates and datetimes, but remember to quote them:

	git spend sum --since 2023-03-21
	git spend sum --since "2023-03-21 13:37:00"

`,
	Run: func(cmd *cobra.Command, args []string) {
		ts := Sum().Normalize()
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

func Sum() *gitime.TimeSpent {
	var gitLog string
	if FlagStdin {
		if len(FlagAuthors) > 0 {
			fail(`Flag --author is not supported with --stdin parsing.
Meanwhile, you can use --author on git log, like so:

    git log --author Bob > log.log && cat log.log | git-spend sum`)
		}
		if FlagExcludeMerge {
			fail(`Flag --no-merges is not supported with --stdin parsing.
Meanwhile, you can use --no-merges on git log, like so:

    git log --no-merges > log.log && cat log.log | git-spend sum`)
		}
		if FlagSince != "" {
			fail(`Flag --since is not supported with --stdin parsing.`)
		}
		if FlagUntil != "" {
			fail(`Flag --until is not supported with --stdin parsing.`)
		}
		if FlagTarget != FlagTargetDefault {
			fail(`Flag --target is not supported with --stdin parsing.`)
		}
		gitLog = reader.ReadStdin()
	} else {
		gitLog = reader.ReadGitLog(FlagAuthors, FlagExcludeMerge, FlagSince, FlagUntil, FlagTarget)
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

func addTargetFlags(command *cobra.Command) {
	command.Flags().StringVar(
		&FlagTarget,
		"target",
		FlagTargetDefault,
		"target this directory instead of the working directory (.)",
	)
	command.Flags().BoolVar(
		&FlagStdin,
		"stdin",
		false,
		"read stdin instead of target's git log",
	)
}

func init() {
	rootCmd.AddCommand(sumCmd)
	addFormatFlags(sumCmd)
	addFilterFlags(sumCmd)
	addTargetFlags(sumCmd)
}
