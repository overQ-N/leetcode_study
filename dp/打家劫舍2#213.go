package dp

// func init() {
// 	expect1 := rob2([]int{1, 2, 3})
// 	fmt.Printf("预期结果：3，输出:%v\n", expect1)

// 	expect2 := rob2([]int{2, 3, 2})
// 	fmt.Printf("预期结果：3，输出:%v\n", expect2)

// 	expect3 := rob2([]int{2, 1, 1, 2})
// 	fmt.Printf("预期结果：3，输出:%v\n", expect3)
// }

func rob2(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	if length == 1 {
		return nums[0]
	}
	dp1 := make([]int, length+2)
	dp1[0] = nums[0]
	dp1[1] = max(nums[0], nums[1])

	dp2 := make([]int, length+2)
	dp2[0] = 0
	dp2[1] = nums[1]
	for i := 2; i < length; i++ {
		dp1[i] = max(dp1[i-1], dp1[i-2]+nums[i])
		dp2[i] = max(dp2[i-1], dp2[i-2]+nums[i])
	}

	return max(dp1[length-2], dp2[length-1])
}
