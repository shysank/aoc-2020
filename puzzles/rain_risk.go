package puzzles

import "io"

type rainRisk struct {
}

func (r rainRisk) Puzzle1(reader io.Reader) (Result, error) {
	ins, err := ParseNavInstructions(reader)
	if err != nil {
		return nil, err
	}

	ship := &vector{
		pos: coordinates{0, 0},
		dir: east,
	}
	for _, i := range ins {
		ship.handleAction(i)
	}

	return intResult(ship.pos.manhattanDistance()), nil
}

func (r rainRisk) Puzzle2(reader io.Reader) (Result, error) {
	ins, err := ParseNavInstructions(reader)
	if err != nil {
		return nil, err
	}

	ship := &coordinates{0, 0}
	waypoint := &coordinates{10, 1}

	for _, i := range ins {
		if i.isDirection() {
			waypoint.handleDirection(dirMapping[i.action], i.value)
		}

		if i.isTurn() {
			if i.action == left {
				waypoint.rotate(i.value)
			}

			if i.action == right {
				waypoint.rotate(i.value * -1)
			}
		}

		if i.isForward() {
			ship.add(coordinates{i.value * waypoint.x, i.value * waypoint.y})
		}
	}
	return intResult(ship.manhattanDistance()), nil
}
