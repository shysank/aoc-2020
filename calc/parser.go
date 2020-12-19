package calc

import (
	"fmt"
)

type IParser interface {
	Expr() (AST, error)
}

type NormalParser struct {
	lexer        *Lexer
	currentToken Token
}

func NewParser(l *Lexer) (NormalParser, error) {
	token, err := l.NextToken()
	if err != nil {
		return NormalParser{}, err
	}
	return NormalParser{lexer: l, currentToken: token}, nil
}

/*

grammar

expr: term ( (ADD | SUB) term)*
term: factor ( (MUL | DIV) factor )*
factor: INTEGER | ( LPAREN expr RPAREN )

*/

// factor: INTEGER
func (p *NormalParser) factor() (AST, error) {
	if p.currentToken.kind == INTEGER {
		intVal, _ := p.currentToken.value.(int)
		err := p.eat(INTEGER)
		if err != nil {
			return GenericNode{}, err
		}
		return Num{intVal}, nil
	} else {
		err := p.eat(LPAREN)
		if err != nil {
			return GenericNode{}, err
		}

		exprVal, err := p.Expr()
		if err != nil {
			return GenericNode{}, err
		}

		err = p.eat(RPAREN)
		if err != nil {
			return GenericNode{}, err
		}

		return exprVal, nil
	}
}

// term: factor ( (MUL | DIV) factor )*
func (p *NormalParser) term() (AST, error) {
	term, err := p.factor()
	if err != nil {
		return GenericNode{}, err
	}

	for p.currentToken.kind == MUL || p.currentToken.kind == DIV {
		op := p.currentToken.value
		err := p.eat(p.currentToken.kind)
		if err != nil {
			return GenericNode{}, err
		}

		right, err := p.factor()
		if err != nil {
			return GenericNode{}, err
		}

		term = BinaryOp{op: Op(op.(string)), left: term, right: right}
	}

	return term, nil
}

// expr: term ( (ADD | SUB) term)*
func (p *NormalParser) Expr() (AST, error) {
	expr, err := p.term()
	if err != nil {
		return GenericNode{}, err
	}

	for p.currentToken.kind == ADD || p.currentToken.kind == SUB {
		op := p.currentToken.value
		err := p.eat(p.currentToken.kind)
		if err != nil {
			return GenericNode{}, err
		}

		right, err := p.term()
		if err != nil {
			return GenericNode{}, err
		}

		expr = BinaryOp{op: Op(op.(string)), left: expr, right: right}
	}

	return expr, nil
}

func (p *NormalParser) eat(tokenType TokenType) error {
	if p.currentToken.kind == tokenType {
		token, err := p.lexer.NextToken()
		if err != nil {
			return err
		}
		p.currentToken = token
		return nil
	}

	return fmt.Errorf("Syntax Error: invalid syntax: `%s` , Expected `%s`", p.currentToken.kind, tokenType)
}
