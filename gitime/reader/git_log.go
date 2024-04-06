package reader

import (
	"fmt"
	"github.com/tsuyoshiwada/go-gitlog"
	"os"
	"os/exec"
)

// ReadGitLog reads the git log of the repository of the specified directpry
func ReadGitLog(onlyAuthors []string, excludeMerge bool, since string, until string, directory string) string {
	git := gitlog.New(&gitlog.Config{
		Path: directory,
	})
	rev := getRevArgsFromFlags(since, until)
	params := &gitlog.Params{
		IgnoreMerges: excludeMerge,
	}
	commits, err := git.Log(rev, params)
	if exitError, isExitError := err.(*exec.ExitError); isExitError {
		fmt.Println("git command unsuccessful:", err, "—", exitError.Stderr)
		os.Exit(exitError.ExitCode())
	}
	if err != nil {
		fmt.Println("cannot read git log:", err)
		os.Exit(1)
	}

	s := ""
	for _, commit := range commits {
		if !isCommitByAnyAuthor(commit, onlyAuthors) {
			continue
		}

		// We read from the raw body because some newlines are eaten when separating subject an body.
		// My non-tech friend commits without separating subject and body, like this:
		//   > style: something amazing
		//   > /spent 0.5h
		// … and the "/spend 0.5h" ends up at the end of the Subject, without newline.
		s += commit.RawBody + "\n"
		// We also read from the note, and it might or might not be correct.
		s += commit.Note + "\n"
	}

	return s
}

func getRevArgsFromFlags(since string, until string) gitlog.RevArgs {
	var rev gitlog.RevArgs = nil
	if since != "" {
		sinceTime := parseTimePerhaps(since)

		if until != "" {
			untilTime := parseTimePerhaps(until)

			if untilTime != nil {
				if sinceTime != nil {
					rev = &gitlog.RevTime{
						Since: *sinceTime,
						Until: *untilTime,
					}
				} else {
					fmt.Println("unsupported mix of dates and refs in --until and --since")
					os.Exit(1)
				}
			} else {
				if sinceTime == nil {
					rev = &gitlog.RevRange{
						New: until,
						Old: since,
					}
				} else {
					fmt.Println("unsupported mix of dates and refs in --since and --until")
					os.Exit(1)
				}
			}
		} else {
			if sinceTime != nil {
				rev = &gitlog.RevTime{
					Since: *sinceTime,
				}
			} else {
				rev = &gitlog.RevRange{
					New: "HEAD",
					Old: since,
				}
			}
		}
	} else {
		if until != "" {
			untilDate := parseTimePerhaps(until)
			if untilDate != nil {
				rev = &gitlog.RevTime{
					Until: *untilDate,
				}
			} else {
				rev = &gitlog.Rev{
					Ref: until,
				}
			}
		}
	}
	return rev
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
