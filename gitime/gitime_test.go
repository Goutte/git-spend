package gitime

import (
	"main/gitime"
	"testing"
)

func TestCollectTimeSpent(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{
			name: "No time is specified",
			args: args{
				message: "feat: implement something",
			},
			want: 0,
		},
		{
			name: "No time unit is specified → minutes",
			args: args{
				message: `feat: implement something

/spent 3
`,
			},
			want: 3,
		},
		{
			name: "/spend 1h",
			args: args{
				message: `feat: implement something

I did good.

/spend 1h
`,
			},
			want: 60,
		},
		{
			name: "/spend 30m",
			args: args{
				message: `/spend 30m`,
			},
			want: 30,
		},
		{
			name: "/spend 42 min",
			args: args{
				message: `/spend 42 min`,
			},
			want: 42,
		},
		{
			name: "/spend 42 minutes",
			args: args{
				message: `/spend 42 minutes`,
			},
			want: 42,
		},
		{
			name: "/spend 1 minute",
			args: args{
				message: `/spend 1 minute`,
			},
			want: 1,
		},
		{
			name: "/spend 1h 90m",
			args: args{
				message: `/spend 1h 90m`,
			},
			want: 150,
		},
		{
			name: "/spend 1d",
			args: args{
				message: `/spend 1d`,
			},
			want: 8 * 60,
		},
		{
			name: "/spend 1d2h10m",
			args: args{
				message: `/spend 1d2h10m`,
			},
			want: 610,
		},
		{
			name: "/spent 1.5h",
			args: args{
				message: `/spent 1.5h`,
			},
			want: 90,
		},
		{
			name: "/spent .5h",
			args: args{
				message: `/spent .5h`,
			},
			want: 30,
		},
		{
			name: "/spent 2. h",
			args: args{
				message: `/spent 2. h`,
			},
			want: 120,
		},
		{
			name: "/spent 1h30",
			args: args{
				message: `/spent 1h30`,
			},
			want: 90,
		},
		{
			name: "Multiple /spend directives are cumulated",
			args: args{
				message: `feat: everything
/spend 1d
/spend 2h
/spend 2h30m
`,
			},
			want: 750,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gitime.CollectTimeSpent(tt.args.message).ToMinutes(); got != tt.want {
				t.Errorf("CollectTimeSpent(%s) = %v, want %v", tt.args.message, got, tt.want)
			}
		})
	}
}
