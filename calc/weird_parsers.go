package calc

type AddMulEqualLevelParser struct {
	NormalParser
}
type InvertedRuleParser struct {
	NormalParser
}

// factor: INTEGER
func (p *AddMulEqualLevelParser) factor() (AST, error) {
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

// expr: term ( (ADD | SUB | MUL | DIV) term)*
func (p *AddMulEqualLevelParser) Expr() (AST, error) {
	expr, err := p.factor()
	if err != nil {
		return GenericNode{}, err
	}

	for p.currentToken.kind == ADD || p.currentToken.kind == SUB || p.currentToken.kind == MUL || p.currentToken.kind == DIV {
		op := p.currentToken.value
		err := p.eat(p.currentToken.kind)
		if err != nil {
			return GenericNode{}, err
		}

		right, err := p.factor()
		if err != nil {
			return GenericNode{}, err
		}

		expr = BinaryOp{op: Op(op.(string)), left: expr, right: right}
	}

	return expr, nil
}

// factor: INTEGER
func (p *InvertedRuleParser) factor() (AST, error) {
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

// term: factor ( (ADD | SUB) factor )*
func (p *InvertedRuleParser) term() (AST, error) {
	term, err := p.factor()
	if err != nil {
		return GenericNode{}, err
	}

	for p.currentToken.kind == ADD || p.currentToken.kind == SUB {
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

// expr: term ( ( MUL | DIV) term)*
func (p *InvertedRuleParser) Expr() (AST, error) {
	expr, err := p.term()
	if err != nil {
		return GenericNode{}, err
	}

	for p.currentToken.kind == MUL || p.currentToken.kind == DIV {
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
