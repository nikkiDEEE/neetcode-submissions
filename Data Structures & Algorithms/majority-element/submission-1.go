func majorityElement(nums []int) int {
    candidate, count := 0, 0
	for i := 0; i < len(nums); i++ {
		if count == 0 {
			candidate = nums[i]
		}

		if candidate == nums[i] {
			count++
		} else {
			count--
		}
	}
	return candidate
}