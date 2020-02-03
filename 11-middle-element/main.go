// 11. 1つのパスの中で片方向連結リストの真ん中の要素を探すにはどうすればよいですか？
// How to Find Middle Element of Linked List in Java in Single Pass
// https://javarevisited.blogspot.com/2012/12/how-to-find-middle-element-of-linked-list-one-pass.html

package main

import (
	"container/list"
	"fmt"
)

func main() {

	l := list.New()
	l.PushBack(0) // head node
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	l.PushBack(4)

	current := l.Front()
	middle := current
	length := 0

	for current.Next() != nil {
		length++
		if length%2 == 0 {
			middle = middle.Next()
		}
		current = current.Next()
	}

	if length%2 == 1 {
		middle = middle.Next()
	}

	fmt.Printf("length of List: %d\n", length)
	fmt.Printf("middle element of List: %d\n", middle.Value)
}
