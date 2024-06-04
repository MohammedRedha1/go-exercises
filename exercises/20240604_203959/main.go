package main

func getMessageCosts(messages []string) []float64 {
	costs := make([]float64, len(messages))
	for i := 0; i < len(costs); i++ {
		costs[i] = float64(len(messages[i])) * 0.01
	}
	return costs
}
