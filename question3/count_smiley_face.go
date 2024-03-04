package question3

func CountSmileys(arr []string) int {
	count := 0
	for _, str := range arr {
		if len(str) < 4 {
			isValidSmiley := isValidSmiley(str)
			if isValidSmiley {
				count++
			}
		}
	}
	return count
}

func isValidSmiley(str string) bool {
	if len(str) < 2 || len(str) > 3 {
		return false
	}

	isValid := false

	isValidPrefix := str[0] == ':' || str[0] == ';'
	isValidSuffix := str[len(str)-1] == ')' || str[len(str)-1] == 'D'

	if len(str) == 2 {
		isValid = isValidPrefix && isValidSuffix
	}

	if len(str) == 3 {
		isValidMiddle := str[1] == '-' || str[1] == '~'
		isValid = isValidPrefix && isValidMiddle && isValidSuffix
	}

	return isValid

}
