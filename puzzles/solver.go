package puzzles

import "io"

type Solver interface {
	Puzzle1(reader io.Reader) (result, error)
	Puzzle2(reader io.Reader) (result, error)
}

type result struct {
	day1 day1Result
}

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
