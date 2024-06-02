package main

func bulkSend(numMessages int) float64 {
	var totalCost float64 = 0.0
	for i := 0; i < numMessages; i++ {
		totalCost += (1.0 + (float64(i) / 100.0))
	}
	return totalCost
}
