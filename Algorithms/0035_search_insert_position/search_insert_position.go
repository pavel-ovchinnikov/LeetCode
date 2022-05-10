package search_insert_position_0035

func searchInsert(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	for l <= r {
		curr := (r + l) / 2
		if nums[curr] == target {
			return curr
		}

		if nums[curr] < target {
			l = curr + 1
		} else {
			r = curr - 1
		}
	}
	return l
}
