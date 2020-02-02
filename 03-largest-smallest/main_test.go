package main

import (
	"fmt"
	"math"
	"reflect"
	"testing"
)

func TestLargestAndSmallest(t *testing.T) {
	patterns := []struct {
		input  []int
		expect Result
	}{
		{
			[]int{-20, 34, 21, -87, 92, math.MaxInt32},
			Result{largest: math.MaxInt32, smallest: -87},
		},
		{
			[]int{10, math.MinInt32, -2},
			Result{largest: 10, smallest: math.MinInt32},
		},
		{
			[]int{math.MaxInt32, 40, math.MaxInt32},
			Result{largest: math.MaxInt32, smallest: 40},
		},
		{
			[]int{1, -1, 0},
			Result{largest: 1, smallest: -1},
		},
	}

	for _, p := range patterns {
		got := largestAndSamallest(p.input)
		if !reflect.DeepEqual(p.expect, got) {
			t.Errorf("Expect output to %s, but %s\n", fmt.Sprint(p.expect), fmt.Sprint(got))
		}
	}
}
