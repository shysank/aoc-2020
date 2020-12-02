package puzzles

import (
	"errors"
	"io"
)

type expenseReport struct {
	sum int
}

func (e *expenseReport) Puzzle1(reader io.Reader) (result, error) {
	nos, err := ParseToInt64Array(reader)
	if err != nil {
		return result{}, err
	}

	for i := 0; i < len(nos); i++ {
		n1 := nos[i]
		for j := i + 1; j < len(nos); j++ {
			n2 := nos[j]

			if n1+n2 == int64(e.sum) {
				return result{day1: day1Result{p1: day1p1Result{n1: n1, n2: n2}}}, nil
			}
		}
	}

	return result{}, errors.New("Cannot find 2 nos with sum 2020")
}

func (e *expenseReport) Puzzle2(reader io.Reader) (result, error) {
	nos, err := ParseToInt64Array(reader)
	if err != nil {
		return result{}, err
	}

	for i := 0; i < len(nos); i++ {
		n1 := nos[i]
		for j := i + 1; j < len(nos); j++ {
			n2 := nos[j]
			for k := j + 1; k < len(nos); k++ {
				n3 := nos[k]
				if n1+n2+n3 == int64(e.sum) {
					return result{day1: day1Result{p2: day1p2Result{n1: n1, n2: n2, n3: n3}}}, nil
				}
			}
		}
	}

	return result{}, errors.New("Cannot find 3 nos with sum 2020")

	return result{}, nil

}
