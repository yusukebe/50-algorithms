package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFindMissingNumbers(t *testing.T) {
	tests := []struct {
		numbers    []int
		totalCount int
		expect     []int
	}{
		{
			[]int{1, 2, 3, 4, 6},
			6,
			[]int{5},
		},
		{
			[]int{1, 2, 3, 4, 6, 7, 9, 8, 10},
			10,
			[]int{5},
		},
		{
			[]int{1, 2, 3, 4, 6, 9, 8},
			10,
			[]int{5, 7, 10},
		},
		{
			[]int{1, 2, 3, 4, 9, 8},
			10,
			[]int{5, 6, 7, 10},
		},
	}

	for _, tt := range tests {
		got := findMissingNumbers(tt.numbers, tt.totalCount)
		if !reflect.DeepEqual(tt.expect, got) {
			t.Errorf("Expect output to %s, but %s\n", fmt.Sprint(tt.expect), fmt.Sprint(got))
		}
	}
}
