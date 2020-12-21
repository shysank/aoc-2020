package puzzles

import (
	"fmt"
	"io"
	"strings"
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
type stringArrayResult []string

func (i intResult) Value() string {
	return fmt.Sprintf("%d", i)
}

func (i int64Result) Value() string {
	return fmt.Sprintf("%d", i)
}

func (s stringArrayResult) Value() string {
	var result string
	for _, v := range s {
		result += v + ","
	}
	return strings.Trim(result, ",")
}
