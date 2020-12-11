package puzzles

import (
	"io"
)

type handyHaversacks struct {
	colorToCheck bagType
}

type bagType string

type contains struct {
	qty   int
	color bagType
}

func (h handyHaversacks) Puzzle1(reader io.Reader) (Result, error) {
	bags, err := ParseBags(reader)
	if err != nil {
		return nil, err
	}

	var waysCount int
	for k, _ := range bags {
		if k == h.colorToCheck {
			continue
		}
		if traverseAndCheck(bags, k, h.colorToCheck) {
			waysCount++
		}
	}
	return intResult(waysCount), nil
}

func (h handyHaversacks) Puzzle2(reader io.Reader) (Result, error) {
	bags, err := ParseBags(reader)
	if err != nil {
		return nil, err
	}

	var bagsCount = intResult(totalBags(bags, h.colorToCheck))

	return bagsCount, nil
}

func traverseAndCheck(bags map[bagType][]contains, curr, toCheck bagType) bool {
	if curr == toCheck {
		return true
	}

	containsList := bags[curr]

	for _, c := range containsList {
		found := traverseAndCheck(bags, c.color, toCheck)
		if found {
			return true
		}
	}

	return false
}

func totalBags(bags map[bagType][]contains, curr bagType) int {
	containsList := bags[curr]

	if len(containsList) == 0 {
		return 0
	}

	bagCount := 0
	for _, c := range containsList {
		bagCount += c.qty + (c.qty * totalBags(bags, c.color))
	}

	return bagCount
}
