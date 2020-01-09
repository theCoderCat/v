package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Which one do you wanna play with?")
	v := ask()
	for v != "1" && v != "2" && v != "3" {
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
	fmt.Println("3. benchmark `merged sort` and `sorted merge`")
	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()

	check(err)
	return string(char)
}

func check(e error) {
	if e != nil {
		// log.Fatalf("cmd.Run() failed with %s\n", e)
		panic(e)
	}
}
