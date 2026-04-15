package solutions

import (
	"fmt"
)

func problem5() string {
	const (
		min uint64 = 2 // Ignore 1 as it is a trivial case
		max uint64 = 20
	)

	allFactorsMap := make(map[uint64]uint64)

	for n := max; n >= min; n-- {
		if c, ok := allFactorsMap[n]; ok && c > 0 {
			continue
		}

		if n*2 <= max || n*3 <= max || n*5 <= max {
			continue
		}

		factorMap := make(map[uint64]uint64)

		primeFactors := getPrimeFactors(n)

		for _, f := range primeFactors {
			if _, ok := factorMap[f]; !ok {
				factorMap[f] = 0
			}

			factorMap[f]++
		}

		for f, c := range factorMap {
			if _, ok := allFactorsMap[f]; !ok {
				allFactorsMap[f] = 0
			}

			if c > allFactorsMap[f] {
				allFactorsMap[f] = c
			}
		}
	}

	var v uint64 = 1

	for f, c := range allFactorsMap {
		for range c {
			v *= f
		}
	}

	return fmt.Sprintf("%d", v)
}
