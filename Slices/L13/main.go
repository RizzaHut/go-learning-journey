package main

import "strings"

type sms struct {
	id      string
	content string
	tags    []string
}

func tagMessages(messages []sms, tagger func(sms) []string) []sms {
	for i := 0; i < len(messages); i++ {
		tagResult := tagger(messages[i])

		messages[i].tags = tagResult
	}
	return messages
}

func tagger(msg sms) []string {
	tags := []string{}
	lowercaseMessage := strings.ToLower(msg.content)
	if strings.Contains(lowercaseMessage, "urgent") {
		tags = append(tags, "Urgent")
	}
	if strings.Contains(lowercaseMessage, "sale") {
		tags = append(tags, "Promo")
	}
	return tags
}
