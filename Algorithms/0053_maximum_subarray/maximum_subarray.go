package maximum_subarray_0053

func maxSubArray(nums []int) int {
	maxSum := -1 << 31
	tempSum := -1 << 31
	for _, num := range nums {
		tempSum = max(tempSum+num, num)
		maxSum = max(maxSum, tempSum)
	}
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
