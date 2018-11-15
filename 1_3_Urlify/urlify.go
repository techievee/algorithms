package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {

	//Testing the UTF COdepoint till 255
	a := "ÿ Dr John ò is a good doctor                                    "
	urlify(a)
	return
}

func urlify(url string) {
	println("\n Input String : ", url)

	//We need to get the string length excluding the end spaces availeble for reshuffle
	length := utf8.RuneCountInString(strings.Trim(url, " "))
	space := 0

	//Count spaces till the end of sentence
	for _, char := range strings.Trim(url, " ") {
		if strings.Compare(string(char), " ") == 0 {
			space++
		}
	}

	//Strings cannot be modified, they are immutable, Rune which is a Int32, that represents the UTF8 bytes are used
	runeString := []rune(url)

	for length > 0 {

		//Space into 2, in used as 2 space for " %2"
		//Already one space is available that would be used for "0"
		if runeString[length-1] == ' ' {
			runeString[(space*2)+(length-1)] = '0'
			runeString[(space*2)+(length-2)] = '2'
			runeString[(space*2)+(length-3)] = '%'
			space--
			length--

		} else {
			runeString[(space*2)+(length-1)] = runeString[(length - 1)]
			length--
		}
	}

	fmt.Printf("\n Output String : %s", string(runeString))

}
