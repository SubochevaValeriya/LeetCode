package main

import "testing"

func Test_minimizeArrayValue(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{{name: "1", args: args{nums: []int{3, 7, 1, 6}}, want: 5},
		{name: "2", args: args{nums: []int{13, 13, 20, 0, 8, 9, 9}}, want: 16},
		{name: "3", args: args{nums: []int{6, 9, 3, 8, 14}}, want: 8},
		{name: "4", args: args{nums: []int{4, 7, 2, 2, 9, 19, 16, 0, 3, 1}}, want: 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimizeArrayValue(tt.args.nums); got != tt.want {
				t.Errorf("minimizeArrayValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
