package puzzles

import "math"

type direction int

type navInstruction struct {
	action string
	value  int
}

type vector struct {
	pos coordinates
	dir direction
}

const (
	north direction = 0
	east  direction = 1
	south direction = 2
	west  direction = 3
	right           = "R"
	left            = "L"

	northStr string = "N"
	eastStr  string = "E"
	southStr string = "S"
	westStr  string = "W"

	forward string = "F"
)

var dirMapping = map[string]direction{
	northStr: north, eastStr: east, southStr: south, westStr: west,
}

func (v *vector) handleAction(n navInstruction) {
	if n.isTurn() {
		v.handleTurn(n.action, n.value)
	}

	if n.isDirection() {
		v.pos.handleDirection(dirMapping[n.action], n.value)
	}

	if n.isForward() {
		v.pos.handleDirection(v.dir, n.value)
	}
}

func (v *vector) handleTurn(turn string, degrees int) {
	quadrants := degrees / 90
	if turn == right {
		newDir := (int(v.dir) + quadrants) % 4
		v.dir = direction(newDir)
	}

	if turn == left {
		newDir := (int(v.dir) - quadrants + 4) % 4
		v.dir = direction(newDir)
	}
}

func (c *coordinates) handleDirection(dir direction, val int) {
	switch dir {
	case north:
		c.y += val
	case east:
		c.x += val
	case south:
		c.y -= val
	case west:
		c.x -= val

	}
}

func (c *coordinates) manhattanDistance() int {
	distX := c.x
	if distX < 0 {
		distX *= -1
	}

	distY := c.y
	if distY < 0 {
		distY *= -1
	}

	return distX + distY
}

func (n navInstruction) isTurn() bool {
	return n.action == right || n.action == left
}

func (n navInstruction) isDirection() bool {
	return n.action == northStr || n.action == southStr || n.action == eastStr || n.action == westStr
}

func (n navInstruction) isForward() bool {
	return n.action == forward
}

func (c *coordinates) rotate(radians int) {
	degrees := float64(radians) * (math.Pi / 180)
	orig := coordinates{c.x, c.y}
	c.x = orig.x*int(math.Cos(degrees)) - orig.y*int(math.Sin(degrees))
	c.y = orig.x*int(math.Sin(degrees)) + orig.y*int(math.Cos(degrees))
}

func (c *coordinates) add(c1 coordinates) {
	c.x = c.x + c1.x
	c.y = c.y + c1.y

}
