package main

import "strings"

func countDistinctWords(messages []string) int {
	// create map of distinct words
	distinctWords := make(map[string]struct{})

	// iterate each sentence in messages slices
	for i := 0; i < len(messages); i++ {
		words := strings.Fields(messages[i])
		// iterate each word in sentence
		for j := 0; j < len(words); j++ {
			// converd word into lowercase
			wordLowerCase := strings.ToLower(words[j])

			distinctWords[wordLowerCase] = struct{}{}
		}
	}
	return len(distinctWords)
}
