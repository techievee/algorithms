package main

func main() {

	const a string = "tactÿcoaÿ"
	const b string = "tactaceae"

	if CheckPermutatedPalindrome(b) == true {
		println("Permutation")
	} else {
		println("Not permutation")
	}

}

/* CheckPermutatedPalindrome.. */
func CheckPermutatedPalindrome(a string) bool {

	flag := make([]uint64, 4)

	//Fill the flag based on the string a
	for _, rune := range a {
		numChar := uint8(rune)
		arraySelection := numChar / 64
		posSelection := numChar % 64

		//Check the flag whether its set or not
		if (flag[arraySelection] & (1 << posSelection)) >= 1 { // Value is set
			// Reset the flag
			flag[arraySelection] &= ^(1 << posSelection)

		} else { //Flag is not set

			//Set the flag
			flag[arraySelection] |= (1 << posSelection)
		}

	}

	//Not whether one bit is set flag is set for all for uint64
	firstOdd := false
	for i := 0; i < 4; i++ {

		if flag[i] == 0 { //No odd
			continue
		}

		if (flag[i] & (flag[i] - 1)) == 0 {

			if firstOdd == true {
				//Already we encountered one add among the variable so make stop and return
				return false
			} else {
				//Out of 4 entries, this is first entry of odd
				firstOdd = true
			}

		} else {
			//Multiple odd entries in single flag
			return false

		}

	}

	return true

}
