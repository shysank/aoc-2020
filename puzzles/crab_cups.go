package puzzles

import (
	"container/ring"
	"io"
	"strconv"
)

type CrabCups struct {
	label string
	moves int
	max   int
}

type cuplist struct {
	start  *ring.Ring
	curr   *ring.Ring
	picked []*ring.Ring

	cache []*ring.Ring
	max   int
}

const (
	ten        = 10
	oneMillion = 1000000
)

func (c CrabCups) Puzzle1(reader io.Reader) (Result, error) {
	l := c.label
	var labels []int
	for _, v := range l {
		lnum, _ := strconv.ParseInt(string(v), 10, 32)
		labels = append(labels, int(lnum))
	}

	cups := cuplist{start: ring.New(len(labels)), cache: make([]*ring.Ring, len(labels)), max: c.max}
	r := cups.start
	for i := 0; i < len(labels); i++ {
		r.Value = labels[i]
		cups.cache[labels[i]-1] = r
		r = r.Next()
	}

	for i := 0; i < c.moves; i++ {
		cups.nextClockwise()
		cups.pickN(3)
		dest := cups.findDestination()
		cups.arrange(dest)
	}

	var oneNode = cups.start
	for {
		if oneNode.Value.(int) == 1 {
			break
		}
		oneNode = oneNode.Next()
	}

	var result string
	var curr = oneNode.Next()
	for {
		if curr.Value.(int) == 1 {
			break
		}
		result += strconv.Itoa(curr.Value.(int))
		curr = curr.Next()
	}

	return stringResult(result), nil
}

func (c CrabCups) Puzzle2(reader io.Reader) (Result, error) {
	l := c.label
	var labels = make([]int, c.max)
	for i, v := range l {
		lnum, _ := strconv.ParseInt(string(v), 10, 32)
		labels[i] = int(lnum)
	}

	for i := 10; i <= c.max; i++ {
		labels[i-1] = i
	}

	cups := cuplist{start: ring.New(len(labels)), cache: make([]*ring.Ring, len(labels)), max: c.max}
	r := cups.start
	for i := 0; i < len(labels); i++ {
		r.Value = labels[i]
		cups.cache[labels[i]-1] = r
		r = r.Next()
	}

	for i := 0; i < c.moves; i++ {
		cups.nextClockwise()
		cups.pickN(3)
		dest := cups.findDestination()
		cups.arrange(dest)
	}

	for {
		if r.Value.(int) == 1 {
			break
		}
		r = r.Next()
	}

	result := r.Next().Value.(int) * r.Next().Next().Value.(int)

	return intResult(result), nil
}

func (c *cuplist) nextClockwise() {
	if c.curr == nil {
		c.curr = c.start
		return
	}
	c.curr = c.curr.Next()
}

func (c *cuplist) pickN(n int) {
	curr := c.curr
	var picked []*ring.Ring
	for i := 0; i < n; i++ {
		curr = curr.Next()
		picked = append(picked, curr)
	}
	c.picked = picked
}

func (c *cuplist) findDestination() int {
	d := c.curr.Value.(int) - 1
	for c.isAlreadyPicked(d) {
		d--
	}
	if d == 0 {
		return c.findMaxDest()
	}
	return d
}

func (c *cuplist) isAlreadyPicked(val int) bool {
	for _, p := range c.picked {
		if val == p.Value.(int) {
			return true
		}
	}
	return false
}

func (c *cuplist) findMaxDest() int {
	max := c.max
	for c.isAlreadyPicked(max) {
		max--
	}

	if max == c.curr.Value.(int) {
		max--
	}

	return max
}

func (c *cuplist) arrange(dest int) {
	pickedStart := c.picked[0]
	pickedEnd := c.picked[len(c.picked)-1]

	c.curr.Link(pickedEnd.Next())

	destNode := c.cache[dest-1]
	destNodeNext := destNode.Next()
	destNode.Link(pickedStart)
	pickedEnd.Link(destNodeNext)
}
