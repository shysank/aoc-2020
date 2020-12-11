package puzzles

import (
	"io"
)

type binaryBoarding struct {
	rows int
	cols int
}

func (b *binaryBoarding) Puzzle1(reader io.Reader) (Result, error) {
	boardingPasses := ParseBoardingPasses(reader)

	var maxSeatId int
	for _, bp := range boardingPasses {
		rowCode := bp[:len(bp)-3]
		colCode := bp[len(bp)-3:]

		row := decode(rowCode, 0, 0, b.rows-1)
		col := decode(colCode, 0, 0, b.cols-1)

		seat := (row * b.cols) + col

		if seat > int(maxSeatId) {
			maxSeatId = seat
		}
	}

	return intResult(maxSeatId), nil
}

func (b *binaryBoarding) Puzzle2(reader io.Reader) (Result, error) {
	boardingPasses := ParseBoardingPasses(reader)

	sorted := make([]int, len(boardingPasses))
	for _, bp := range boardingPasses {
		rowCode := bp[:len(bp)-3]
		colCode := bp[len(bp)-3:]

		row := decode(rowCode, 0, 0, b.rows-1)
		col := decode(colCode, 0, 0, b.cols-1)

		seat := (row * b.cols) + col
		sorted = insertSort(sorted, seat)
	}

	var missingSeat int
	for i, seat := range sorted {
		if sorted[i+1]-seat == 2 {
			missingSeat = seat + 1
			break
		}
	}

	return intResult(missingSeat), nil
}

func insertSort(arr []int, num int) []int {
	var index int
	for i, n := range arr {
		if num > n && n != 0 {
			continue
		}
		index = i
		break
	}
	arr = append(arr[:index+1], arr[index:]...)
	arr[index] = num
	return arr
}

func decode(code string, index, min, max int) int {

	if min == max {
		return min
	}

	currChar := code[index]

	if currChar == 'F' || currChar == 'L' {
		max = (min + max) / 2
	}

	if currChar == 'B' || currChar == 'R' {
		min = ((min + max) / 2) + 1
	}

	index++

	return decode(code, index, min, max)

}
