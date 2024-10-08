package alg

func Factorial(n int) int {
	if n < 0 {
		return 0
	}

	if n == 1 || n == 0 {
		return 1
	}

	return n * Factorial(n-1)
}

func Fibonacci(n int) int {
	if n <= 0 {
		return 0
	}

	if n == 1 || n == 2 {
		return 1
	}

	return Fibonacci(n-1) + Fibonacci(n-2)
}
