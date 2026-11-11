package main

func updateCounts(messagedUsers []string, validUsers map[string]int) {
	for i := 0; i < len(messagedUsers); i++ {
		user := messagedUsers[i]
		_, ok := validUsers[user]
		if ok == true {
			validUsers[user]++
		}
	}
}
