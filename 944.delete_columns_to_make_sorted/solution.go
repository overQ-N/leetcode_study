package delete_columns_to_make_sorted

// 944. 删列造序
func minDeletionSize(strs []string) (res int) {
	if len(strs) == 0 {
		return 0
	}
	str := strs[0]
	length := len(strs)

	for i := 0; i < len(str); i++ {
		item := string(str[i])
		for j := 1; j < length; j++ {
			last := string(strs[j][i])
			if last < string(item[len(item)-1]) {
				res += 1
				break
			}
			item += last
		}
	}

	return res
}
