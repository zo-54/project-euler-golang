package solutions

func triangleNumber(n int) int {
	return n * (n + 1) / 2
}

/*
 * Fast doubling Fibonacci algorithm
 * Based on JavaScript version by Project Nayuki
 * https://www.nayuki.io/page/fast-fibonacci-algorithms
 */
func fibonacci(n uint64) uint64 {
	v, _ := fib(n)

	return v
}

func fib(n uint64) (uint64, uint64) {
	if n == 0 {
		return 0, 1
	}

	a, b := fib(n / 2)
	c := a * (b*2 - a)
	d := a*a + b*b

	if n%2 == 0 {
		return c, d
	} else {
		return d, c + d
	}
}
