package main

import (
	"fmt"
	"github.com/shysank/aoc-2020/calc"
	"github.com/shysank/aoc-2020/puzzles"
	"math/big"
	"os"
)

func main() {
	gridTest()
}

func bitsTest() {
	var bits = big.NewInt(int64(11))
	bits.SetBit(bits, 1, 0)
	bits.SetBit(bits, 6, 1)
	fmt.Println(bits)

}

func crtTest() {
	var eqs = puzzles.CRTEquations{puzzles.NewEquation(0, 17), puzzles.NewEquation(11, 13),
		puzzles.NewEquation(16, 19)}
	x := eqs.Solve()
	fmt.Println(x)
}

func waysTest() {
	var arr = []int{12, 13, 15, 16, 17}
	w := ways(arr, len(arr)-1)

	fmt.Println(w)
}

func ways(arr []int, index int) int {
	if index == 0 {
		return 1
	}

	curr := arr[index]
	prevRoutes := ways(arr, index-1)
	var prev1, prev2 int
	if index-2 >= 0 && curr-arr[index-2] <= 3 {
		prev1 = ways(arr, index-2)
	}

	if index-3 >= 0 && curr-arr[index-3] <= 3 {
		prev2 = ways(arr, index-3)
	}

	return prevRoutes + prev1 + prev2

}

func calcTest() {
	text := "(8 * 3 + 9 + 3 * 4 * 3)"
	lexer := calc.NewLexer(text)
	parser, err := calc.NewParser(&lexer)
	addMulInvertedParser := calc.InvertedRuleParser{parser}
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	ast, err := addMulInvertedParser.Expr()
	if err != nil {
		fmt.Println(err.Error())
	}
	cal := calc.Calculator{}
	r, _ := cal.Eval(ast)
	fmt.Println(r)
}

func gridTest() {
	grid := `#....####.
#..#.##...
#.##..#...
######.#.#
.#...#.#.#
.#########
.###.#..#.
########.#
##...##.#.
..###.#.#.
`
	l := puzzles.ParseLayoutFromString(grid)
	l.Print()
	l.RotateLeft()
	l.RotateLeft()
	l.RotateLeft()
	l.Flip()
	l.Print()
}
