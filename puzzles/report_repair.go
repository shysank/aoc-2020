package puzzles

import (
	"errors"
	"fmt"
	"io"
)

type ReportRepair struct {
	Sum int
}

type reportResult struct {
	nos []int
}

func (e *ReportRepair) Puzzle1(reader io.Reader) (Result, error) {
	nos, err := ParseToIntArray(reader)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(nos); i++ {
		n1 := nos[i]
		for j := i + 1; j < len(nos); j++ {
			n2 := nos[j]

			if n1+n2 == e.Sum {
				return reportResult{[]int{n1, n2}}, nil
			}
		}
	}

	return nil, errors.New("Cannot find 2 nos with Sum 2020")
}

func (e *ReportRepair) Puzzle2(reader io.Reader) (Result, error) {
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
				if n1+n2+n3 == e.Sum {
					return reportResult{[]int{n1, n2, n3}}, nil
				}
			}
		}
	}

	return nil, errors.New("Cannot find 3 nos with Sum 2020")

}

func (r reportResult) Value() string {
	product := 1
	for _, n := range r.nos {
		product *= n
	}

	return fmt.Sprintf("%d", product)
}
