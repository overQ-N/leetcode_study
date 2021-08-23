package other

import (
	"math"
)

// func init() {
// 	ghosts := [][]int{{1, 0}, {0, 3}}
// 	target := []int{0, 1}
// 	fmt.Println(escapeGhosts(ghosts, target))
// }
func escapeGhosts(ghosts [][]int, target []int) bool {
	targetX := float64(target[0])
	targetY := float64(target[1])
	dis := int(math.Abs(targetX) + math.Abs(targetY))
	for _, v := range ghosts {
		if int(math.Abs(float64(v[0])-targetX)+math.Abs(float64(v[1])-targetY)) <= dis {
			return false
		}
	}
	return true
}
