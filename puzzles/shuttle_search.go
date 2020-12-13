package puzzles

import (
	"io"
	"math"
)

type ShuttleSearch struct {
}

type notes struct {
	earliestTime int
	busIDs       []int
}

func (s ShuttleSearch) Puzzle1(reader io.Reader) (Result, error) {
	n, err := ParseNotes(reader)
	if err != nil {
		return nil, err
	}

	minWaitTime := math.MaxInt32
	minWaitID := 0
	for _, id := range n.busIDs {
		if id == -1 {
			continue
		}
		q := n.earliestTime / id
		nearestTime := (q + 1) * id
		if nearestTime < minWaitTime {
			minWaitTime = nearestTime
			minWaitID = id
		}
	}

	return intResult((minWaitTime - n.earliestTime) * minWaitID), nil
}

func (s ShuttleSearch) Puzzle2(reader io.Reader) (Result, error) {
	n, err := ParseNotes(reader)
	if err != nil {
		return nil, err
	}
	eqs := toEqs(n.busIDs)
	x := eqs.Solve()
	return int64Result(x), nil
}

func toEqs(ids []int) (eqs CRTEquations) {
	for i, id := range ids {
		if id == -1 {
			continue
		}
		eqs = append(eqs, CRTEquation{id - i, id})
	}
	return eqs
}
