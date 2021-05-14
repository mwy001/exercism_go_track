package diffsquares

func SquareOfSum(input int) int {
	
	var sum int

	for i:=1; i <= input; i++ {
		sum += i
	}

	return sum * sum
}

func SumOfSquares(input int) int {
	var res int

	for i:=1; i <= input; i++ {
		res += i * i
	}

	return res
}

func Difference(input int) int {
	return SquareOfSum(input) - SumOfSquares(input)
}
