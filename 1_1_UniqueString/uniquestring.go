package main

import (
	"fmt"
)

func main() {

	// Would work for any UTF Character in the Range 0-255
	const s = "tte6ü35ÿþþÿÿüÿÿþ"

	//ÿ - 255 - Border case
	// 0 - Control Character

	//Since we are concerned with 0-255 characters, we are making a Slice of UINT64 of 4 Nos
	//64X4 = 256 Entries, used to make a flag
	var flag = make([]uint64, 4)

	//Print String content
	fmt.Printf("Input String = %s\n", s)

	//rune of Go would automatically take care of the UTF8 Extra bytes when iterated via range
	for i, runeValue := range s {

		//Print flag variable before Setting
		fmt.Printf("[ %b, %b, %b, %b ]", flag[0], flag[1], flag[2], flag[3])
		arrayPos := uint(runeValue)

		if SetAndCheck(arrayPos, flag) == true {
			fmt.Printf("\nNot unique at position %d : %#U , %d : Array No: %d, Array Index : %d\n", i, runeValue, arrayPos, (arrayPos / 64), (arrayPos % 64))

			//Print flag variable after Setting
			fmt.Printf("[ %b, %b, %b, %b ]\n\n", flag[0], flag[1], flag[2], flag[3])
			continue

		}
		fmt.Printf("\nUnique at position %d : %#U , %d : Array No: %d, Array Index : %d\n", i, runeValue, arrayPos, (arrayPos / 64), (arrayPos % 64))
		//Print flag variable after Setting
		fmt.Printf("[ %b, %b, %b, %b ]\n\n", flag[0], flag[1], flag[2], flag[3])

	}

}

/* SetAndCheck functions checks the flag and if its already set then it returns true, if not set then it will set it and return false */
func SetAndCheck(code uint, flag []uint64) bool {

	var pos uint
	var arraySelection uint
	//Find the corresponding array
	arraySelection = code / 64

	//Find the position in that array
	pos = (code % 64)

	//Sets the first bit to 1, of the 64bit unsigned integer.
	var value uint64 = 0x8000000000000000

	if (flag[arraySelection] & (value >> pos)) > 0 {
		//Value Exist, so simply return true
		return true
	} else {
		//Value doesnt exist, so Set the flag and return false
		flag[arraySelection] |= value >> (pos)
		return false
	}
}
