package main

type cost struct {
	day   int
	value float64
}

func getDayCosts(costs []cost, day int) []float64 {
	costResult := []float64{}

	for _, cost := range costs {
		if cost.day == day {
		costResult = append(costResult, cost.value)
		}
	}
	return costResult
}
