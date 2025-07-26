package main

func maxMessages(thresh int) int {
	cost := 0
	messages := -1
	for i := 0; cost <= thresh; i++ {
		cost += 100 + i
		messages++
	}
	return messages
}
