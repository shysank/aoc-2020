package puzzles

import (
	"io"
)

type CustomCustoms struct {
}

type answer struct {
	yeses      map[rune]int
	noOfPeople int
}

func (c *CustomCustoms) Puzzle1(reader io.Reader) (Result, error) {
	answers := ParseCustomsAnswers(reader)

	var yesCount int
	for _, a := range answers {
		yesCount += len(a.yeses)
	}

	return intResult(yesCount), nil
}

func (c *CustomCustoms) Puzzle2(reader io.Reader) (Result, error) {
	answers := ParseCustomsAnswers(reader)

	var everyoneYesCount int
	for _, a := range answers {
		for _, v := range a.yeses {
			if v == a.noOfPeople {
				everyoneYesCount++
			}
		}
	}

	return intResult(everyoneYesCount), nil
}
