package solutions

import (
	"fmt"
)

func problem6() string {
	const (
		minN uint64 = 1
		maxN uint64 = 100
	)

	var (
		sumOfSquares uint64 = 0
		sum          uint64 = 0
	)

	for n := minN; n <= maxN; n++ {
		sumOfSquares += n * n
		sum += n
	}

	return fmt.Sprintf("%d", sum*sum-sumOfSquares)
}
