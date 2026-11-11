package main

func indexOfFirstBadWord(msg []string, badWords []string) int {
	badWordsFound := false
	for m, _ := range msg{
		for b,_ := range badWords {
			if msg[m] == badWords[b] {
				badWordsFound = true
				return m
			}
		}	
	}
	if badWordsFound == false {
		return -1
	}
	return 0
}
