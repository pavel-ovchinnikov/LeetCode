package longest_common_prefix_0014

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	minStr := strs[0]
	minLen := len(minStr)
	for _, str := range strs {
		lenStr := len(str)
		if lenStr < minLen {
			minLen = lenStr
			minStr = str
		}
	}

	for i, r := range minStr {
		for j := 0; j < len(strs); j++ {
			if byte(r) != strs[j][i] {
				return strs[j][:i]
			}
		}
	}

	return ""
}
