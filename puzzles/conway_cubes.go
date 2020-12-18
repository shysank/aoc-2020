package puzzles

import "io"

type ConwayCubes struct {
	cycles         int
	neighborsCache dimensionCache
}

type coordinate interface {
	neighbors() coordinateList
}

type coordinateList []coordinate

type coordinates2d struct {
	x, y int
}

type coordinates3d struct {
	coordinates2d
	z int
}

type coordinates4d struct {
	coordinates3d
	w int
}

type pocketDimension struct {
	grid           map[coordinate]string
	neighborsCache dimensionCache
}

type dimensionCache struct {
	limit int
	vals  map[coordinate]coordinateList
}

const (
	active   string = "#"
	inactive string = "."
)

var _ coordinate = coordinates2d{}
var _ coordinate = coordinates3d{}
var _ coordinate = coordinates4d{}

func (c ConwayCubes) Puzzle1(reader io.Reader) (Result, error) {
	l, err := ParseLayout(reader)
	if err != nil {
		return nil, err
	}

	c.neighborsCache = dimensionCache{
		limit: 100,
		vals:  map[coordinate]coordinateList{},
	}

	pd := toPocketDimension3D(l)
	for i := 0; i < c.cycles; i++ {
		newPd := pd.clone()
		pd.runCycle(newPd)
		pd = newPd
	}

	return intResult(pd.activeCount()), nil
}

func (c ConwayCubes) Puzzle2(reader io.Reader) (Result, error) {
	l, err := ParseLayout(reader)
	if err != nil {
		return nil, err
	}

	pd := toPocketDimension4D(l)
	for i := 0; i < c.cycles; i++ {
		newPd := pd.clone()
		pd.runCycle(newPd)
		pd = newPd
	}

	return intResult(pd.activeCount()), nil
}

func toPocketDimension3D(l *layout) *pocketDimension {
	p := pocketDimension{grid: map[coordinate]string{}}
	z := 0
	for k, v := range l.grid {
		p.grid[coordinates3d{coordinates2d{k.x, k.y}, z}] = v
	}
	p.neighborsCache = dimensionCache{vals: map[coordinate]coordinateList{}}
	return &p
}

func toPocketDimension4D(l *layout) *pocketDimension {
	p := pocketDimension{grid: map[coordinate]string{}}
	z := 0
	w := 0
	for k, v := range l.grid {
		p.grid[coordinates4d{coordinates3d{coordinates2d{k.x, k.y}, z}, w}] = v
	}
	p.neighborsCache = dimensionCache{vals: map[coordinate]coordinateList{}}
	return &p
}

func (p *pocketDimension) activeCount() int {
	activeCount := 0
	for k, _ := range p.grid {
		if p.grid[k] == active {
			activeCount++
		}
	}
	return activeCount
}

func (p *pocketDimension) runCycle(newPd *pocketDimension) {
	expandedGrid := make(map[coordinate]string)
	for k, _ := range p.grid {
		p.applyRules(k, p.grid, newPd.grid, expandedGrid)
	}
	for k, _ := range expandedGrid {
		p.applyRules(k, p.grid, newPd.grid, map[coordinate]string{})
	}
}

func (p *pocketDimension) clone() *pocketDimension {
	pClone := pocketDimension{grid: map[coordinate]string{}}
	for k, v := range p.grid {
		pClone.grid[k] = v
	}
	pClone.neighborsCache = p.neighborsCache
	return &pClone
}

func (p *pocketDimension) applyRules(c coordinate, old, new, expanded map[coordinate]string) {
	var neighbors []coordinate
	if n, ok := p.neighborsCache.vals[c]; ok {
		neighbors = n
	} else {
		neighbors = c.neighbors()
		p.neighborsCache.vals[c] = neighbors
	}

	activeCount := 0
	for _, n := range neighbors {
		if _, ok := old[n]; !ok {
			expanded[n] = inactive
			continue
		}
		if old[n] == active {
			activeCount++
		}
	}

	if old[c] == active {
		if !(activeCount == 2 || activeCount == 3) {
			new[c] = inactive
		}
	} else {
		if activeCount == 3 {
			new[c] = active
		} else {
			new[c] = inactive
		}
	}
}

func (c coordinates4d) neighbors() coordinateList {
	neighbors3D := c.coordinates3d.neighbors()

	var neighbors coordinateList
	neighbors = threeDFourD(neighbors3D, c.w)
	neighbors = append(neighbors, threeDFourD(neighbors3D, c.w+1)...)
	neighbors = append(neighbors, threeDFourD(neighbors3D, c.w-1)...)
	neighbors = append(neighbors, coordinates4d{coordinates3d{coordinates2d{c.x, c.y}, c.z}, c.w + 1},
		coordinates4d{coordinates3d{coordinates2d{c.x, c.y}, c.z}, c.w - 1})
	return neighbors
}

func (c coordinates3d) neighbors() coordinateList {
	var neighbors coordinateList

	neighbors2D := c.coordinates2d.neighbors()
	neighbors = twoDThreeD(neighbors2D, c.z)
	neighbors = append(neighbors, twoDThreeD(neighbors2D, c.z+1)...)
	neighbors = append(neighbors, twoDThreeD(neighbors2D, c.z-1)...)
	neighbors = append(neighbors, coordinates3d{coordinates2d{c.x, c.y}, c.z - 1},
		coordinates3d{coordinates2d{c.x, c.y}, c.z + 1})

	return neighbors
}

func (c coordinates2d) neighbors() coordinateList {
	x, y := c.x, c.y
	level := 1

	var neighborsList coordinateList
	neighborsList = append(neighborsList, coordinates2d{x - level, y - level}, coordinates2d{x, y - level},
		coordinates2d{x + level, y - level})
	neighborsList = append(neighborsList, coordinates2d{x - level, y}, coordinates2d{x + level, y})
	neighborsList = append(neighborsList, coordinates2d{x - level, y + level}, coordinates2d{x, y + level},
		coordinates2d{x + level, y + level})

	return neighborsList
}

func twoDThreeD(coords2d coordinateList, z int) []coordinate {
	var coords3d []coordinate
	for _, c := range coords2d {
		c2d := c.(coordinates2d)
		coords3d = append(coords3d, coordinates3d{coordinates2d{c2d.x, c2d.y}, z})
	}
	return coords3d
}

func threeDFourD(coords3d coordinateList, w int) []coordinate {
	var coords4d []coordinate
	for _, c := range coords3d {
		c3d := c.(coordinates3d)
		coords4d = append(coords4d, coordinates4d{coordinates3d{coordinates2d{c3d.x, c3d.y}, c3d.z}, w})
	}
	return coords4d
}
