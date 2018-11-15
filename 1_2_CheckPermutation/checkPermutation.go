package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	// Would work for any UTF Character in the Range 0-255
	const a = "aÿfdbe"
	const b = "abydef"
	const c = "afbÿde"

	//ÿ - 255 - Border cases
	// 0 - Control Character

	var flag = make([]int8, 256)

	if CheckStringPermutation(a, c, flag) == true {
		fmt.Printf("String a is the permutation of string b")
	} else {
		fmt.Printf("String a is not the permutation of string b")
	}

}

/* CheckStringPermutation .. */
func CheckStringPermutation(a string, b string, flag []int8) bool {

	//len cannot be used to count the utf8 string, it returns differnt size for other then ASCII character
	if utf8.RuneCountInString(a) != utf8.RuneCountInString(b) {
		return false
	}

	for _, rune := range a {
		value := int(rune)
		flag[value]++

	}

	for _, rune := range b {
		value := int(rune)
		flag[value]--
		if flag[value] < 0 {
			return false
		}
	}

	return true

}
