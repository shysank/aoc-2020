package calc

import "fmt"

type Interpreter interface {
	Eval(ast AST) interface{}
}

type Calculator struct {
}

func (c Calculator) Eval(ast AST) (int, error) {
	return ast.Accept(c).(int), nil
}

func (c Calculator) VisitBinaryOp(b BinaryOp) interface{} {
	switch b.op {
	case "+":
		return b.left.Accept(c).(int) + b.right.Accept(c).(int)
	case "-":
		return b.left.Accept(c).(int) - b.right.Accept(c).(int)
	case "*":
		return b.left.Accept(c).(int) * b.right.Accept(c).(int)
	case "/":
		return b.left.Accept(c).(int) / b.right.Accept(c).(int)
	}
	panic("Unknown Op")
}

func (c Calculator) VisitNum(num Num) interface{} {
	return num.val
}

type ReversePolish struct {
}

func (r ReversePolish) Eval(ast AST) (string, error) {
	return ast.Accept(r).(string), nil
}

func (r ReversePolish) VisitBinaryOp(b BinaryOp) interface{} {
	switch b.op {
	case "+":
		return b.left.Accept(r).(string) + " " + b.right.Accept(r).(string) + " " + "+" + " "
	case "-":
		return b.left.Accept(r).(string) + " " + b.right.Accept(r).(string) + " " + "-" + " "
	case "*":
		return b.left.Accept(r).(string) + " " + b.right.Accept(r).(string) + " " + "*" + " "
	case "/":
		return b.left.Accept(r).(string) + " " + b.right.Accept(r).(string) + " " + "/" + " "
	}
	panic("Unknown Op")
}

func (r ReversePolish) VisitNum(num Num) interface{} {
	return fmt.Sprintf("%d", num.val)
}

type Lisp struct {
}

func (r Lisp) Eval(ast AST) (string, error) {
	return ast.Accept(r).(string), nil
}

func (r Lisp) VisitBinaryOp(b BinaryOp) interface{} {
	switch b.op {
	case "+":
		return "(" + "+" + " " + b.left.Accept(r).(string) + " " + b.right.Accept(r).(string) + ")"
	case "-":
		return "(" + "-" + " " + b.left.Accept(r).(string) + " " + b.right.Accept(r).(string) + ")"
	case "*":
		return "(" + "*" + " " + b.left.Accept(r).(string) + " " + b.right.Accept(r).(string) + ")"
	case "/":
		return "(" + "/" + " " + b.left.Accept(r).(string) + " " + b.right.Accept(r).(string) + ")"
	}
	panic("Unknown Op")
}

func (l Lisp) VisitNum(num Num) interface{} {
	return fmt.Sprintf("%d", num.val)
}
