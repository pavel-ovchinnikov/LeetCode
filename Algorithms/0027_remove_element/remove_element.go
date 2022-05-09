package remove_element_0027

func removeElement(nums []int, val int) int {
	i := 0
	k := -1
	for i < len(nums) {
		if nums[i] != val {
			k++
			nums[k] = nums[i]
		}
		i++
	}
	return k + 1
}
