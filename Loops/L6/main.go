package main

func countConnections(groupSize int) int {
	connectionCounter := 0

	for i:= 1; i < groupSize; i++ {
		connectionCounter += 1 * i
	}
	return connectionCounter
}
