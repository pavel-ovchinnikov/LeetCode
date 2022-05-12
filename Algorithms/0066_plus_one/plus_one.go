package plus_one_0066

func plusOne(digits []int) []int {
	length := len(digits)

	digits[length-1]++
	for i := length - 1; 0 < i; i-- {
		if digits[i] < 10 {
			break
		}
		digits[i] -= 10
		digits[i-1]++
	}

	if digits[0] >= 10 {
		digits[0] -= 10
		digits = append([]int{1}, digits...)
	}

	return digits
}
