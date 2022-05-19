package valid_palindrome_0125

import "strings"

func isChar(c byte) bool {
	return ('a' <= c && c <= 'z') || ('0' <= c && c <= '9')
}

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	left := 0
	right := len(s) - 1
	for left < right {

		for left < right && !isChar(s[left]) {
			left++
		}

		for left < right && !isChar(s[right]) {
			right--
		}

		if s[left] != s[right] {
			return false
		}

		left++
		right--
	}

	return true
}
