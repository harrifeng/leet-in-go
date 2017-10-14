package leetcode

import "strings"

func longestPalindrome(s string) string {
	ns := "^#" + strings.Join(strings.Split(s, ""), "#") + "#$"

	tmp_centr_pos := 0
	tmp_right_pos := 0

	maxv := 0
	maxv_pos := 0

	sizes := [3000]int{0}

	for i := 1; i < len(ns)-1; i++ {
		size := 0

		if i < tmp_right_pos {
			size = minInt2(tmp_right_pos-i, sizes[2*tmp_centr_pos-i])
		}

		for ns[i-size-1] == ns[i+size+1] {
			size++
		}

		sizes[i] = size

		if sizes[i]+size > tmp_right_pos {
			tmp_right_pos = sizes[i] + size
			tmp_centr_pos = i
		}

		if sizes[i] > maxv {
			maxv = sizes[i]
			maxv_pos = i
		}
	}

	return s[(maxv_pos-maxv)/2 : (maxv_pos+maxv)/2]
}

func minInt2(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
