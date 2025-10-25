package easy

// prev2 = to all solutions you append '2'
// prev = to all solutions you append '1'
// current = prev + prev2
func climbStairs(n int) int {
	if n <= 3 {
		return n
	}
	prev := 3
	prev2 := 2
	var current int

	for i := 4; i <= n; i++ {
		current = prev + prev2

		prev2 = prev
		prev = current
	}

	return current
}
