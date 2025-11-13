package main

import (
	"fmt"
	"os"
	"unicode"
)

func countLetters(text string) int {
	count := 0

	for _, ch := range text {
		if unicode.IsLetter(ch) || unicode.IsDigit(ch) {
			count++
		}
	}
	return count
}

func countWords(text string) int {
	wordCount := 0
	inWord := false

	for i := 0; i < len(text); i++ {
		ch := text[i]
		
		if ch != ' ' && ch != '\n' && ch != '\t' {
			if !inWord {
				wordCount++
				inWord = true
			}
		} else {
			inWord = false
		}
	}

	return wordCount
}

func countSentences(text string) int {
	count := 0

	for _, ch := range text {
		if ch == '.' || ch == '!' || ch == '?' {
			count++
		}
	}
	return count
}


func main() {
	data, err := os.ReadFile("text.txt")

	if err != nil {
		panic(err)
	}

	text := string(data)

	fmt.Println("Letters: ", countLetters(text))
	fmt.Println("Words: ", countWords(text))
	fmt.Println("Sentences: ", countSentences(text))
}






