package main

/*
 *   Stats:
 *   Runtime: <3ms (beats 63.08% of Go solutions)
 *   Memory: 4.42mb (beats 39.66% of Go solutions)
 */

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	start := 0
	maxLength := 1

	// Check each character as a potential center
	for i := 0; i < len(s); i++ {
		// Check odd length palindromes
		left := i
		right := i
		for left >= 0 && right < len(s) && s[left] == s[right] {
			if right-left+1 > maxLength {
				start = left
				maxLength = right - left + 1
			}
			left--
			right++
		}

		// Check even length palindromes
		left = i
		right = i + 1
		for left >= 0 && right < len(s) && s[left] == s[right] {
			if right-left+1 > maxLength {
				start = left
				maxLength = right - left + 1
			}
			left--
			right++
		}
	}

	return s[start : start+maxLength]
}
