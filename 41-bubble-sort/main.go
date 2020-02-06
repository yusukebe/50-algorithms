package main

import (
	"fmt"
)

func bubbleSort(numbers []int) {
	fmt.Printf("Unsorted array in Go: %v\n", numbers)

	for i := 0; i < len(numbers); i++ {
		for j := len(numbers) - 1; j > i; j-- {
			if numbers[j] < numbers[j-1] {
				swap(&numbers, j, j-1)
			}
		}
	}

	fmt.Printf("Sorted array using Bubble sort algorithm : %v\n", numbers)
}

func swap(array *[]int, from int, to int) {
	temp := (*array)[from]
	(*array)[from] = (*array)[to]
	(*array)[to] = temp
}

func main() {
	bubbleSort([]int{20, 12, 45, 19, 91, 55})
	bubbleSort([]int{-1, 0, 1})
	bubbleSort([]int{-3, -9, -2, -1})
}
