package puzzles

import "io"

type RambunctiousRecitation struct {
	startingNos []int
	limit       int
}

func (r RambunctiousRecitation) Puzzle1(reader io.Reader) (Result, error) {
	nos := r.startingNos
	mem := initialMemory(nos)

	lastSpoken := nos[len(nos)-1]
	var nextVal int
	for i := len(nos) + 1; i <= r.limit; i++ {
		m := mem[lastSpoken]
		// not in memory
		if m == 0 {
			nextVal = 0
		} else {
			nextVal = (i -1) - m
		}
		mem[lastSpoken] = i - 1
		lastSpoken = nextVal
	}

	return intResult(nextVal), nil
}

func (r RambunctiousRecitation) Puzzle2(reader io.Reader) (Result, error) {
	return r.Puzzle1(reader)
}

func initialMemory(nos []int) map[int]int {
	mem := make(map[int]int)
	for i := 0; i < len(nos)-1; i++ {
		mem[nos[i]] = i + 1
	}
	return mem
}
