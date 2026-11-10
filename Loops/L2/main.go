package main

func maxMessages(thresh int) int {
	messageCounter := 0
	totalBudget := thresh
	for i := 0; i < totalBudget; i++ {
		currentCost := 100 + (1 *i)
		if totalBudget < currentCost {
			break
		}
		totalBudget -= currentCost
		messageCounter += 1
	}
	return messageCounter
}
