//package main
//
//import (
//	"fmt"
//)
//
//func countConsecutiveChars(s string) map[rune]int {
//	counts := make(map[rune]int)
//	currentChar := '\x00'
//	currentCount := 0
//
//	for _, char := range s {
//		if char == currentChar {
//			currentCount++
//		} else {
//			if currentChar != '\x00' {
//				counts[currentChar] = currentCount
//			}
//			currentChar = char
//			currentCount = 1
//		}
//	}
//	if currentChar != '\x00' {
//		counts[currentChar] = currentCount
//	}
//
//	return counts
//}
//
//func main() {
//	str := "aabbbccccdhhd"
//	counts := countConsecutiveChars(str)
//	for char, count := range counts {
//		fmt.Printf("%c: %d\n", char, count)
//	}
//}
package main

import (
	"fmt"
)

func countConsecutiveCharsRecursive(s string) map[rune]int {
	counts := make(map[rune]int)
	if len(s) == 0 {
		return counts
	}
	countConsecutiveCharsRecursiveHelper(s, counts, 0, rune(s[0]), 1)
	return counts
}
//递归函数
//

func countConsecutiveCharsRecursiveHelper(s string, counts map[rune]int, index int, currentChar rune, currentCount int) {
	if index == len(s)-1 {
		counts[currentChar] = currentCount
		return
	}
	if rune(s[index+1]) == currentChar {
		countConsecutiveCharsRecursiveHelper(s, counts, index+1, currentChar, currentCount+1)
	} else {
		counts[currentChar] = currentCount
		countConsecutiveCharsRecursiveHelper(s, counts, index+1, rune(s[index+1]), 1)
	}
}

func main() {
	str := "aabbbccccddhjkll"
	counts := countConsecutiveCharsRecursive(str)
	for char, count := range counts {
		fmt.Printf("%c: %d\n", char, count)
	}
}
