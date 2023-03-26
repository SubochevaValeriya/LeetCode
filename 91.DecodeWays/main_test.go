package main

import "testing"

func Test_numDecodings(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{{name: "12", args: args{s: "12"}, want: 2},
		{name: "226", args: args{s: "226"}, want: 3},
		{name: "06", args: args{s: "06"}, want: 0},
		{name: "10", args: args{s: "10"}, want: 1},
		{name: "2206", args: args{s: "2206"}, want: 1},
		{name: "1123", args: args{s: "1123"}, want: 5},
		{name: "27", args: args{s: "27"}, want: 1}} // TODO: Add test cases.

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numDecodings(tt.args.s); got != tt.want {
				t.Errorf("numDecodings() = %v, want %v", got, tt.want)
			}
		})
	}
}
