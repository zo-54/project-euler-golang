package solutions

import (
	"fmt"
	"maps"
	"slices"
	"time"
)

type solutionFunc func() string

var solutionsMap = map[string]solutionFunc{
	"1": problem1,
	"2": problem2,
	"3": problem3,
	"4": problem4,
	"5": problem5,
	"6": problem6,
	"7": problem7,
	"8": problem8,
}

func Run(problems []string) {
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

		solutionString, elapsed := runSolution(s)

		fmt.Printf("%s (completed in %s)\n\n", solutionString, elapsed)
	}
}

func Benchmark(problem string) {
	solutionFunc, ok := solutionsMap[problem]

	if !ok {
		fmt.Printf("No solution available for problem %s\n\n", problem)
	}

	_, firstElapsed := runSolution(solutionFunc)

	executions := 1

	min := firstElapsed
	max := firstElapsed
	sum := firstElapsed

	const (
		maxDuration   = 5 * time.Minute
		maxExecutions = 1000
	)

	for executions < maxExecutions && sum+max < maxDuration {
		_, e := runSolution(solutionFunc)

		sum += e

		if e < min {
			min = e
		}

		if e > max {
			max = e
		}

		executions++

		resetMaths()
	}

	ave := sum / time.Duration(executions)

	fmt.Printf("Problem %s benchmark\n"+
		"\tRan solution %d times\n"+
		"\tMin: %s\n"+
		"\tAve: %s\n"+
		"\tMax: %s\n",
		problem, executions, min, ave, max,
	)
}

func runSolution(s solutionFunc) (string, time.Duration) {
	startTime := time.Now()

	soln := s()

	elapsed := time.Since(startTime)

	return soln, elapsed
}
