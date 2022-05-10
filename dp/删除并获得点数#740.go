package dp

import "fmt"

func init() {
	return
	expect1 := deleteAndEarn([]int{3, 4, 2})
	fmt.Printf("期待值6，输出值:%v", expect1)
	expect2 := deleteAndEarn([]int{2, 2, 3, 3, 3, 4})
	fmt.Printf("期待值9，输出值:%v", expect2)
}

func deleteAndEarn(nums []int) int {

	var targetArr = make([]int, 10001)
	for i := 0; i < len(nums); i++ {
		targetArr[nums[i]] += nums[i]
	}
	var dp = make([]int, 10001)
	dp[0] = 0
	dp[1] = targetArr[1]
	for i := 2; i < len(targetArr); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+targetArr[i])
	}
	return dp[len(dp)-1]
}
