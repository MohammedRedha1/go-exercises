package main

type cost struct {
	day   int
	value float64
}

func getMaxDay(costs []cost) int {
	max := 0
	for i := 0; i < len(costs); i++ {
		if costs[i].day > max {
			max = costs[i].day
		}
	}
	return max
}
func getCostsByDay(costs []cost) []float64 {
	costsByDay := make([]float64, getMaxDay(costs)+1)
	for i := 0; i < len(costs); i++ {
		costsByDay[costs[i].day] += costs[i].value
	}

	return costsByDay
}
