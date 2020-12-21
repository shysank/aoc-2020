package puzzles

import (
	"fmt"
	"io"
	"math"
	"strconv"
)

type JurassicJigsaw struct {
}

type tile struct {
	id      int
	layout  *layout
	matches []match
}

type match struct {
	id      int
	reverse bool
	dir     string
}

type tiles map[int]*tile

const (
	top = "top"
	rt  = "right"
	lt  = "left"
	bot = "bot"
)

var ops = []func(t *tile){nil, rotateRight, rotateRight, rotateRight, rotateRight, flip, rotateRight, rotateRight, rotateRight}

var seaMonsterLayout = `                  # 
#    ##    ##    ###
 #  #  #  #  #  #  `

func (j JurassicJigsaw) Puzzle1(reader io.Reader) (Result, error) {
	tilesList, err := ParseTiles(reader)
	if err != nil {
		return nil, err
	}

	tilesList.findBorders()
	tilesList.findMatches()

	product := 1
	for id, tl := range tilesList {
		if len(tl.matches) == 2 {
			product *= id
		}
	}

	return intResult(product), nil
}

func (j JurassicJigsaw) Puzzle2(reader io.Reader) (Result, error) {
	tilesList, err := ParseTiles(reader)
	if err != nil {
		return nil, err
	}

	tilesList.findBorders()
	tilesList.findMatches()

	// bootstrap tile is one of the borders that is already aligned (reverse = true for all matches)
	bootstrap := tilesList.findBootstrapTile()

	alignedTiles := make(map[int]bool)
	// alignTiles aligns all the tiles in the right direction after rotating and flipping
	alignTiles(bootstrap, tilesList, alignedTiles)

	// topLeft is the top left tile in the puzzle
	topLeft := tilesList.findTopLeftTile()

	origin := coordinates{0, 0}
	tilesMatrix := make(map[coordinates]int)
	tilesMatrix[origin] = topLeft.id
	// tilesMatrix aligns all the tiles in the form of grid with each tile's position determined by its alignment
	tilesList.formMatrix(topLeft, origin, tilesMatrix, map[int]bool{})

	// merge all tiles removing the borders into one big tile
	mergedTile := tilesList.mergeMatrix(tilesMatrix)

	// parse sea monster layout
	sml := ParseLayoutFromString(seaMonsterLayout)
	sml.onlyHave("#")

	// check for sea monsters in mergedTile by rotating and flipping and calculate roughness
	var roughness int
	for _, op := range ops {
		if op != nil {
			op(mergedTile)
		}
		if smlCoords, ok := mergedTile.layout.checkForSeaMonster(sml); ok {
			for k, v := range mergedTile.layout.grid {
				if v == "#" && !smlCoords[k] {
					roughness++
				}
			}
			break
		}
	}

	return intResult(roughness), nil
}

func (ts tiles) mergeMatrix(matrix map[coordinates]int) *tile {
	t := tile{}
	l := layout{}
	grid := make(map[coordinates]string)
	size := math.Sqrt(float64(len(matrix)))
	var tileSize coordinates
	for y := 0; y < int(size); y++ {
		for x := 0; x < int(size); x++ {
			tileId := matrix[coordinates{x, y}]
			t := ts[tileId]
			t.layout.removeBorders()
			tileSize = t.layout.size
			for k, v := range t.layout.grid {
				newCoords := coordinates{x*t.layout.size.x + k.x, y*t.layout.size.y + k.y}
				grid[newCoords] = v
			}
		}
	}
	l.grid = grid
	l.size = coordinates{tileSize.x * int(size), tileSize.y * int(size)}
	t.layout = &l
	return &t
}

