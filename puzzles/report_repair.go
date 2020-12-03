package puzzles

import (
	"errors"
	"fmt"
	"io"
)

type reportRepair struct {
	sum int
}

type reportResult struct {
	nos []int
}

func (e *reportRepair) Puzzle1(reader io.Reader) (Result, error) {
	nos, err := ParseToIntArray(reader)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(nos); i++ {
		n1 := nos[i]
		for j := i + 1; j < len(nos); j++ {
			n2 := nos[j]

			if n1+n2 == e.sum {
				return reportResult{[]int{n1, n2}}, nil
			}
		}
	}

	return nil, errors.New("Cannot find 2 nos with sum 2020")
}

func (e *reportRepair) Puzzle2(reader io.Reader) (Result, error) {
	nos, err := ParseToIntArray(reader)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(nos); i++ {
		n1 := nos[i]
		for j := i + 1; j < len(nos); j++ {
			n2 := nos[j]
			for k := j + 1; k < len(nos); k++ {
				n3 := nos[k]
				if n1+n2+n3 == e.sum {
					return reportResult{[]int{n1, n2, n3}}, nil
				}
			}
		}
	}

	return nil, errors.New("Cannot find 3 nos with sum 2020")

}

func (r reportResult) Value() string {
	product := 1
	for _, n := range r.nos {
		product *= n
	}

	return fmt.Sprintf("%d", product)
}
