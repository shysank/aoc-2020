package puzzles

import (
	"fmt"
	"io"
	"sort"
)

type adapterArray struct {
}

type jolts int

func (a adapterArray) Puzzle1(reader io.Reader) (Result, error) {
	joltages, err := ParseToIntArray(reader)
	if err != nil {
		return nil, err
	}

	sort.Ints(joltages)
	joltDiffs := make(map[int]int)

	currJoltage := 0
	for i := 0; i < len(joltages); i++ {
		diff := joltages[i] - currJoltage
		joltDiffs[diff] = joltDiffs[diff] + 1
		currJoltage = joltages[i]
	}

	joltDiffs[3] = joltDiffs[3] + 1

	return jolts(joltDiffs[3] * joltDiffs[1]), nil
}

func (a adapterArray) Puzzle2(reader io.Reader) (Result, error) {
	joltages, err := ParseToIntArray(reader)
	if err != nil {
		return nil, err
	}

	joltages = append(joltages, 0)

	sort.Ints(joltages)

	routes := 1
	i := len(joltages) - 1

	for {

		if i == 0 {
			break
		}

		subIndex := i - 1
		curr := joltages[i]
		if i-2 >= 0 && curr-joltages[i-2] <= 3 {
			subIndex = findSubSequenceStartIndex(joltages, i)
		}

		if subIndex != i-1 {
			subSequence := joltages[subIndex : i+1]
			w := waysInSubSequence(subSequence, len(subSequence)-1)
			routes = routes * w
		}

		i = subIndex

	}

	return jolts(routes), nil
}

func findSubSequenceStartIndex(arr []int, index int) int {
	i := index - 2

	if i-2 >= 0 && arr[i]-arr[i-2] <= 3 {
		i = i - 2
	}

	if i-1 >= 0 {
		return i - 1
	}
	return i
}

func waysInSubSequence(arr []int, index int) int {
	if index == 0 {
		return 1
	}

	curr := arr[index]
	prevRoutes := waysInSubSequence(arr, index-1)
	var prev1, prev2 int
	if index-2 >= 0 && curr-arr[index-2] <= 3 {
		prev1 = waysInSubSequence(arr, index-2)
	}

	if index-3 >= 0 && curr-arr[index-3] <= 3 {
		prev2 = waysInSubSequence(arr, index-3)
	}

	return prevRoutes + prev1 + prev2
}

func (j jolts) Value() string {
	return fmt.Sprintf("%d", j)
}
