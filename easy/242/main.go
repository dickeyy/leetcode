package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	count := make(map[rune]int)
	for _, c := range s {
		count[c]++
	}

	for _, c := range t {
		count[c]--
	}

	for _, v := range count {
		if v != 0 {
			return false
		}
	}

	return true
}
