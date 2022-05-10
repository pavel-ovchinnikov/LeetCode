package implement_strstr_0028

func strStr(haystack string, needle string) int {
	haystackLen := len(haystack)
	needleLen := len(needle)

	for i := 0; i <= haystackLen-needleLen; i++ {
		if haystack[i:i+needleLen] == needle {
			return i
		}
	}

	return -1
}
