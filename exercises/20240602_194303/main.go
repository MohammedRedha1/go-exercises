package main

func maxMessages(thresh int) int {
	messagesCount := 0
	for i := 0; thresh > 0; i++ {
		thresh -= (100 + i)
		messagesCount++
	}
	return messagesCount
}
