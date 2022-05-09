package remove_duplicates_from_sorted_array_0026

func removeDuplicates(nums []int) int {

	if len(nums) == 1 {
		return 1
	}
	i := 1
	k := 0
	for i < len(nums) {
		if nums[i] != nums[k] {
			k++
			nums[k] = nums[i]
		}
		i++
	}

	return k + 1
}
