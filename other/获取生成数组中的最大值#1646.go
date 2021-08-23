package other

import (
	"fmt"
	"math"
)

func init() {
	fmt.Println(int(math.Floor(float64(3 / 2))))
	fmt.Println(getMaximumGenerated(0))
}
func getMaximumGenerated(n int) int {
	if n <= 1 {
		return n
	}
	res := make([]int, n+1)
	res[0] = 0
	res[1] = 1
	for i := 2; i <= n; i++ {
		// 偶数
		if i%2 == 0 {
			index := i / 2
			res[i] = res[index]
		} else {
			// 奇数
			index := int(math.Floor(float64(i / 2)))
			res[i] = res[index] + res[index+1]
		}
	}
	fmt.Println(res)
	max := 0
	for i := 0; i < n+1; i++ {
		if res[i] > max {
			max = res[i]
		}
	}
	return max
}
