package leetcode

import "bytes"

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	group := 2*numRows - 2
	ret := []bytes.Buffer{}

	for i := 0; i < numRows; i++ {
		ret = append(ret, bytes.Buffer{})
	}

	for i := 0; i < len(s); i++ {
		curr := i % group
		if curr < numRows {
			ret[curr].WriteByte(s[i])
		} else {
			ret[group-curr].WriteByte(s[i])
		}
	}
	rr := ""
	for i := 0; i < numRows; i++ {
		rr += ret[i].String()
	}
	return rr
}
