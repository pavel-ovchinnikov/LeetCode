package length_of_last_word_0053

func lengthOfLastWord(s string) int {
	r := len(s) - 1

	for 0 <= r && s[r] == ' ' {
		r--
	}

	l := r - 1

	for 0 <= l && s[l] != ' ' {
		l--
	}

	if l < 0 {
		return r + 1
	}
	return r - l
}
