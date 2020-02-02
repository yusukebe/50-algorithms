// 4. 合計すると与えられた数字と同じになる整数の配列のすべての組み合わせを探すにはどうすればよいですか？
// How to Find all Pairs in Array of Integers Whose sum is Equal to a Given Number
// https://javarevisited.blogspot.com/2014/08/how-to-find-all-pairs-in-array-of-integers-whose-sum-equal-given-number-java.html

package main

func getPairs(numbers []int, sum int) [][]int {
	var result [][]int

	set := map[int]int{}

	for _, number := range numbers {
		target := sum - number
		if _, ok := set[target]; !ok {
			set[number] = 1
		} else {
			result = append(result, []int{number, target})
		}
	}

	return result
}
