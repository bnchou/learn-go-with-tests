package main

func Sum(numbers []int) int {
	var total int
	for _, n := range numbers {
		total += n
	}
	return total
}
