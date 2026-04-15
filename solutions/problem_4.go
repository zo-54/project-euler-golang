package solutions

import (
	"fmt"
	"slices"
)

func problem4() string {
	toPalindrome := func(n uint64, combineMiddleDigit bool) uint64 {
		palindrome := n

		if combineMiddleDigit {
			palindrome = palindrome / 10
		}

		for n > 0 {
			palindrome = palindrome*10 + n%10
			n = n / 10
		}

		return palindrome
	}

	// Maximum product of 2 3-digit numbers is 999 * 999 = 998001
	// Largest palindrome number less than this is 997799, therefore start at 997
	var startNumber uint64 = (997)

	// Minimum product of 2 3-digit numbers is 100 * 100 = 10000
	// Smallest palindrome number greater than this is 10001, therefore stop at 100
	var stopNumber uint64 = 100

	check := func(n uint64) bool {
		primeFactors := getPrimeFactors(n)

		// 4-digit numbers are not allowed
		if slices.ContainsFunc(primeFactors, func(f uint64) bool {
			return f > 999
		}) {
			return false
		}

		// Only 2 prime factors which are both 3-digit numbers
		if len(primeFactors) == 2 && primeFactors[0] >= 100 && primeFactors[1] >= 100 {
			return true
		} else if len(primeFactors) == 2 {
			return false
		}

		soln := combineFactors(primeFactors)

		return len(soln) == 2
	}

	// Check 6-digit palindromes first
	for halfP := startNumber; halfP >= stopNumber; halfP-- {
		p := toPalindrome(halfP, false)

		if check(p) {
			return fmt.Sprintf("%d", p)
		}
	}

	// Then check 5-digit palindromes
	for n := uint64(997); n >= 100; n-- {
		p := toPalindrome(n, true)

		if check(p) {
			return fmt.Sprintf("%d", p)
		}
	}

	panic("No solution found")
}

func combineFactors(factors []uint64) []uint64 {
	l := len(factors)

	for i := 0; i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			product := factors[i] * factors[j]

			if product > 999 {
				continue
			}

			newFactors := []uint64{product}

			for k := range l {
				if k == i || k == j {
					continue
				}

				newFactors = append(newFactors, factors[k])
			}

			if len(newFactors) == 2 {
				return newFactors
			}

			newCombinedFactors := combineFactors(newFactors)

			if len(newCombinedFactors) == 2 {
				return newCombinedFactors
			}
		}
	}

	return []uint64{}
}
