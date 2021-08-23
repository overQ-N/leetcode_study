package dp

import (
	"math"
)

// func init() {
// 	cost := []int{10, 15, 20}
// 	fmt.Println(minCostClimbingStairs(cost))
// }
func minCostClimbingStairs(cost []int) int {
	l := len(cost)
	for i := 2; i < l; i++ {
		cost[i] += int(math.Min(float64(cost[i-1]), float64(cost[i-2])))
	}
	return int(math.Min(float64(cost[l-1]), float64(cost[l-2])))
}
