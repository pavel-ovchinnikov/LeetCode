package pascals_triangle_0119

func getRow(rowIndex int) []int {
	size := rowIndex + 1
	row := make([]int, size)
	for i := 0; i < size; i++ {
		line := row[size-i-1 : size]
		for j := 1; j < i; j++ {
			line[j] += line[j+1]
		}
		line[0] = 1
	}
	return row
}
