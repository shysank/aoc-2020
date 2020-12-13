package puzzles

import "fmt"

type layout struct {
	grid map[coordinates]string
	size coordinates
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

func (l *layout) duplicate() *layout {
	newLayout := new(layout)
	newLayout.grid = make(map[coordinates]string)
	for k, v := range l.grid {
		newLayout.grid[k] = v
	}

	newLayout.size = l.size

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

func (l *layout) print() {
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
