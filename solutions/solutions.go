package solutions

import (
	"fmt"
	"maps"
	"slices"
	"time"
)

type solution func() string

func Run(problems []string) {
	solutionsMap := map[string]solution{
		"1": problem1,
	}

	if len(problems) == 0 {
		problems = slices.Collect(maps.Keys(solutionsMap))
	}

	for _, p := range problems {
		s, ok := solutionsMap[p]

		if !ok {
			fmt.Printf("No solution available for problem %s\n\n", p)
			continue
		}

		fmt.Printf("======== Problem %s ========\n", p)

		soln, elapsed := runSolution(s)

		fmt.Printf("%s (completed in %s)\n\n", soln, elapsed)
	}
}

func runSolution(s solution) (string, time.Duration) {
	startTime := time.Now()

	soln := s()

	elapsed := time.Since(startTime)

	return soln, elapsed
}
