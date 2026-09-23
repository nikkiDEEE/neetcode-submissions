func majorityElement(nums []int) []int {
	cand1, cand2 := 0, 0
	ct1, ct2 := 0, 0

	for _, num := range nums {
		if num == cand1 {
			ct1++
		} else if num == cand2 {
			ct2++
		} else if ct1 == 0 {
			cand1 = num
			ct1 = 1
		} else if ct2 == 0 {
			cand2 = num
			ct2 = 1
		} else {
			ct1--
			ct2--
		}
	}

	var result []int
	count1, count2 := 0, 0

	for _, num := range nums {
		if num == cand1 {
			count1++
		}
		if num == cand2 {
			count2++
		}
	}

	if count1 > len(nums)/3 {
		result = append(result, cand1)
	}

	if count2 > len(nums)/3 {
		result = append(result, cand2)
	}
	
	return result
}
