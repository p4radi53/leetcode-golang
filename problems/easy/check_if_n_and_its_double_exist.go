package easy

// 1346. Check If N and Its Double Exist
func checkIfExist(arr []int) bool {
	hashmap := map[int]struct{}{}

	for _, v := range arr {
		if _, ok := hashmap[v * 2]; ok {
			return true
		}
		if _, ok := hashmap[v / 2]; ok && v % 2 == 0 {
			return true
		}

		hashmap[v] = struct{}{}
	}

	return false
}
