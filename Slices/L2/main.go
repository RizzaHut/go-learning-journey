package main

import (
	"errors"
)

const (
	planFree = "free"
	planPro  = "pro"
)

func getMessageWithRetriesForPlan(plan string, messages [3]string) ([]string, error) {
	if plan == planFree {
		freeMessages := messages[:2]
		return freeMessages, nil
	} 
	if plan == planPro {
		proMessages := messages[:]
		return proMessages, nil
	} 
	return nil, errors.New("unsupported plan")
}
