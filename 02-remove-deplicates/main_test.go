package main

import (
	"testing"
	"reflect"
	"fmt"
)

func TestRemoveDeplicates(t *testing.T) {
	patterns := []struct {
		input  []int
		expect []int
	}{
		{
			[]int{1, 1, 2, 2, 3, 4, 5},
			[]int{1, 2, 3, 4, 5},
		},
		{
			[]int{1, 1, 1, 1, 1, 1, 1},
			[]int{1},
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7},
			[]int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			[]int{1, 2, 1, 1, 1, 1, 1},
			[]int{1, 2},
		},
	}

	for _, p := range patterns {
		got := removeDeplicates(p.input)
		if !reflect.DeepEqual(p.expect, got) {
			t.Errorf("Expect output to %s, but %s\n", fmt.Sprint(p.expect), fmt.Sprint(got))
		}
	}
}
