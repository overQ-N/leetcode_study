package main

func isValid(s string) bool {
	arr := []rune{}
	m := map[rune]rune{'{': '}', '(': ')', '[': ']'}
	for _, v := range s {
		if value, ok := m[v]; ok {
			arr = append(arr, value)
		} else {

			if len(arr) > 0 && arr[len(arr)-1] == v {
				arr = arr[:len(arr)-1]
			} else {
				return false
			}
		}
	}
	return len(arr) == 0
}
