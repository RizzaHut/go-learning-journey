package main

func sum(nums ...int) int {
	sumNumbers := 0
	for i := 0; i < len(nums); i++ {
		sumNumbers += nums[i]
	}
	return sumNumbers
}
