package string

import (
	"sort"
)

// g 5 6 7 8 8 9  11 11
// s 5 6 7 9 9 10 10 10
func findContentChildren(g []int, s []int) int {
	sort.Ints(g[:])
	sort.Ints(s[:])

	sIndex, gIndex := 0, 0
	counter := 0
    
	for gIndex < len(g) && sIndex < len(s){
		if s[sIndex] >= g[gIndex] {
			counter++
			sIndex++
			gIndex++
		} else {
			sIndex++
		}
	}
	return counter
}
