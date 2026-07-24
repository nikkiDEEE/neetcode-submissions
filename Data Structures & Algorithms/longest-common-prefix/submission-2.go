func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	shortest := strs[0]
	for _, str := range strs {
		if len(str) < len(shortest) {
			shortest = str
		}
	}

	for _, str := range strs {
		var i int
		for i = 0; i < len(str) && i < len(shortest); i++ {
			if shortest[i] != str[i] {
				break
			}
		}
		shortest = shortest[:i]
	}

	return shortest
}
