package single_number_0136

func singleNumber(nums []int) int {
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		res = res ^ nums[i]
	}
	return res
}
