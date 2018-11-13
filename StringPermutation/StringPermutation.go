package main

func main() {

	input := "abcdef"
	println("%d", input)
	permuation("", input)
	println("..COMPLETED..")

}

func permuation(prefix string, str string) {
	if len(str) == 0 {
		println(prefix)
	}

	for i := 0; i < len(str); i++ {
		newPrefix := str[i : i+1]
		remainingString := str[0:i] + str[i+1:]
		permuation(prefix+newPrefix, remainingString)
	}
}
