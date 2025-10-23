package string

func generate(numRows int) [][]int {
	result := make([][]int, numRows)

	result[0] = []int{1}
	for i := 2; i <= numRows; i++ {
		newList := make([]int, i)
		for j := 0; j < i; j++ {
			if j == 0 || j == i-1 {
				newList[j] = 1
			} else {
				previousList := result[i-2]
				newList[j] = previousList[j-1] + previousList[j]
			}
		}
		result[i-1] = newList
	}

	return result
}
