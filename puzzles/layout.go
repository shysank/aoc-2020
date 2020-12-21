package puzzles

import "fmt"

type layout struct {
	grid    map[coordinates]string
	size    coordinates
	borders map[string]string
}

type coordinates struct {
	x, y int
}

func (l *layout) valueAt(x, y int) string {
	return l.grid[coordinates{x, y}]
}

func (l *layout) setValueAt(coords coordinates, value string) {
	l.grid[coords] = value
}

func (l *layout) Clone() *layout {
	newLayout := new(layout)
	newLayout.grid = make(map[coordinates]string)
	for k, v := range l.grid {
		newLayout.grid[k] = v
	}

	newLayout.size = l.size
	newLayout.border()

	return newLayout
}

func (l *layout) getAdjacentCoords(coords coordinates, level int) []coordinates {
	x, y := coords.x, coords.y
	top := []coordinates{{x - level, y - level}, {x, y - level}, {x + level, y - level}}
	mid := []coordinates{{x - level, y}, {x + level, y}}
	bot := []coordinates{{x - level, y + level}, {x, y + level}, {x + level, y + level}}

	all := append(top, append(mid, bot...)...)

	var validCoords []coordinates
	for _, c := range all {
		if !l.isOutOfBoundary(c) {
			validCoords = append(validCoords, c)
		}

	}
	return validCoords
}

func (l *layout) isOutOfBoundary(coords coordinates) bool {
	x, y := coords.x, coords.y
	if x < 0 || x >= l.size.x {
		return true
	}

	if y < 0 || y >= l.size.y {
		return true
	}

	return false
}

func (l *layout) border() {
	b := make(map[string]string)
	y := 0
	for x := 0; x < l.size.x; x++ {
		b[top] += l.valueAt(x, y)
	}

	y = l.size.y - 1
	for x := 0; x < l.size.x; x++ {
		b[bot] += l.valueAt(x, y)
	}

	x := 0
	for y := 0; y < l.size.y; y++ {
		b[lt] += l.valueAt(x, y)
	}

	x = l.size.x - 1
	for y := 0; y < l.size.y; y++ {
		b[rt] += l.valueAt(x, y)
	}
	l.borders = b
}

func (l *layout) RotateRight() {
	rotatedGrid := map[coordinates]string{}
	for k, v := range l.grid {
		c := coordinates{l.size.y - k.y - 1, k.x}
		rotatedGrid[c] = v
	}

	l.grid = rotatedGrid
	l.border()
}

func (l *layout) RotateLeft() {
	rotatedGrid := map[coordinates]string{}
	for k, v := range l.grid {
		c := coordinates{k.y, l.size.x - k.x - 1}
		rotatedGrid[c] = v
	}

	l.grid = rotatedGrid
	l.border()
}

func (l *layout) Flip() {
	rotatedGrid := map[coordinates]string{}
	for k, v := range l.grid {
		c := coordinates{k.x, l.size.y - k.y - 1}
		rotatedGrid[c] = v
	}

	l.grid = rotatedGrid
	l.border()
}

func (l *layout) removeBorders() {
	newGrid := make(map[coordinates]string)
	for y := 1; y < l.size.y-1; y++ {
		for x := 1; x < l.size.x-1; x++ {
			newGrid[coordinates{x - 1, y - 1}] = l.grid[coordinates{x, y}]
		}
	}
	l.grid = newGrid
	l.size = coordinates{l.size.x - 2, l.size.y - 2}
}

func (l *layout) Print() {
	layoutString := "\n"
	for y := 0; y < l.size.y; y++ {
		row := ""
		for x := 0; x < l.size.x; x++ {
			row = row + l.valueAt(x, y)
		}
		layoutString = layoutString + row + "\n"
	}

	fmt.Println(layoutString)
}
