package puzzles

import (
	"fmt"
	"io"
)

type TobogganTrajectory struct {
	slopes []coordinates
}

type tobogganResult struct {
	trees []int
}

const (
	emptySpace string = "."
	tree       string = "#"
)

func (t TobogganTrajectory) Puzzle1(reader io.Reader) (Result, error) {
	return t.findTreeCount(reader)
}

func (t TobogganTrajectory) Puzzle2(reader io.Reader) (Result, error) {
	return t.findTreeCount(reader)
}

func (t TobogganTrajectory) findTreeCount(reader io.Reader) (Result, error) {
	tmap, err := ParseLayout(reader)
	if err != nil {
		return nil, err
	}

	result := tobogganResult{}

	for _, slope := range t.slopes {
		trees := tmap.treeCountForSlope(slope)
		result.trees = append(result.trees, trees)
	}

	return result, nil
}

func (l layout) isTreeAt(x, y int) bool {
	return l.valueAt(x, y) == tree
}

func (l layout) isEmptySpaceAt(x, y int) bool {
	return l.valueAt(x, y) == emptySpace
}

func (l layout) treeCountForSlope(slope coordinates) int {
	x := 0
	trees := 0
	for y := 0; y < l.size.y; {
		if l.isTreeAt(x, y) {
			trees++
		}
		x = (x + slope.x) % l.size.x
		y = y + slope.y
	}

	return trees
}

func (t tobogganResult) Value() string {
	product := 1
	for _, t := range t.trees {
		product *= t
	}
	return fmt.Sprintf("%d", product)
}
