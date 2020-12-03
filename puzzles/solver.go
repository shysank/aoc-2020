package puzzles

import "io"

type Solver interface {
	Puzzle1(reader io.Reader) (result, error)
	Puzzle2(reader io.Reader) (result, error)
}

type result struct {
	day1 day1Result
	day2 day2Result
}

// day 1 result types
type day1p1Result struct {
	n1 int64
	n2 int64
}

type day1p2Result struct {
	n1 int64
	n2 int64
	n3 int64
}

type day1Result struct {
	p1 day1p1Result
	p2 day1p2Result
}

// day 2 result types
type day2p1Result struct {
	validPasswords []string
}

type day2p2Result struct {
	validPasswords []string
}

type day2Result struct {
	p1 day2p1Result
	p2 day2p2Result
}