func (ts tiles) formMatrix(curr *tile, pos coordinates, matrix map[coordinates]int, visited map[int]bool) {
	if visited[curr.id] {
		return
	}
	visited[curr.id] = true
	for _, m := range curr.matches {
		nextPos := coordinates{pos.x, pos.y}
		if m.dir == rt {
			nextPos.x++
		}
		if m.dir == lt {
			nextPos.x--
		}
		if m.dir == bot {
			nextPos.y++
		}
		if m.dir == top {
			nextPos.y--
		}
		matrix[nextPos] = m.id
		ts.formMatrix(ts[m.id], nextPos, matrix, visited)
	}
}
func (ts tiles) findTopLeftTile() *tile {
	for _, tl := range ts {
		if len(tl.matches) == 2 {
			if (tl.matches[0].dir == bot && tl.matches[1].dir == rt) || (tl.matches[0].dir == rt && tl.matches[1].dir == bot) {
				return tl
			}
		}
	}
	return nil
}

func (ts tiles) findBootstrapTile() *tile {
	for _, tl := range ts {
		if len(tl.matches) == 2 {
			if tl.matches[0].reverse && tl.matches[1].reverse {
				return tl
			}
		}
	}
	return nil
}

func (ts tiles) findBorders() {
	for _, t := range ts {
		t.layout.border()
	}
}

func (ts tiles) findMatches() {
	for _, t := range ts {
		t.findMatches(ts)
	}
}

func (t *tile) findMatches(ts tiles) {
	for _, tl := range ts {
		if tl.id == t.id {
			continue
		}

		for _, b := range t.layout.borders {
			for _, b1 := range tl.layout.borders {
				if b == b1 {
					t.matches = append(t.matches, match{id: tl.id})
					break
				}

				if b == reverse(b1) {
					t.matches = append(t.matches, match{id: tl.id, reverse: true})
					break
				}
			}
		}
	}
}

func alignTiles(t *tile, ts tiles, aligned map[int]bool) {
	if aligned[t.id] {
		return
	}
	aligned[t.id] = true
	for i, m := range t.matches {
		matchTile := ts[m.id]
		for _, op := range ops {
			if op != nil {
				op(matchTile)
			}
			if ok, dir := matchBorder(t, matchTile); ok {
				t.matches[i].dir = dir
				alignTiles(matchTile, ts, aligned)
				break
			}
		}
	}

}

func matchBorder(t1, t2 *tile) (bool, string) {
	if t1.layout.borders[rt] == t2.layout.borders[lt] {
		return true, rt
	}
	if t1.layout.borders[lt] == t2.layout.borders[rt] {
		return true, lt
	}
	if t1.layout.borders[top] == t2.layout.borders[bot] {
		return true, top
	}
	if t1.layout.borders[bot] == t2.layout.borders[top] {
		return true, bot
	}
	return false, ""
}

func rotateRight(t *tile) {
	t.layout.RotateRight()
}

func flip(t *tile) {
	t.layout.Flip()
}
func (l *layout) checkForSeaMonster(sml *layout) (map[coordinates]bool, bool) {
	smlCoords := make(map[coordinates]bool)
	for y := 0; y < l.size.y; y++ {
		for x := 0; x < l.size.x; x++ {
			if matches, ok := l.doCheck(coordinates{x, y}, sml); ok {
				smlCoords = mergeMaps(smlCoords, matches)
			}
		}
	}
	if len(smlCoords) > 0 {
		return smlCoords, true
	}
	return nil, false
}

func (l *layout) doCheck(pos coordinates, sml *layout) (map[coordinates]bool, bool) {
	smlCoords := make(map[coordinates]bool)
	for k, _ := range sml.grid {
		c := coordinates{pos.x + k.x, pos.y + k.y}
		if l.valueAt(c.x, c.y) != "#" {
			return nil, false
		}
		smlCoords[c] = true
	}
	return smlCoords, true
}

func (l *layout) onlyHave(val string) {
	for k, v := range l.grid {
		if v != val {
			delete(l.grid, k)
		}
	}
}

func mergeMaps(m1, m2 map[coordinates]bool) map[coordinates]bool {
	result := m1
	for k, v := range m2 {
		result[k] = v
	}

	return result

}

func print(matrix map[coordinates]int) {
	size := math.Sqrt(float64(len(matrix)))
	printOutput := ""
	for y := 0; y < int(size); y++ {
		for x := 0; x < int(size); x++ {
			printOutput += strconv.Itoa(matrix[coordinates{x, y}]) + "  "
		}

		printOutput += "\n"
	}
	fmt.Println(printOutput)
}
