package main

func Sum(numbers [5]int) int {
	var total int
	for _, n := range numbers {
		total += n
	}
	return total
}
