func majorityElement(nums []int) int {
    nmap := map[int]int{}
	for i := 0; i < len(nums); i++ {
		nmap[nums[i]] += 1
	}
	highest_count, highest_num := 0, 0
	for k,v := range nmap {
		if v > highest_count {
			highest_num = k
			highest_count = v
		}
	}
	return highest_num
}
