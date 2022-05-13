package one_away_LCCI

// 一次编辑

func oneEditAway(first string, second string) bool {
	m, n := len(first), len(second)
	if m < n {
		return oneEditAway(second, first)
	}
	if m-n > 1 {
		return false
	}
	for i, char := range second {
		if first[i] != byte(char) {
			if m == n {
				return first[i+1:] == second[i+1:]
			}
			return first[i+1:] == second[i:]
		}
	}
	return true
}
