package cmd

import (
	"fmt"
	"github.com/goutte/gitime/gitime"
	"github.com/spf13/cobra"
	"github.com/tsuyoshiwada/go-gitlog"
	"log"
)

var (
	FlagAuthors []string
	FlagMinutes bool
	FlagHours   bool
)

// sumCmd represents the sum command
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
		ts := sum(FlagAuthors).Normalize()
		fmt.Println(formatTimeSpent(ts))
	},
}

func formatTimeSpent(ts *gitime.TimeSpent) string {
	out := ""
	if FlagMinutes {
		out = fmt.Sprintf("%d", ts.ToMinutes())
	} else if FlagHours {
		out = fmt.Sprintf("%d", ts.ToHours())
	} else {
		out = ts.String()
	}
	if out == "" {
		out = "No time-tracking directives /spend or /spent found in commits."
	}
	return out
}

func sum(onlyAuthors []string) *gitime.TimeSpent {
	git := gitlog.New(&gitlog.Config{})

	commits, err := git.Log(nil, nil)
	if err != nil {
		log.Fatalln("Cannot read git log:", err)
	}

	ts := &gitime.TimeSpent{}
	for _, commit := range commits {
		if !isCommitByAnyAuthor(commit, onlyAuthors) {
			continue
		}
		ts.Add(gitime.CollectTimeSpent(commit.Subject))
		ts.Add(gitime.CollectTimeSpent(commit.Body))
	}

	return ts
}

func isCommitByAnyAuthor(commit *gitlog.Commit, authors []string) bool {
	if len(authors) == 0 {
		return true
	}

	if commit.Author == nil {
		return false
	}

	for _, author := range authors {
		if commit.Author.Name == author {
			return true
		}
		if commit.Author.Email == author {
			return true
		}
	}

	return false
}

func init() {
	rootCmd.AddCommand(sumCmd)
	addFormatFlags(sumCmd)
	addFilterFlags(sumCmd)
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
}

func addFilterFlags(command *cobra.Command) {
	command.Flags().StringArrayVarP(
		&FlagAuthors,
		"author",
		"",
		[]string{},
		"only use commits by these authors (can be repeated)",
	)
}
