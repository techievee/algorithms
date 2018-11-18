package main

import (
	"fmt"
	"math"
)

func main() {

	const a string = "pale"
	const b string = "bale"

	//Convert to rune as we are more converned about the unicode character
	ra := []rune(a)
	rb := []rune(b)

	if CheckOneDifference(ra, rb) == true {
		fmt.Printf("TRUE")
	} else {
		fmt.Printf("FALSE")
	}

}

/* CheckOneDifference ..*/
func CheckOneDifference(a []rune, b []rune) bool {

	var runeLarger []rune
	var runeSmaller []rune

	//If the difference is more than one then return false immediately
	if math.Abs(float64(len(a)-len(b))) > 1 {
		return false
	}

	//Find which is larger and shorter string. And assign the variables accordingle.
	//We will interate only through the larger string
	if len(a) >= len(b) {
		runeLarger = a
		runeSmaller = b

	} else {
		runeLarger = b
		runeSmaller = a

	}

	//Default assignment
	i, j := 0, 0
	oneChange := false
	deletion := true

	//Deletion variable contains whether varia one character ble is inserted/deleted or same
	if len(runeLarger) == len(runeSmaller) {
		deletion = false
	}

	for i = 0; i < len(runeLarger) && j < len(runeSmaller); i++ {

		if runeLarger[i] == runeSmaller[j] {
			//i gets incremented via loop, increment only j
			j++
			continue
		} else { //Mismatch case
			if oneChange == true {
				//ALready one variable is changed
				return false
			} else {
				//First change set the variable
				oneChange = true
				//If deletion case, one character need to be moved
				if deletion == false {
					j++
				}

			}
		}

	}

	// If still some i value exist it denotes that all the values are chrekced and last character is mismatch
	//Check whether one edit and take appropriate action
	if i < len(runeLarger) {
		if oneChange == true {
			return false
		} else {
			return true
		}

	}

	return true

}
