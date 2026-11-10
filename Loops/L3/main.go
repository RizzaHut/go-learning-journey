package main

func getMaxMessagesToSend(costMultiplier float64, budgetInPennies int) int {
	actualCostInPennies := 1.0
	maxMessagesToSend := 1
	balance := float64(budgetInPennies) - actualCostInPennies
	for balance > actualCostInPennies{
		actualCostInPennies *= costMultiplier
		balance -= actualCostInPennies
		maxMessagesToSend++
	}
	return maxMessagesToSend
}
