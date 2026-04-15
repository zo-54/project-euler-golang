package main

import (
	"fmt"
	"os"
	"regexp"

	"project-euler-golang/solutions"
)

func main() {
	shouldBenchmark := false
	problems := []string{}

	benchmarkRegex := regexp.MustCompile(`^(?:--bench|-b)$`)
	problemRegex := regexp.MustCompile(`^\d+$`)

	for _, arg := range os.Args[1:] {
		switch {
		case problemRegex.MatchString(arg):
			problems = append(problems, arg)
		case benchmarkRegex.MatchString(arg):
			shouldBenchmark = true
		default:
			fmt.Printf("WARNING: unknown argument supplied: \"%s\". Ignoring...\n"+
				"\tUse -b or --bench to benchmark a single solution.\n"+
				"\tOther arguments should be a list of problem numbers separated by a space.\n\n",
				arg)
		}
	}

	if !shouldBenchmark {
		solutions.Run(problems)
		return
	}

	if len(problems) != 1 {
		fmt.Println("Must supply only one problem when benchmarking a solution.")
		return
	}

	solutions.Benchmark(problems[0])
}
