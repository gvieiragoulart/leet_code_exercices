package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func longestCommonPrefix(strs []string) string {
	result := ""

	if len(strs) == 0 {
		return result
	}

	for i := 0; i < len(strs[0]); i++ {
		for _, str := range strs {
			fmt.Println(str, string(str[i]), string(strs[0][i]))
			if len(str) <= i || strs[0][i] != str[i] {
				return result
			}
		}

		result += string(strs[0][i])
	}

	return result
}

func main() {
	result1 := longestCommonPrefix([]string{"flower", "flow", "flight"})
	result2 := longestCommonPrefix([]string{"dog", "racecar", "car"})

	fmt.Println(result1, result2)
}
