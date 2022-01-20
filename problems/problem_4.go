package problems

import (
	"math"
	"sort"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var m = len(nums1)
	var n = len(nums2)

	var mergedArr []int

	if m == 0 && n == 0 {
		return 0.0
	}

	if m == 0 {
		mergedArr = append(mergedArr, nums2...)
	} else if n == 0 {
		mergedArr = append(mergedArr, nums1...)
	} else {
		mergedArr = append(nums1, nums2...)
	}

	// sort the array
	sort.Ints(mergedArr)

	var length = len(mergedArr)
	var median = 0.0

	// if even
	if length%2 == 0 && length >= 2 {
		upperIndex := int(math.Ceil(float64(length) / 2))
		lowerIndex := upperIndex - 1

		median = float64(mergedArr[lowerIndex]+mergedArr[upperIndex]) / 2.0

	} else if length%2 == 1 {
		medianIndex := int(math.Ceil(float64(length)/2)) - 1

		median = float64(mergedArr[medianIndex])
	}

	return median
}

func Problem4() {
	//	println(findMedianSortedArrays([]int{1, 2}, []int{2}))
	println(findMedianSortedArrays([]int{1, 3}, []int{2}))
}
