package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"sort"
	"sync"
)

type t struct {
	name *string
	id   int
}

func main() {

	ctx := context.Background()
	eg, ctx := errgroup.WithContext(ctx)
	result := make([]int, 0, 2)
	resultS := make([]string, 0, 2)
	input := make([]t, 2, 2)
	a := "hello"
	b := "bye"
	input[0] = t{
		name: &a,
		id:   1,
	}
	input[1] = t{
		name: &b,
		id:   2,
	}

	var mu sync.Mutex
	for _, i := range input {
		i := i
		//iS := *i.name
		eg.Go(func() error {
			mu.Lock()
			result = append(result, i.id)
			resultS = append(resultS, *i.name)
			mu.Unlock()
			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		fmt.Errorf("failed to create one or more volumes: %w", err)
	}
	fmt.Println(result, resultS)
}

func sortTheStudents(score [][]int, k int) [][]int {

	sort.Slice(score, func(i, j int) bool {
		return score[i][k] < score[j][k]
	})
	return score
}
