package main

import (
	"strings"
)

func countDistinctWords(messages []string) int {
	occured := make(map[string]bool)
	for _, msg := range messages {
		words := strings.Fields(msg)
		for _, word := range words {
			if !occured[strings.ToLower(word)] {
				occured[strings.ToLower(word)] = true
			}

		}
	}
	return len(occured)
}
