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

//Using variadic functions to pass slice list of any number
func SumAll(numbersToSum ...[]int) []int {
	lengthOfNumberSlicesPassed := len(numbersToSum)
	sums := make([]int, lengthOfNumberSlicesPassed)
	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}
	return sums
}
