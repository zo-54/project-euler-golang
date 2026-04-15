package solutions

import (
	"slices"
)

var first100Primes = []uint64{
	2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71,
	73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173,
	179, 181, 191, 193, 197, 199, 211, 223, 227, 229, 233, 239, 241, 251, 257, 263, 269, 271, 277, 281,
	283, 293, 307, 311, 313, 317, 331, 337, 347, 349, 353, 359, 367, 373, 379, 383, 389, 397, 401, 409,
	419, 421, 431, 433, 439, 443, 449, 457, 461, 463, 467, 479, 487, 491, 499, 503, 509, 521, 523, 541,
}

func isPrime(n uint64) bool {
	if n < 2 {
		panic("Not able to determine if values less than 2 are prime")
	}

	if slices.Contains(first100Primes, n) {
		return true
	}

	if n < first100Primes[99] {
		return false
	}

	for i := first100Primes[99] - first100Primes[99]%6; i <= sqrt(n)+1; i += 6 {
		for _, j := range []uint64{i - 1, i + 1} {
			if n%j == 0 {
				return false
			}
		}
	}

	return true
}

func getPrimeFactors(n uint64) []uint64 {
	primeFactors := []uint64{}

	for _, p := range first100Primes {
		for n%p == 0 {
			n = n / p
			primeFactors = append(primeFactors, p)

			if isPrime(n) {
				primeFactors = append(primeFactors, n)
				return primeFactors
			}
		}
	}

	for i := first100Primes[99] + 2; i < n; i += 2 {
		for n%i == 0 {
			n = n / i
			primeFactors = append(primeFactors, i)

			if isPrime(n) {
				primeFactors = append(primeFactors, n)
				return primeFactors
			}
		}
	}

	return primeFactors
}

func sqrt(n uint64) uint64 {
	if n < 4 {
		return 1
	}

	var maybeRoot uint64 = 2

	for (maybeRoot+1)*(maybeRoot+1) < n {
		nOverR := n / maybeRoot
		nextNOverR := (nOverR + maybeRoot) / 2
		maybeRoot = n/nextNOverR + 1
	}

	return maybeRoot
}

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
