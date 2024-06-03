package main

func getMessageWithRetries(primary, secondary, tertiary string) ([3]string, [3]int) {
	messages := [3]string{primary, secondary, tertiary}
	var costs [3]int
	for i := 0; i < 3; i++ {
		if i == 0 {
			costs[i] = len(messages[i])
			continue
		}
		costs[i] = len(messages[i]) + costs[i-1]
	}
	return messages, costs
}
