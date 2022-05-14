package reverse_string

func ReverseString(input string) (output string) {
	tmp := []rune(input)
	strLen := len(tmp)
	// output = input

	for i, l := 0, strLen/2; i < l; i++ {
		tmp[strLen-i-1], tmp[i] = tmp[i], tmp[strLen-i-1]
		//output[i] = input[strLen-i-1]
	}

	return string(tmp)
}
