package question1

func Permutation(str string) []string {
	if len(str) <= 1 {
		return []string{str}
	}
	var result []string
	for i := 0; i < len(str); i++ {
		remaining := str[:i] + str[i+1:]
		permutations := Permutation(remaining)
		for _, perm := range permutations {
			result = append(result, string(str[i])+perm)
		}
	}
	result = removeDuplicate(result)
	return result
}

func removeDuplicate(strs []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range strs {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
