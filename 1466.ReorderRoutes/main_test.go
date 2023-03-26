package main

import (
	"testing"
)

func Test_minReorder(t *testing.T) {
	type args struct {
		n           int
		connections [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{{name: "d", args: args{
		n:           5,
		connections: [][]int{{4, 3}, {2, 3}, {1, 2}, {1, 0}},
	}, want: 2},
		{name: "test2", args: args{
			n:           7,
			connections: [][]int{{0, 6}, {6, 3}, {1, 3}, {2, 1}, {4, 0}, {4, 5}},
		}, want: 3},
		{name: "test1", args: args{
			n:           6,
			connections: [][]int{{0, 1}, {1, 3}, {2, 3}, {4, 0}, {4, 5}},
		}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minReorder(tt.args.n, tt.args.connections); got != tt.want {
				t.Errorf("minReorder() = %v, want %v", got, tt.want)
			}
		})
	}
}
