package arrayslices

//Sum numbers passed
func SumArray(numbers [5]int) int {

	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func Sum(numbers []int) int {

	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}
