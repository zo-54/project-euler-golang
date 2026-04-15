# Project Euler (Go)

[Project Euler](https://projecteuler.net) is a fantastic site that offers complex mathematical and computational problems that are available to be solved for free. I highly recommend attempting these problems yourself before viewing more of this repository, as my solutions to the problems will be contained here.

Since I began working on these problems around 2019, I have made significant progress on the archive using JavaScript/TypeScript. While I am still working on that project, I have decided to create a fresh account and work through the problems again from the beginning, this time using Go.

I am also placing a heavy precedence on making the solutions fast, as well as adaptable to values other than the ones provided by the problem.

## Usage

### Run the program

```bash
go run . <problem numbers (optional)>
```

By default, the program will output all solutions. Providing problem numbers will limit the program to only outputting the corresponding solutions.

Problem numbers that correspond to problems that I have not yet solved will result in a message stating that there is no solution available for the given problem.

#### Examples

- To get all solutions: `go run .`
- To get the solution for problem 7: `go run . 7`
- To get the solutions for problem 35, 38, and 42: `go run . 35 38 42`

### Build the program

```bash
go build .
```

The resulting executable can be used in place of `go run` when running the program.
