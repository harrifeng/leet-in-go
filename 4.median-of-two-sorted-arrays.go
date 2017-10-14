package leetcode

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	cnt := len(nums1) + len(nums2)
	if cnt%2 == 1 {
		return float64(findNth(nums1, nums2, cnt/2+1))
	} else {
		return float64(findNth(nums1, nums2, cnt/2+1)+findNth(nums1, nums2, cnt/2)) / 2.0
	}
}

func findNth(nums1 []int, nums2 []int, nth int) int {
	if len(nums1) > len(nums2) {
		return findNth(nums2, nums1, nth)
	}
	if len(nums1) == 0 {
		return nums2[nth-1]
	}
	if len(nums2) == 0 {
		return nums1[nth-1]
	}
	if nth == 1 {
		return minInt(nums1[0], nums2[0])
	}

	partA := minInt(nth/2, len(nums1))
	partB := nth - partA

	if nums1[partA-1] < nums2[partB-1] {
		return findNth(nums1[partA:], nums2, nth-partA)
	} else {
		return findNth(nums1, nums2[partB:], nth-partB)
	}
}

func minInt(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
