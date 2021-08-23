package dp

// func init() {
// 	fmt.Println(climbStairs(4))
// }

/**
* 斐波那契数列：解法1：记忆法递归
* 超时了
 */
// func climbStairs(n int) int {
// 	var cache []int = make([]int, n+1)
// 	var dp = func(n int) int {
// 		if cache[n] > 0 {
// 			return cache[n]
// 		}
// 		if n == 1 {
// 			cache[n] = 1
// 		} else if n == 2 {
// 			cache[n] = 2
// 		} else {
// 			cache[n] = climbStairs(n-1) + climbStairs(n-2)
// 		}
// 		return cache[n]
// 	}
// 	return dp(n)
// }

/*
* 解法2：滚动数组法
 */
func climbStairs(n int) int {
	// if n == 1 {
	// 	return 1
	// }
	// if n == 2 {
	// 	return 2
	// }
	a, b, c := 0, 0, 1
	for i := 1; i <= n; i++ {
		a = b
		b = c
		c = a + b
	}
	return c
}
