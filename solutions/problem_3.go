package solutions

import (
	"fmt"
	"slices"
)

func problem3() string {
	const n = 600851475143

	factors := getPrimeFactors(n)

	largestPrimeFactor := slices.Max(factors)

	return fmt.Sprintf("%d", largestPrimeFactor)
}
