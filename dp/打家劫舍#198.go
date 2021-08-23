package dp

// func init() {
// 	fmt.Println(rob([]int{2, 3, 2}))
// }

//
func rob(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	if length == 1 {
		return nums[0]
	}
	var dp = make([]int, length)
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < length; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}
	return dp[length-1]
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
