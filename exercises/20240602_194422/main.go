package main

func maxMessages(thresh int) int {
	messagesCount := 0
	for i := 0; ; i++ {
		thresh -= (100 + i)
		if thresh < 0 {
			break
		}
		messagesCount++
	}
	return messagesCount
}
