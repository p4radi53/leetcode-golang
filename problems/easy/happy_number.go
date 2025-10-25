package easy

func isHappy(n int) bool {
	var current int
	previous := n
	alreadySeen := map[int]struct{}{}

	for current != 1 {
		current = 0
		for previous != 0 {
			residual := previous % 10
			previous = previous / 10
			current += residual * residual
		}
		previous = current
		if _, ok := alreadySeen[current]; ok {
			return false
		} else {
			alreadySeen[current] = struct{}{}
		}

	}
	return true
}
