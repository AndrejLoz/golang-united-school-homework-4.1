package reverse_string

//Implement a function, that reverses the input string

// Consider that the string may have international symbols, emojis and consist of several lines

func ReverseString(input string) (output string) {
	runes := []rune(input)

	for i := len(runes) - 1; i >= 0; i-- {

		output += string(runes[i])
	}

	return output
}
