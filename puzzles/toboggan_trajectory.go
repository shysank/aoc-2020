package puzzles

import (
	"fmt"
	"io"
)

type tobogganTrajectory struct {
	slopes []tobogganCoordinates
}

type tobogganMap struct {
	grid map[tobogganCoordinates]tobogganElement
	size tobogganCoordinates
}

type tobogganCoordinates struct {
	x, y int
}

type tobogganElement string

type tobogganResult struct {
	trees []int
}

const (
	emptySpace tobogganElement = "."
	tree       tobogganElement = "#"
)

func (t tobogganTrajectory) Puzzle1(reader io.Reader) (Result, error) {
	return t.findTreeCount(reader)
}

func (t tobogganTrajectory) Puzzle2(reader io.Reader) (Result, error) {
	return t.findTreeCount(reader)
}

func (t tobogganTrajectory) findTreeCount(reader io.Reader) (Result, error) {
	tmap, err := ParseTobogganMap(reader)
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

func (t tobogganMap) treeCountForSlope(slope tobogganCoordinates) int {
	x := 0
	trees := 0
	for y := 0; y < t.size.y; {
		if t.isTreeAt(x, y) {
			trees++
		}
		x = (x + slope.x) % t.size.x
		y = y + slope.y
	}

	return trees
}

func (t tobogganMap) valueAt(x, y int) tobogganElement {
	return t.grid[tobogganCoordinates{x, y}]
}

func (t tobogganMap) isTreeAt(x, y int) bool {
	return t.valueAt(x, y) == tree
}

func (t tobogganMap) isEmptySpaceAt(x, y int) bool {
	return t.valueAt(x, y) == emptySpace
}

func (t tobogganResult) Value() string {
	product := 1
	for _, t := range t.trees {
		product *= t
	}
	return fmt.Sprintf("%d", product)
}
