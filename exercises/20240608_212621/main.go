package main

func getNameCounts(names []string) map[rune]map[string]int {
	counts := make(map[rune]map[string]int)
	for _, name := range names {
		runes := []rune(name) // just to handle +1 byte chars
		firstCahr := runes[0]
		if counts[firstCahr] == nil {
			counts[firstCahr] = make(map[string]int)
		}
		counts[firstCahr][name]++
	}
	return counts
}
