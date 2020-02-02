package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetPairs(t *testing.T) {
	patterns := []struct {
		array  []int
		sum    int
		expect [][]int
	}{
		{
			[]int{0, 14, 0, 4, 7, 8, 3, 5, 7},
			11,
			[][]int{
				{7, 4},
				{3, 8},
				{7, 4},
			},
		},
		{
			[]int{10, 9, 5, 9, 0, 10, 2, 10, 1, 9},
			12,
			[][]int{
				{2, 10},
			},
		},
	}

	for _, p := range patterns {
		got := getPairs(p.array, p.sum)
		if !reflect.DeepEqual(p.expect, got) {
			t.Errorf("Expect output to %s, but %s\n", fmt.Sprint(p.expect), fmt.Sprint(got))
		}
	}
}
