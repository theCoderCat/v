package main

import (
	"fmt"
	"sort"
)

func merge() {
	listA := []int{1, 3, 9}
	listB := []int{2, 9, 4}
	fmt.Println("list A: ", listA)
	fmt.Println("list B: ", listB)
	merged := execSortedMerge(listA, listB)
	fmt.Println("merged and sorted result: ", merged)
	fmt.Println("total time complexity: O(sAlogsA + sBlogsB + sA + sB)")
	fmt.Println("total space complexity: O(sA + sB)")
}

func execSortedMerge(listA []int, listB []int) []int {
	sA := len(listA)
	sB := len(listB)

	// sort 2 lists first
	// time complexity: O(sA log sA + sB log sB)
	sort.Ints(listA)
	sort.Ints(listB)

	// then merge
	// create new slice to store sorted list
	merged := make([]int, 0, sA+sB)

	iA, iB := 0, 0

	// three for loops have O(sA + sB) time complexity in total
	for iA < sA && iB < sB {
		if listA[iA] <= listB[iB] {
			merged = append(merged, listA[iA])
			//fmt.Printf("put %d from list A to merged list \n", listA[iA])
			//fmt.Println("merged list became: ", merged)
			iA++
		} else {
			merged = append(merged, listB[iB])
			//fmt.Printf("put %d from list B to merged list \n", listB[iB])
			//fmt.Println("merged list became: ", merged)
			iB++
		}
	}

	// then put the rest of list A to the merged list (if any)
	for iA < sA {
		merged = append(merged, listA[iA])
		//fmt.Printf("put %d from list A to merged list \n", listA[iA])
		//fmt.Println("merged list became: ", merged)
		iA++
	}

	// then put the rest of list B to the merged list (if any)
	for iB < sB {
		merged = append(merged, listB[iB])
		//fmt.Printf("put %d from list B to merged list \n", listB[iB])
		//fmt.Println("merged list became: ", merged)
		iB++
	}

	return merged
}
