package arraysandslices

// return the sum of a list of int32
func Sum(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, slice := range numbersToSum {
		sums = append(sums, Sum(slice))
	}
	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, slice := range numbersToSum {
		if len(slice) == 0 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, Sum(slice[1:]))
		}
	}
	return sums
}
