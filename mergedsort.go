package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func generateBigSlice(size int) []int {
	randomSlice := make([]int, 0, size)
	for i := 0; i < size; i++ {
		randomSlice = append(randomSlice, rand.Int())
	}
	return randomSlice
}

func startBenchmarking() {
	for i:=1; i < 11; i++ {
		fmt.Println("test ", i)
		benchmark()
	}
}

func benchmark() {
	sliceASize := rand.Intn(90000000) + 10000000
	sliceBSize := rand.Intn(90000000) + 10000000
	fmt.Println("slices with ", sliceASize, " and ", sliceBSize, " items")
	sliceA := generateBigSlice(sliceASize)
	sliceB := generateBigSlice(sliceBSize)

	start := time.Now()
	execMergedSort(sliceA, sliceB)
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println("merged sort took ", elapsed)

	start = time.Now()
	execSortedMerge(sliceA, sliceB)
	t = time.Now()
	elapsed = t.Sub(start)
	fmt.Println("sorted merge took ", elapsed)
	fmt.Print("\n\n")
}

func execMergedSort(sliceA []int, sliceB []int) []int {
	merged := append(sliceA, sliceB...)
	sort.Ints(merged)
	return merged
}
