package solutions

import (
	"fmt"
)

func problem7() string {
	const nthPrime uint64 = 10001

	return fmt.Sprintf("%d", getNthPrime(nthPrime))
}
