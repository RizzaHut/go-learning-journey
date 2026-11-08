package main

import "fmt"

func reformat(message string, formatter func(string) string) string {
	firstMessage := formatter(message)
	secondMessage := formatter(firstMessage)
	thirdMessage := formatter(secondMessage)
	return fmt.Sprintf("TEXTIO: %v", thirdMessage)
}
