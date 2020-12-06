package puzzles

import (
	"fmt"
	"io"
)

type customCustoms struct {
}

type answer struct {
	yeses      map[rune]int
	noOfPeople int
}

type answerCount int

func (c *customCustoms) Puzzle1(reader io.Reader) (Result, error) {
	answers := ParseCustomsAnswers(reader)

	var yesCount answerCount
	for _, a := range answers {
		yesCount += answerCount(len(a.yeses))
	}

	return yesCount, nil
}

func (c *customCustoms) Puzzle2(reader io.Reader) (Result, error) {
	answers := ParseCustomsAnswers(reader)

	var everyoneYesCount answerCount
	for _, a := range answers {
		for _, v := range a.yeses {
			if v == a.noOfPeople {
				everyoneYesCount++
			}
		}
	}

	return everyoneYesCount, nil
}

func (a answerCount) Value() string {
	return fmt.Sprintf("%d", a)
}
