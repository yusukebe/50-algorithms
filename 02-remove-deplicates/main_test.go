package main

import (
	"testing"
	"reflect"
	"fmt"
)

func TestRemoveDeplicates(t *testing.T) {
	tests := []struct {
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

	for _, tt := range tests {
		got := removeDeplicates(tt.input)
		if !reflect.DeepEqual(tt.expect, got) {
			t.Errorf("Expect output to %s, but %s\n", fmt.Sprint(tt.expect), fmt.Sprint(got))
		}
	}
}
