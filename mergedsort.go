package main

import (
	"sort"
)

// time complexity : O ( (sizeA + sizeB) log (sizeA + sizeB) )
// extra space used : O ( (sizeA + sizeB) )
func execMergedSort(sliceA []int, sliceB []int) []int {
	merged := append(sliceA, sliceB...)
	sort.Ints(merged)
	return merged
}
