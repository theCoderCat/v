package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	fmt.Println("Which one do you wanna play with?")
	v := ask()
	if v != "1" && v != "2" && v != "3" {
		fmt.Println("invalid option, please pick again")
		v = ask()
	}
	switch v {
	case "1":
		execPalindrome()
		break
	case "2":
		merge()
		break
	case "3":
		startBenchmarking()
		break
	}
}

func ask() string {
	fmt.Println("1. palindrome validator")
	fmt.Println("2. a look at list merging solution")
	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()

	check(err)
	return string(char)
}

func generateBigSlice(size int) []int {
	randomSlice := make([]int, 0, size)
	for i := 0; i < size; i++ {
		randomSlice = append(randomSlice, rand.Int())
	}
	return randomSlice
}

func startBenchmarking() {
	// create new file to store benchmark result
	now := time.Now()
	filename := "./benchmark results/benchmark_" + now.Format("20060102150405") + ".txt"

	f, err := os.Create(filename)
	check(err)
	defer f.Close()
	for i := 1; i < 11; i++ {
		fmt.Println("test ", i)
		f.WriteString(fmt.Sprintf("test %d\n", i))
		result := benchmark()
		f.WriteString(result)
	}
	f.Sync()
}

func benchmark() string {
	result := ""
	sliceASize := rand.Intn(90000000) + 10000000
	sliceBSize := rand.Intn(90000000) + 10000000
	fmt.Println("generating arrays...")
	start := time.Now()
	sliceA1 := generateBigSlice(sliceASize)
	sliceB1 := generateBigSlice(sliceBSize)

	sliceA2 := generateBigSlice(sliceASize)
	sliceB2 := generateBigSlice(sliceBSize)
	t := time.Now()
	elapsed := t.Sub(start)

	fmt.Println("generated 2 arrays with", sliceASize, "and 2 array with", sliceBSize, "items in", elapsed)
	result += fmt.Sprintf("generated 2 arrays with %d and %d items in %v\n", sliceASize, sliceBSize, elapsed)

	start = time.Now()
	execMergedSort(sliceA2, sliceB2)
	t = time.Now()
	elapsed = t.Sub(start)
	fmt.Println("merged sort finished in", elapsed)
	result += fmt.Sprintf("- merged then sort finished in:\t\t%v\n", elapsed)

	start = time.Now()
	execSortedMerge(sliceA1, sliceB1)
	t = time.Now()
	elapsed = t.Sub(start)
	fmt.Println("sorted merge finished in", elapsed)
	result += fmt.Sprintf("- sorted then merge finished in:\t%v\n\n", elapsed)

	fmt.Print("\n\n")
	return result
}

func check(e error) {
	if e != nil {
		// log.Fatalf("cmd.Run() failed with %s\n", e)
		panic(e)
	}
}
