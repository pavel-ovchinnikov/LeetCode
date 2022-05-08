package roman_to_integer_0013

var mapper = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	res := 0

	last := 0
	for i := len(s) - 1; i >= 0; i-- {
		temp := mapper[s[i]]

		if temp < last {
			res -= temp
		} else {
			res += temp
		}

		last = temp
	}
	return res
}
