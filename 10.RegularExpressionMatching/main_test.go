package main

import "testing"

func Test_isMatch(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "c*a*b", args: args{
			s: "aab",
			p: "c*a*b",
		}, want: true},
		{name: "c*a*bb", args: args{
			s: "aab",
			p: "c*a*bb",
		}, want: false},
		{name: "mississippi", args: args{
			s: "mississippi",
			p: "mis*is*p*.",
		}, want: false},
		{name: "mississippiTrue", args: args{
			s: "mississippi",
			p: "mis*is*ip*.",
		}, want: true},
		{name: "a", args: args{
			s: "aa",
			p: "a",
		}, want: false},
		{name: "aaa", args: args{
			s: "aaa",
			p: "ab*a*c*a",
		}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMatch(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("isMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
