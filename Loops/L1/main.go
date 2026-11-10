package main

func bulkSend(numMessages int) float64 {
	var totalCost float64 = 0.0

	for m := 0; m < numMessages; m++ {
		currentCost := 1.0 + (float64(m) * 0.01)
		totalCost += currentCost
	}
	return float64(totalCost)
	
}
