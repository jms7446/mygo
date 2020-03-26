package tool

func count(s string, codeCount map[rune]int) {
	for _, r := range s {
		codeCount[r]++
	}
}

func hasDupeRune(s string) bool {
	runeSet := map[rune]struct{}{}
	for _, r := range s {
		if _, exists := runeSet[r]; exists {
			return true
		}
		runeSet[r] = struct{}{}
	}
	return false
}
