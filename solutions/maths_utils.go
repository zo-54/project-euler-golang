package solutions

import (
	"slices"
)

const maxLimit uint64 = 1 << 30

var first100Primes = []uint64{
	2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71,
	73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173,
	179, 181, 191, 193, 197, 199, 211, 223, 227, 229, 233, 239, 241, 251, 257, 263, 269, 271, 277, 281,
	283, 293, 307, 311, 313, 317, 331, 337, 347, 349, 353, 359, 367, 373, 379, 383, 389, 397, 401, 409,
	419, 421, 431, 433, 439, 443, 449, 457, 461, 463, 467, 479, 487, 491, 499, 503, 509, 521, 523, 541,
}

var (
	limit      uint64 = 30 << 12
	lastLimit  uint64 = uint64(len(first100Primes))
	primesList        = slices.Clone(first100Primes)
)

// Reset variables for benchmarking purposes
func resetMaths() {
	limit = 30 << 17
	lastLimit = uint64(len(first100Primes))
	primesList = slices.Clone(first100Primes)
}

// Based on the Sieve of Atkin (https://en.wikipedia.org/wiki/Sieve_of_Atkin)
func genMorePrimes() {
	if limit == lastLimit {
		panic("cannot generate more prime numbers")
	}

	period := uint64(60)
	s := []uint64{1, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 49, 53, 59} // set of wheel "hit" positions for a 2/3/5 wheel rolled twice as per the Atkin algorithm
	maybePrimes := make(map[uint64]bool)
	currMaxPrime := primesList[len(primesList)-1]

	// Condition 1
	s1 := []uint64{1, 13, 17, 29, 37, 41, 49, 53}
	for x := uint64(1); 4*x*x+1 <= limit; x++ {
		for y := uint64(1); 4*x*x+y*y <= limit; y += 2 {
			if n := 4*x*x + y*y; slices.Contains(s1, n%60) {
				maybePrimes[n] = !maybePrimes[n]
			}
		}
	}

	// Condition 2
	s2 := []uint64{7, 19, 31, 43}
	for x := uint64(1); 3*x*x+4 <= limit; x += 2 {
		for y := uint64(2); 3*x*x+y*y <= limit; y += 2 {
			if n := 3*x*x + y*y; slices.Contains(s2, n%60) {
				maybePrimes[n] = !maybePrimes[n]
			}
		}
	}

	// Condition 3
	s3 := []uint64{11, 23, 47, 59}
	for x := uint64(2); 3*x*x-(x-1)*(x-1) <= limit; x++ {
		for y := x - 1; y < x && y >= 1 && 3*x*x-y*y <= limit; y -= 2 {
			if n := 3*x*x - y*y; slices.Contains(s3, n%60) {
				maybePrimes[n] = !maybePrimes[n]
			}
		}
	}

	forAllCandidates := func(fn func(n uint64) bool) {
		for w := range limit / period {
			for _, x := range s {
				n := period*w + x

				if fn(n) {
					return
				}
			}
		}
	}

	forAllCandidates(func(n uint64) bool {
		if n < 7 {
			return false
		}

		if n*n > limit {
			return true
		}

		if maybePrimes[n] {
			forAllCandidates(func(f uint64) bool {
				c := n * n * f

				if c > limit {
					return true
				}

				maybePrimes[c] = false

				return false
			})
		}

		return false
	})

	forAllCandidates(func(n uint64) bool {
		if n <= currMaxPrime {
			return false
		}

		isPrime, ok := maybePrimes[n]

		if !ok && isPrime {
			panic("should not happen")
		}

		if isPrime {
			primesList = append(primesList, n)
		}

		return false
	})

	lastLimit = limit
	limit = min(limit<<5, maxLimit)
}

func getNthPrime(n uint64) uint64 {
	if n < 1 {
		panic("value of n for getting nth prime must be at least 1")
	}

	for int(n) > len(primesList) {
		genMorePrimes()
	}

	return primesList[n-1]
}

func isPrime(n uint64) bool {
	if n < 2 {
		panic("not able to determine if values less than 2 are prime")
	}

	if slices.Contains(primesList, n) {
		return true
	}

	currMaxPrime := primesList[len(primesList)-1]

	if n < currMaxPrime {
		return false
	}

	for i := currMaxPrime + 2; i <= isqrt(n); i += 2 {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func getPrimeFactors(n uint64) []uint64 {
	primeFactors := []uint64{}

	for _, p := range primesList {
		for n%p == 0 {
			n = n / p
			primeFactors = append(primeFactors, p)

			if n == 1 {
				return primeFactors
			}
		}
	}

	currMaxPrime := primesList[len(primesList)-1]

	for i := currMaxPrime + 2; i <= n; i += 2 {
		for n%i == 0 {
			n = n / i
			primeFactors = append(primeFactors, i)

			if n == 1 {
				return primeFactors
			}
		}
	}

	panic("did not find all prime factors")
}

// uint64 square root function. Returns largest integer less than or equal to the root of the input number.
func isqrt(n uint64) uint64 {
	var (
		x uint64 = n
		y uint64 = 1
	)

	for x > y {
		x = (x + y) / 2
		y = n / x
	}

	return x
}

func triangleNumber(n int) int {
	return n * (n + 1) / 2
}

// Fast doubling Fibonacci algorithm (based on JavaScript version by Project Nayuki: https://www.nayuki.io/page/fast-fibonacci-algorithms)
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
