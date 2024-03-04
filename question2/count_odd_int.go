package question2

func CountOddOccur(intArr []int) int {
	countMap := make(map[int]int)
	for _, num := range intArr {
		countMap[num]++
	}
	for key, count := range countMap {
		if count%2 != 0 {
			return key
		}
	}
	return 0
}
