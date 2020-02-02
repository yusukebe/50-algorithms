// 3. ソートされていない整数の配列から最大値と最小値を探すにはどうすればよいですか？
// https://www.java67.com/2014/02/how-to-find-largest-and-smallest-number-array-in-java.html

package main

import (
	"math"
)

type Result struct {
	largest, smallest int
}

func largestAndSamallest(numbers []int) Result {
	largest := math.MinInt32  // XXX
	smallest := math.MaxInt32 // XXX

	for _, number := range numbers {
		if number > largest {
			largest = number
		} else if number < smallest {
			smallest = number
		}
	}
	return Result{largest: largest, smallest: smallest}
}
