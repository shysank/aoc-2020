package puzzles

import (
	"io"
)

type SeatingSystem struct {
}

const (
	empty    string = "L"
	floor    string = "."
	occupied string = "#"
)

func (s SeatingSystem) Puzzle1(reader io.Reader) (Result, error) {
	l, err := ParseLayout(reader)
	if err != nil {
		return nil, err
	}
	var newLayout, oldLayout *layout
	oldLayout = l.duplicate()
	for {
		newLayout = oldLayout.duplicate()
		var unchanged = true
		for y := 0; y < l.size.y; y++ {
			for x := 0; x < l.size.x; x++ {
				changed := applyRules(oldLayout, newLayout, coordinates{x, y})
				unchanged = unchanged && !changed
			}
		}

		if unchanged {
			break
		}

		oldLayout = newLayout
	}

	return intResult(occupiedSeatsCount(newLayout)), nil

}

func (s SeatingSystem) Puzzle2(reader io.Reader) (Result, error) {
	l, err := ParseLayout(reader)
	if err != nil {
		return nil, err
	}
	var newLayout, oldLayout *layout
	oldLayout = l.duplicate()
	for {
		newLayout = oldLayout.duplicate()
		var unchanged = true
		for y := 0; y < l.size.y; y++ {
			for x := 0; x < l.size.x; x++ {
				changed := applyNewRules(oldLayout, newLayout, coordinates{x, y})
				unchanged = unchanged && !changed
			}
		}

		if unchanged {
			break
		}

		oldLayout = newLayout
	}

	return intResult(occupiedSeatsCount(newLayout)), nil
}

func applyRules(old, newLayout *layout, c coordinates) bool {
	seat := old.valueAt(c.x, c.y)
	if seat == floor {
		return false
	}
	adjs := old.getAdjacentCoords(c, 1)

	var occupiedCount int
	for _, c := range adjs {
		if old.valueAt(c.x, c.y) == occupied {
			occupiedCount++
		}
	}

	if seat == empty {
		if occupiedCount == 0 {
			newLayout.setValueAt(c, occupied)
			return true
		}
	}
	if seat == occupied {
		if occupiedCount >= 4 {
			newLayout.setValueAt(c, empty)
			return true
		}
	}
	return false

}

func applyNewRules(old, newLayout *layout, c coordinates) bool {
	seat := old.valueAt(c.x, c.y)
	if seat == floor {
		return false
	}

	var occupiedCount int
	var occupiedFound bool

	occupiedFound = findFirstOccupied(old, c, nextTopLeft)
	if occupiedFound {
		occupiedCount++
	}

	occupiedFound = findFirstOccupied(old, c, nextTop)
	if occupiedFound {
		occupiedCount++
	}

	occupiedFound = findFirstOccupied(old, c, nextTopRight)
	if occupiedFound {
		occupiedCount++
	}

	occupiedFound = findFirstOccupied(old, c, nextLeft)
	if occupiedFound {
		occupiedCount++
	}

	occupiedFound = findFirstOccupied(old, c, nextRight)
	if occupiedFound {
		occupiedCount++
	}

	occupiedFound = findFirstOccupied(old, c, nextBotLeft)
	if occupiedFound {
		occupiedCount++
	}

	occupiedFound = findFirstOccupied(old, c, nextBot)
	if occupiedFound {
		occupiedCount++
	}

	occupiedFound = findFirstOccupied(old, c, nextBotRight)
	if occupiedFound {
		occupiedCount++
	}

	if seat == empty {
		if occupiedCount == 0 {
			newLayout.setValueAt(c, occupied)
			return true
		}
	}
	if seat == occupied {
		if occupiedCount >= 5 {
			newLayout.setValueAt(c, empty)
			return true
		}
	}
	return false

}

func findFirstOccupied(l *layout, c coordinates, next func(x, y, l int) coordinates) bool {
	for i := 1; i < l.size.x; i++ {
		nxt := next(c.x, c.y, i)
		if l.isOutOfBoundary(nxt) || l.valueAt(nxt.x, nxt.y) == empty {
			return false
		}

		if l.valueAt(nxt.x, nxt.y) == occupied {
			return true
		}

	}
	return false
}

func nextTopLeft(x, y, level int) coordinates {
	return coordinates{x - level, y - level}
}

func nextTop(x, y, level int) coordinates {
	return coordinates{x, y - level}
}

func nextTopRight(x, y, level int) coordinates {
	return coordinates{x + level, y - level}
}

func nextRight(x, y, level int) coordinates {
	return coordinates{x + level, y}
}

func nextLeft(x, y, level int) coordinates {
	return coordinates{x - level, y}
}

func nextBotLeft(x, y, level int) coordinates {
	return coordinates{x - level, y + level}
}

func nextBot(x, y, level int) coordinates {
	return coordinates{x, y + level}
}

func nextBotRight(x, y, level int) coordinates {
	return coordinates{x + level, y + level}
}

func occupiedSeatsCount(l *layout) int {
	count := 0
	for _, v := range l.grid {
		if v == occupied {
			count++
		}
	}
	return count
}
