package main

func getMonthlyPrice(tier string) int {
	if tier == "basic" {
		pennies := 10000 
		return pennies
	} else if tier == "premium" {
		return 15000
	} else if tier == "enterprise" {
		return 50000
	} else {
		return 0
	}
}
