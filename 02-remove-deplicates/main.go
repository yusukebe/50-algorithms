// 2. 与えられた整数の配列において重複した数字を探すにはどうすればよいですか？
// https://javarevisited.blogspot.com/2014/01/how-to-remove-duplicates-from-array-java-without-collection-API.html

package main

import "sort"

func removeDeplicates(numbers []int) []int {
	sort.Ints(numbers)

	var result []int
	previous := numbers[0]
	result = append(result, previous)

	for _, number := range numbers {
		if previous != number {
			result = append(result, number)
		}
		previous = number
	}

	return result
}
