package main

func getMessageCosts(messages []string) []float64 {
	messageCosts := make([]float64, len(messages))
	for i := 1; i < len(messages); i++ {
		messageCosts = append(messageCosts, float64(len(messages[i])) * 0.01)
	}
	return messageCosts
}
