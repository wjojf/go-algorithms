package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(LongestWord("Hello, my dear friend. This is a stoofvlees"))
}

func LongestWord(s string) string {

	var longestWord []string
	var currentWord []string

	for i := 0; i < len(s); i++ {
		letter := string(s[i])

		if letter == " " || letter == "," || letter == "." {
			if len(currentWord) > len(longestWord) {
				longestWord = currentWord[0:]
			}
			currentWord = []string{}
			continue
		}

		currentWord = append(currentWord, letter)
	}

	if len(currentWord) > len(longestWord) {
		longestWord = currentWord[0:]
	}

	return strings.Join(longestWord, "")
}
