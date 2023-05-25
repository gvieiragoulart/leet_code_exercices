package main

import (
	"fmt"
)

func romanToInt(s string) int {
	newArray := []int{}
	result := 0

	for _, char := range s {
		newArray = append(newArray, getValue(string(char)))
	}

	for i, number := range newArray {
		if i+1 < len(newArray) && newArray[i] < newArray[i+1] {
			newArray[i+1] -= number
		} else {
			result += number
		}
	}

	return result
}

func getValue(char string) int {
	switch char {
	case "I":
		return 1
	case "V":
		return 5
	case "X":
		return 10
	case "L":
		return 50
	case "C":
		return 100
	case "D":
		return 500
	case "M":
		return 1000
	}

	return 0
}

func main() {
	result := romanToInt("III")
	result2 := romanToInt("LVIII")
	result3 := romanToInt("MCMXCIV")

	fmt.Println(result, result2, result3)
}

// Symbol       Value
// I             1
// V             5
// X             10
// L             50
// C             100
// D             500
// M             1000
