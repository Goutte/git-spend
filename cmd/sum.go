package cmd

import (
	"fmt"
	"github.com/goutte/gitime/gitime"
	"github.com/tsuyoshiwada/go-gitlog"
	"log"

	"github.com/spf13/cobra"
)

var (
	FlagMinutes bool
	FlagHours   bool
)

// sumCmd represents the sum command
var sumCmd = &cobra.Command{
	Use:   "sum",
	Short: "Sum /spent time recorded in commit messages",
	Long: `The commit messages of the currently checked out branch of the git repository of the current working directory will be read and their /spend and /spent directives will be parsed and summed.
You can also get a raw number in a specific unit:

    gitime sum --minutes
`,
	Run: func(cmd *cobra.Command, args []string) {
		ts := sum().Normalize()
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
	return out
}

func sum() *gitime.TimeSpent {
	git := gitlog.New(&gitlog.Config{})

	commits, err := git.Log(nil, nil)
	if err != nil {
		log.Fatalln("Cannot read git log:", err)
	}

	ts := &gitime.TimeSpent{}
	for _, commit := range commits {
		ts.Add(gitime.CollectTimeSpent(commit.Subject))
		ts.Add(gitime.CollectTimeSpent(commit.Body))
	}

	return ts
}

func init() {
	rootCmd.AddCommand(sumCmd)
	addFormatFlags(sumCmd)
}

func addFormatFlags(command *cobra.Command) {
	command.Flags().BoolVarP(
		&FlagMinutes,
		"minutes",
		"",
		false,
		"Show sum in minutes",
	)
	command.Flags().BoolVarP(
		&FlagHours,
		"hours",
		"",
		false,
		"Show sum in hours",
	)
}
