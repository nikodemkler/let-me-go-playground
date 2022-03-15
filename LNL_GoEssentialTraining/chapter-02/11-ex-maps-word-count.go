package main

import (
	"fmt"
	"strings"
)

func main () {
	sampleText := `
		Needles and pins
		Needles and pins
		See me a sail
		To catch me the wind
	`

	loweredText := strings.ToLower(sampleText)
	parts := strings.Fields(loweredText)

	fmt.Println(parts)
	fmt.Println(len(parts))

	countsMap := map[string]int {}
	for _, word := range parts {
		countsMap[word] += 1 // or countsMap[word]++
	}

	fmt.Println(countsMap)
}

