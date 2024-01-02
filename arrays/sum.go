package main

func Sum(numbers []int) int {
	var total int
	for _, n := range numbers {
		total += n
	}
	return total
}

func SumAll(numbersArr ...[]int) []int {
	var res []int
	for _, arr := range numbersArr {
		res = append(res, Sum(arr))
	}
	return res
}

func SumAllTails(numbersArr ...[]int) []int {
	var res []int
	for _, arr := range numbersArr {
		if len(arr) == 0 {
			res = append(res, 0)
		} else {
			res = append(res, Sum(arr[1:]))
		}
	}
	return res
}
