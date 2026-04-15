package solutions

import "fmt"

func problem1() string {
	const (
		max     = 1000
		factorA = 3
		factorB = 5
	)

	cycle := factorA * factorB

	baseMultiples := []int{}
	baseMultiplesSum := 0

	for m := 1; m <= cycle; m++ {
		if m%factorA == 0 || m%factorB == 0 {
			baseMultiples = append(baseMultiples, m)
			baseMultiplesSum += m
		}
	}

	numCycles := max / cycle
	remainder := max % cycle

	sum := triangleNumber(numCycles-1)*len(baseMultiples)*cycle + baseMultiplesSum*numCycles

	for _, m := range baseMultiples {
		if m >= remainder {
			break
		}

		sum += m + max - remainder
	}

	return fmt.Sprintf("%d", sum)
}
