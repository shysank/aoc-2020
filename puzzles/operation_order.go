package puzzles

import (
	"github.com/shysank/aoc-2020/calc"
	"io"
)

type OperationOrder struct {
}

func (o OperationOrder) Puzzle1(reader io.Reader) (Result, error) {
	expressions := ParseToStringArray(reader)

	sum := 0
	for _, expr := range expressions {
		result, err := evalExprWithAddMulEqual(expr)
		if err != nil {
			return nil, err
		}
		sum += result
	}
	return intResult(sum), nil
}

func (o OperationOrder) Puzzle2(reader io.Reader) (Result, error) {
	expressions := ParseToStringArray(reader)

	sum := 0
	for _, expr := range expressions {
		result, err := evalExprWithAddMulInverted(expr)
		if err != nil {
			return nil, err
		}
		sum += result
	}
	return intResult(sum), nil
}

func evalExprWithAddMulEqual(expr string) (int, error) {
	lexer := calc.NewLexer(expr)
	parser, err := calc.NewParser(&lexer)
	addMulEqualParser := calc.AddMulEqualLevelParser{parser}
	if err != nil {
		return 0, err
	}

	ast, err := addMulEqualParser.Expr()
	if err != nil {
		return 0, err
	}

	cal := calc.Calculator{}
	r, err := cal.Eval(ast)
	if err != nil {
		return 0, err
	}
	return r, nil
}

func evalExprWithAddMulInverted(expr string) (int, error) {
	lexer := calc.NewLexer(expr)
	parser, err := calc.NewParser(&lexer)
	addMulInvertedParser := calc.InvertedRuleParser{parser}
	if err != nil {
		return 0, err
	}

	ast, err := addMulInvertedParser.Expr()
	if err != nil {
		return 0, err
	}

	cal := calc.Calculator{}
	r, err := cal.Eval(ast)
	if err != nil {
		return 0, err
	}
	return r, nil
}
