package leetcode

import "math"

func reverse(x int) int {
	neg := false
	if x < 0 {
		neg = true
		x = -1 * x
	}

	ret := 0

	for x != 0 {
		ret = ret*10 + x%10
		if ret > int(math.Pow(2, 31))-1 {
			return 0
		}

		x /= 10
	}

	if neg {
		return -1 * ret
	} else {
		return ret
	}
}
