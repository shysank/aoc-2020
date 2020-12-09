package puzzles

import (
	"errors"
	"fmt"
	"io"
	"sort"
)

type encodingError struct {
	preambleLength int
	invalidNum     int64
}

type xmasNum int

type sortInt64 []int64

func (e *encodingError) Puzzle1(reader io.Reader) (Result, error) {
	nos, err := ParseToInt64Array(reader)
	if err != nil {
		return nil, err
	}

	for i := e.preambleLength; i < len(nos); i++ {
		preamble := nos[i-e.preambleLength : i]
		if !isXMASValid(preamble, nos[i]) {
			e.invalidNum = nos[i]
			return xmasNum(nos[i]), nil
		}
	}
	return nil, errors.New("Cannot find XMAS weakness")
}

func (e *encodingError) Puzzle2(reader io.Reader) (Result, error) {
	nos, err := ParseToInt64Array(reader)
	if err != nil {
		return nil, err
	}

	c := findContinuousSet(nos, e.invalidNum)
	sort.Sort(sortInt64(c))
	min, max := c[0], c[len(c)-1]
	return xmasNum(min + max), nil
}

func isXMASValid(preamble []int64, num int64) bool {
	for i := 0; i < len(preamble); i++ {
		for j := i + 1; j < len(preamble); j++ {
			if num == preamble[i]+preamble[j] {
				return true
			}
		}
	}
	return false
}

func findContinuousSet(nos []int64, num int64) []int64 {
	for i := 0; i < len(nos); i++ {
		sum := nos[i]
		j := i + 1
		for {
			sum += nos[j]
			if sum == num {
				return nos[i : j+1]
			}

			if sum > num {
				break
			}
			j++
		}
	}
	return []int64{}
}

func (s sortInt64) Len() int {
	return len(s)
}
func (s sortInt64) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s sortInt64) Less(i, j int) bool {
	return s[i] < s[j]
}

func (x xmasNum) Value() string {
	return fmt.Sprintf("%d", x)
}
