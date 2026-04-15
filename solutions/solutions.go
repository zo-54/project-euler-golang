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

		startTime := time.Now()

		soln := s()

		elapsed := time.Since(startTime)

		fmt.Printf("%s (completed in %dms)\n\n", soln, elapsed.Milliseconds())
	}
}
