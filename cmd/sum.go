package cmd

import (
	"fmt"
	"github.com/goutte/gitime/gitime"
	"github.com/spf13/cobra"
	"github.com/tsuyoshiwada/go-gitlog"
	"io"
	"log"
	"os"
	"time"
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
		ts := Sum(FlagAuthors, FlagExcludeMerge).Normalize()
		fmt.Println(formatTimeSpent(ts))
	},
}

func doesStdinHaveData() bool {
	fileInfo, err := os.Stdin.Stat()
	if err != nil {
		return false
	}

	if os.Getenv("GITIME_NO_STDIN") == "1" {
		return false
	}

	// Both yield false positives in CI, careful
	//if (fileInfo.Mode() & os.ModeCharDevice) == 0 { // alternatively?
	if (fileInfo.Mode() & os.ModeNamedPipe) != 0 {
		return true
	}

	return false
}

func readStdin() string {
	stdin, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%s", stdin)
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

func parseTimePerhaps(input string) *time.Time {
	if input == "" {
		return nil
	}

	layouts := []string{
		time.RFC3339,
		time.DateTime,
		time.DateOnly,
		time.RFC822,
		time.RFC850,
	}

	for _, layout := range layouts {
		parse, err := time.Parse(layout, input)
		if err == nil {
			return &parse
		}
	}

	return nil
}

func getRevArgsFromFlags() gitlog.RevArgs {
	var rev gitlog.RevArgs = nil
	if FlagSince != "" {
		flagSinceDate := parseTimePerhaps(FlagSince)

		if FlagUntil != "" {
			flagUntilDate := parseTimePerhaps(FlagUntil)

			if flagUntilDate != nil {
				if flagSinceDate != nil {
					rev = &gitlog.RevTime{
						Since: *flagSinceDate,
						Until: *flagUntilDate,
					}
				} else {
					fmt.Println("you cannot mix dates and refs in --until and --since")
					os.Exit(1)
				}
			} else {
				if flagSinceDate != nil {
					fmt.Println("you cannot mix dates and refs in --since and --until")
					os.Exit(1)
				} else {
					rev = &gitlog.RevRange{
						New: FlagUntil,
						Old: FlagSince,
					}
				}
			}
		} else {
			if flagSinceDate != nil {
				rev = &gitlog.RevTime{
					Since: *flagSinceDate,
				}
			} else {
				rev = &gitlog.RevRange{
					New: "HEAD",
					Old: FlagSince,
				}
			}
		}
	} else {
		if FlagUntil != "" {
			flagUntilDate := parseTimePerhaps(FlagUntil)
			if flagUntilDate != nil {
				rev = &gitlog.RevTime{
					Until: *flagUntilDate,
				}
			} else {
				rev = &gitlog.Rev{
					Ref: FlagUntil,
				}
			}
		}
	}
	return rev
}

// ReadGitLog reads the git log of the repository of the specified directpry
func ReadGitLog(onlyAuthors []string, excludeMerge bool, directory string) string {
	git := gitlog.New(&gitlog.Config{
		Path: directory,
	})
	rev := getRevArgsFromFlags()
	params := &gitlog.Params{
		IgnoreMerges: excludeMerge,
	}
	commits, err := git.Log(rev, params)
	if err != nil {
		fmt.Println("Cannot read git log:", err)
		os.Exit(1)
		//log.Fatalln("Cannot read git log:", err)
	}

	s := ""
	for _, commit := range commits {
		if !isCommitByAnyAuthor(commit, onlyAuthors) {
			continue
		}

		s += commit.Subject + "\n"
		s += commit.Body + "\n"
	}

	return s
}

func Sum(onlyAuthors []string, excludeMerge bool) *gitime.TimeSpent {
	var gitLog string
	if doesStdinHaveData() {
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
		gitLog = readStdin()
	} else {
		gitLog = ReadGitLog(onlyAuthors, excludeMerge, ".")
	}

	return gitime.CollectTimeSpent(gitLog)
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
