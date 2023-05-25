package main

import (
	"fmt"
)

func isPalindrome(x int) bool {
	invertedNumber := ""
	number := fmt.Sprintf("%d", x)

	if string(number[0]) != string(number[len(number)-1]) {
		return false
	}

	for i := 1; i <= len(number); i++ {
		invertedNumber += string(number[len(number)-i])
	}

	return (invertedNumber == number)
}

func main() {
	result := isPalindrome(12414201)

	fmt.Println(result)
}
