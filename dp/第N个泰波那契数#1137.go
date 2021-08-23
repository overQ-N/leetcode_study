package dp

func init() {
	// res := tribonacci(4)
	// fmt.Println("res:", res)
}

func tribonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n <= 2 {
		return 1
	}
	/*
	* 滚动数组,记忆前累加值
	 */
	a, b, c, d := 0, 0, 1, 1
	for i := 3; i <= n; i++ {
		a = b
		b = c
		c = d
		d = a + b + c
	}
	return d
}
