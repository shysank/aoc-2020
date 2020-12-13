package puzzles

import (
	"fmt"
	"io"
)

type Solver interface {
	Puzzle1(reader io.Reader) (Result, error)
	Puzzle2(reader io.Reader) (Result, error)
}

type Result interface {
	Value() string
}
type intResult int
type int64Result int64

func (i intResult) Value() string {
	return fmt.Sprintf("%d", i)
}

func (i int64Result) Value() string {
	return fmt.Sprintf("%d", i)
}
