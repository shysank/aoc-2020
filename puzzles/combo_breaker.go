package puzzles

import (
	"io"
)

type ComboBreaker struct {
}

const (
	modulus             = 20201227
	pubKeySubjectNumber = 7
)

func (c ComboBreaker) Puzzle1(reader io.Reader) (Result, error) {
	pubKeys, err := ParseToIntArray(reader)
	if err != nil {
		return nil, err
	}
	cardPubKey, doorPubKey := pubKeys[0], pubKeys[1]
	cardLoopSize := findLoopSize(cardPubKey)
	encKey := findEncKey(doorPubKey, cardLoopSize)
	return intResult(encKey), nil
}

func (c ComboBreaker) Puzzle2(reader io.Reader) (Result, error) {
	panic("implement me")
}

func findEncKey(pubKey, loopSize int) int {
	encKey := 1
	for i := 0; i < loopSize; i++ {
		transform(&encKey, pubKey)
	}
	return encKey
}

func findLoopSize(pubKey int) int {
	n := 1
	loopSize := 0
	for n != pubKey {
		transform(&n, pubKeySubjectNumber)
		loopSize++
	}
	return loopSize
}

func transform(n *int, subjectNumber int) {
	*n *= subjectNumber
	*n = *n % modulus
}
