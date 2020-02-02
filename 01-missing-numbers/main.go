// 1. 1から100までの与えられた整数の配列の中から足りない数字を探すにはどうすればよいですか？
// https://javarevisited.blogspot.com/2014/11/how-to-find-missing-number-on-integer-array-java.html

package main

import "github.com/willf/bitset"

func findMissingNumbers(numbers []int, count int) []int {
	missingCount := count - len(numbers)
	bitset := bitset.New(uint(count))

	for _, number := range numbers {
		bitset.Set(uint(number - 1))
	}

	var lastMissingIndex uint = 0
	var missingNumbers []int

	for i := 0; i < missingCount; i++ {
		lastMissingIndex, _ = bitset.NextClear(lastMissingIndex)
		lastMissingIndex = lastMissingIndex + 1
		missingNumbers = append(missingNumbers, int(lastMissingIndex))
	}

	return missingNumbers
}
