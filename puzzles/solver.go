package puzzles

import "io"

type Solver interface {
	Puzzle1(reader io.Reader) (Result, error)
	Puzzle2(reader io.Reader) (Result, error)
}

type Result interface {
	Value() string
}