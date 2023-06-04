package main

import (
	"reflect"
	"testing"
)

func Test_findSmallestSetOfVertices(t *testing.T) {
	type args struct {
		n     int
		edges [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "[[1,2],[3,2],[1,3],[1,0],[0,2],[0,3]]", args: args{
			n:     4,
			edges: [][]int{{1, 2}, {3, 2}, {1, 3}, {1, 0}, {0, 2}, {0, 3}},
		}, want: []int{1}},
		{name: "[[1,3],[2,0],[2,3],[1,0],[4,1],[0,3]]", args: args{
			n:     5,
			edges: [][]int{{1, 3}, {2, 0}, {2, 3}, {1, 0}, {4, 1}, {0, 3}},
		}, want: []int{2, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSmallestSetOfVertices(tt.args.n, tt.args.edges); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findSmallestSetOfVertices() = %v, want %v", got, tt.want)
			}
		})
	}
}
