package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"unicode/utf8"
)

// take a string,
// keep only alphanumeric character, strip all other characters
// then make them all lowercase
func normalize(str string) string {
	// make a Regex to say we only want letters and numbers
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	check(err)

	processedString := reg.ReplaceAllString(str, "")

	lowerCaseStr := strings.ToLower(processedString)

	return lowerCaseStr
}

// take a string and compare its head and tail
func test(str string) bool {
	runes := []rune(str)
	match := true
	count := utf8.RuneCountInString(str) - 1
	// time complexity: O(count/2)
	// space complexity: O(count)
	//this loop will stop as soon as match turn to false
	for i, j := 0, count; i < j && match; i, j = i+1, j-1 {
		fmt.Printf("Comparing %s with %s \n", string(runes[i]), string(runes[j]))
		match = runes[i] == runes[j]
	}
	return match
}

func execPalindrome() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("may i have some text please? \n")
	text, _ := reader.ReadString('\n')

	// make it lowercase
	processedString := normalize(text)
	fmt.Println("original string:\t", text)
	fmt.Println("string :\t", processedString)

	equal := test(processedString)

	if equal {
		fmt.Println("this is indeed a palindrome!")
	} else {
		fmt.Println("no! it's not a palindrome")
	}
}

// helper function to check error
func check(e error) {
	if e != nil {
		// log.Fatalf("cmd.Run() failed with %s\n", e)
		panic(e)
	}
}
