package solutions

import "fmt"

func problem2() string {
	const max = 4000000

	a, b := 1, 2
	sum := 0

	for b <= max {
		if b&1 == 0 {
			sum += b
		}

		// Iterate fibonacci numbers
		b = a + b
		a = b - a
	}

	return fmt.Sprintf("%d", sum)
}
