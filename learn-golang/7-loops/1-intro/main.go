package main

func bulkSend(numMessages int) float64 {
	cost := 0.0
	for i := range numMessages {
		cost += 1.0 + float64(i)*.01
	}
	return cost
}
