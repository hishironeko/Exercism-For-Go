package differenceofsquares

func SquareOfSum(n int) int {
	sum := n * (n + 1) / 2
    return sum * sum
}

func SumOfSquares(n int) int {
	sumSq := n * (n + 1) * (2 * n + 1) / 6
    return sumSq
}

func Difference(n int) int {
	sqSum := SquareOfSum(n)
    sumSq := SumOfSquares(n)
    return sqSum - sumSq
}
