package pascals_triangle_0118

func generate(numRows int) [][]int {
	res := [][]int{}
	res = append(res, []int{1})

	for i := 1; i < numRows; i++ {
		line := append([]int{1}, res[i-1]...)

		for j := 1; j < i; j++ {
			line[j] += line[j+1]
		}

		res = append(res, line)

	}
	return res
}
