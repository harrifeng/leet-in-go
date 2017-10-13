package leetcode

func lengthOfLongestSubstring(s string) int {
	maxv := 0
	start := 0

	m := make(map[rune]int)

	for i, c := range s {
		if v, ok := m[c]; ok == true {
			start = maxInt(start, v+1)
		}
		maxv = maxInt(maxv, i-start+1)
		m[c] = i
	}

	return maxv
}

func maxInt(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
