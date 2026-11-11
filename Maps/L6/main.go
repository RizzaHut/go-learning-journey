package main

func getNameCounts(names []string) map[rune]map[string]int {
	// make new maps for counting names
	nameCounts := make(map[rune]map[string]int)

	// iterate each name in names slices
	for i := 0; i < len(names); i++ {
		// convert each name into slices of runes (word)
		runes := []rune(names[i])
		// catch the first word of each name
		firstWord := runes[0]

		// check if the first word is already on the list
		_, ok := nameCounts[firstWord]
		// if NOT create a new key 
		if !ok {
			nameCounts[firstWord] = make(map[string]int)
		}

		// add each name and the counter into the inner list
		nameCounts[firstWord][names[i]]++
	} 
	return nameCounts
}
