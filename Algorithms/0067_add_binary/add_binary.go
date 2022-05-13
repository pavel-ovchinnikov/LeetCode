package add_binary_0067

func addBinary(a string, b string) string {
	aLen := len(a)
	bLen := len(b)
	maxLen := max(aLen, bLen)
	bytes := make([]byte, maxLen)

	temp := 0
	for i := 0; i < maxLen; i++ {
		aValue := 0
		bValue := 0

		if aLen-i-1 < 0 {
			aValue = 0
		} else {
			aValue = int(a[aLen-i-1] - '0')
		}
		if bLen-i-1 < 0 {
			bValue = 0
		} else {
			bValue = int(b[bLen-i-1] - '0')
		}

		temp += aValue + bValue
		bytes[maxLen-i-1] = byte(temp%2) + '0'
		temp = temp / 2
	}

	if temp != 0 {
		return string(append([]byte{'1'}, bytes...))
	}
	return string(bytes)
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
