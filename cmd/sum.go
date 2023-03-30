package cmd

import (
	"fmt"
	"github.com/goutte/git-spend/gitime"
	"github.com/goutte/git-spend/gitime/reader"
	"github.com/goutte/git-spend/locale"
	"github.com/spf13/cobra"
)

const (
	FlagTargetDefault = "."
)

var (
	FlagAuthors  []string
	FlagTarget   string
	FlagStdin    bool
	FlagSince    string
	FlagUntil    string
	FlagMinutes  bool
	FlagHours    bool
	FlagDays     bool
	FlagWeeks    bool
	FlagMonths   bool
	FlagNoMerges bool
)

var sumCmd = &cobra.Command{
	Use:   "sum",
	Short: locale.T("CommandSumSummary"),
	Long:  locale.T("CommandSumDescription"),
	Run: func(cmd *cobra.Command, args []string) {
		ts, err := Sum()
		if err != nil {
			fail(err.Error(), cmd)
		}
		if ts != nil {
			fmt.Println(formatTimeSpent(ts.Normalize()))
		}
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
		out = locale.T("CommandSumDescription")
	}
	return out
}

func Sum() (*gitime.TimeSpent, error) {
	var gitLog string
	if FlagStdin {
		if len(FlagAuthors) > 0 {
			return nil, fmt.Errorf(locale.T("CommandSumFailureStdinAuthors"))
		}
		if FlagNoMerges {
			return nil, fmt.Errorf(locale.T("CommandSumFailureStdinNoMerges"))
		}
		if FlagSince != "" {
			return nil, fmt.Errorf(locale.T("CommandSumFailureStdinSince"))
		}
		if FlagUntil != "" {
			return nil, fmt.Errorf(locale.T("CommandSumFailureStdinUntil"))
		}
		if FlagTarget != FlagTargetDefault {
			return nil, fmt.Errorf(locale.T("CommandSumFailureStdinTarget"))
		}
		gitLog = reader.ReadStdin()
	} else {
		gitLog = reader.ReadGitLog(FlagAuthors, FlagNoMerges, FlagSince, FlagUntil, FlagTarget)
	}

	return gitime.CollectTimeSpent(gitLog), nil
}

func addFormatFlags(command *cobra.Command) {
	command.Flags().BoolVarP(
		&FlagMinutes,
		"minutes",
		"",
		false,
		locale.T("CommandSumFlagMinutesHelp"),
	)
	command.Flags().BoolVarP(
		&FlagHours,
		"hours",
		"",
		false,
		locale.Tf("CommandSumFlagHoursHelp", gitime.MinutesInOneHour),
	)
	command.Flags().BoolVarP(
		&FlagDays,
		"days",
		"",
		false,
		locale.Tf("CommandSumFlagDaysHelp", gitime.HoursInOneDay),
	)
	command.Flags().BoolVarP(
		&FlagWeeks,
		"weeks",
		"",
		false,
		locale.Tf("CommandSumFlagWeeksHelp", gitime.DaysInOneWeek),
	)
	command.Flags().BoolVarP(
		&FlagMonths,
		"months",
		"",
		false,
		locale.Tf("CommandSumFlagMonthsHelp", gitime.WeeksInOneMonth),
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
		locale.T("CommandSumFlagAuthorsHelp"),
	)
	command.Flags().BoolVar(
		&FlagNoMerges,
		"no-merges",
		false,
		locale.T("CommandSumFlagNoMergesHelp"),
	)
	command.Flags().StringVar(
		&FlagSince,
		"since",
		"",
		locale.T("CommandSumFlagSinceHelp"),
	)
	command.Flags().StringVar(
		&FlagUntil,
		"until",
		"",
		locale.T("CommandSumFlagUntilHelp"),
	)
}

func addTargetFlags(command *cobra.Command) {
	command.Flags().StringVar(
		&FlagTarget,
		"target",
		FlagTargetDefault,
		locale.T("CommandSumFlagTargetHelp"),
	)
	command.Flags().BoolVar(
		&FlagStdin,
		"stdin",
		false,
		locale.T("CommandSumFlagStdinHelp"),
	)
}

func init() {
	rootCmd.AddCommand(sumCmd)
	sumCmd.Flags().SortFlags = false
	addTargetFlags(sumCmd)
	addFilterFlags(sumCmd)
	addFormatFlags(sumCmd)
}
