package merge_sorted_array_0088

func merge(nums1 []int, m int, nums2 []int, n int) {

	i := m - 1
	j := n - 1

	for k := len(nums1) - 1; k >= 0; k-- {
		if j < 0 {
			nums1[k] = nums1[i]
			i--
			continue
		}
		if i < 0 {
			nums1[k] = nums2[j]
			j--
			continue
		}

		if nums1[i] <= nums2[j] {
			nums1[k] = nums2[j]
			j--
		} else {
			nums1[k] = nums1[i]
			i--
		}
	}
}
