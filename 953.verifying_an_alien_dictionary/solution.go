package verifying_an_alien_dictionary

func isAlienSorted(words []string, order string) bool {
	if len(words) < 2 {
		return true
	}
	m := make(map[rune]int)
	for i, v := range order {
		m[v] = i
	}

	// for i := 0; i < words; i++ {

	// }
	return false
}
