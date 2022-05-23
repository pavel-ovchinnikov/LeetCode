package longest_palindromic_substring_0005

func longestPalindrome(s string) string {

	begin, end := 0, 0

	for i := 0; i < len(s); i++ {
		c := s[i]
		left, right := i, i

		for left >= 0 && s[left] == c {
			left--
		}

		for right < len(s) && s[right] == c {
			right++
		}

		for left >= 0 && right < len(s) {
			if s[left] != s[right] {
				break
			}
			left--
			right++
		}

		left += 1
		if end-begin < right-left {
			end = right
			begin = left

			if (end - begin) >= len(s)/2 {
				break
			}
		}
	}

	return s[begin:end]
}
