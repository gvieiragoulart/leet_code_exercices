package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	for i, number := range nums {
		for j, number2 := range nums {
			if i != j && number+number2 == target {
				return []int{i, j}
			}
		}
	}

	return []int{}
}

func main() {
	result := twoSum([]int{2, 7, 11, 15}, 9)
	result2 := twoSum([]int{3, 2, 4}, 6)
	result3 := twoSum([]int{3, 3}, 6)

	fmt.Println(result, result2, result3)
}
