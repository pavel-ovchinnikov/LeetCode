package longest_substring_without_repeating_characters_0003

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lengthOfLongestSubstring(s string) int {
	maxLength := 0
	left := 0
	buffer := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		index, exist := buffer[byte(s[i])]
		if exist && index >= left {
			left = index + 1
		}

		buffer[byte(s[i])] = i
		maxLength = max(i-left+1, maxLength)
	}
	return maxLength
}
