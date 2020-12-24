package puzzles

import (
	"io"
)

type LobbyLayout struct {
	days int
}

type tilePath struct {
	dirs []string
}

const (
	dirEast      = "e"
	dirWest      = "w"
	dirNorthEast = "ne"
	dirNorthWest = "nw"
	dirSouthEast = "se"
	dirSouthWest = "sw"

	black = "black"
	white = "white"
)

var base = coordinates{20, 12}
var dirMovement = map[string]coordinates{
	dirEast:      {base.x, 0},
	dirWest:      {-1 * base.x, 0},
	dirNorthEast: {base.x / 2, 3 * base.y / 2},
	dirNorthWest: {-1 * base.x / 2, 3 * base.y / 2},
	dirSouthEast: {base.x / 2, -3 * base.y / 2},
	dirSouthWest: {-1 * base.x / 2, -3 * base.y / 2},
}

func (l LobbyLayout) Puzzle1(reader io.Reader) (Result, error) {
	tiles := ParseTileDirections(reader)
	var tileColor = determineInitialColors(tiles)
	blackTileCount := calcBlackTileCount(tileColor)
	return intResult(blackTileCount), nil
}

func (l LobbyLayout) Puzzle2(reader io.Reader) (Result, error) {
	tiles := ParseTileDirections(reader)
	var tileColor = determineInitialColors(tiles)

	for i := 0; i < l.days; i++ {
		newTileColor := copyTileColor(tileColor)
		expanded := make(map[coordinates]string)
		for k, _ := range tileColor {
			applyTileColorRules(k, tileColor, newTileColor, expanded)
		}

		for k, _ := range expanded {
			applyTileColorRules(k, tileColor, newTileColor, map[coordinates]string{})
		}
		tileColor = newTileColor
	}
	blackTileCount := calcBlackTileCount(tileColor)
	return intResult(blackTileCount), nil
}

func applyTileColorRules(c coordinates, old, new, expanded map[coordinates]string) {
	adjBlackTileCount := 0
	var adjCoords []coordinates
	for _, v := range dirMovement {
		adjCoords = append(adjCoords, coordinates{c.x + v.x, c.y + v.y})
	}

	for _, adj := range adjCoords {

		if _, ok := old[adj]; !ok {
			expanded[adj] = white
			continue
		}

		if old[adj] == black {
			adjBlackTileCount++
		}
	}
	if old[c] == black {
		if adjBlackTileCount == 0 || adjBlackTileCount > 2 {
			new[c] = white
		}
	} else {
		if adjBlackTileCount == 2 {
			new[c] = black
		} else {
			new[c] = white
		}
	}

}

func determineInitialColors(tiles []tilePath) map[coordinates]string {
	tileColor := make(map[coordinates]string)
	for _, t := range tiles {

		pos := coordinates{0, 0}
		for _, d := range t.dirs {
			move, _ := dirMovement[d]
			pos = coordinates{pos.x + move.x, pos.y + move.y}
		}

		color, ok := tileColor[pos]
		if !ok {
			tileColor[pos] = black
		} else {
			if color == black {
				tileColor[pos] = white
			} else {
				tileColor[pos] = black
			}
		}
	}
	return tileColor
}

func calcBlackTileCount(tileColor map[coordinates]string) int {
	blackTileCount := 0
	for _, v := range tileColor {
		if v == black {
			blackTileCount++
		}
	}

	return blackTileCount
}

func copyTileColor(src map[coordinates]string) map[coordinates]string {
	dst := make(map[coordinates]string)
	for k, v := range src {
		dst[coordinates{k.x, k.y}] = v
	}
	return dst
}
