// 21. 文字列から重複する文字を表示するにはどうすればよいですか？
// How to Find Duplicate Characters on String - Java Programming Problems
// https://www.java67.com/2014/03/how-to-find-duplicate-characters-in-String-Java-program.html

package main

import (
	"fmt"
	"strings"
)

func printDuplicateCharacters(word string) {
	var characters []string
	characters = strings.Split(word, "")
	charMap := map[string]int{}

	for _, ch := range characters {
		if _, ok := charMap[ch]; ok {
			charMap[ch] = charMap[ch] + 1
		} else {
			charMap[ch] = 1
		}
	}

	fmt.Printf("List of duplicate characters in String '%s' \n", word)
	for key, val := range charMap {
		if val > 1 {
			fmt.Printf("%s : %d\n", key, val)
		}
	}
}

func main() {
	printDuplicateCharacters("Programming")
	printDuplicateCharacters("Combination")
	printDuplicateCharacters("golang")
}
