package palindrome_number_0009

func isPalindrome(x int) bool {
	if 0 > x {
		return false
	}
	value := x
	res := 0
	for 0 < value {
		res = res*10 + value%10
		value /= 10
	}
	return res == x
}
