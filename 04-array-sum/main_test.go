package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetPairs(t *testing.T) {
	tests := []struct {
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

	for _, tt := range tests {
		got := getPairs(tt.array, tt.sum)
		if !reflect.DeepEqual(tt.expect, got) {
			t.Errorf("Expect output to %s, but %s\n", fmt.Sprint(tt.expect), fmt.Sprint(got))
		}
	}
}
