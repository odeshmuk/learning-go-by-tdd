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

	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

//Sum of tails ie leaving the 1st element in the slice out
func SumTails(numberTailsToSum ...[]int) int {
	sumOfTails := 0
	for _, numbers := range numberTailsToSum {
		if len(numbers) != 0 {
			sumOfTails += Sum(numbers[1:])
		}
	}
	return sumOfTails
}
