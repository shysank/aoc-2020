package calc

import "fmt"

type Visitor interface {
	VisitBinaryOp(b BinaryOp) interface{}
	VisitNum(n Num) interface{}
}

type AST interface {
	Accept(visitor Visitor) interface{}
}

type Num struct {
	val int
}

type BinaryOp struct {
	op    Op
	left  AST
	right AST
}

type GenericNode struct {
}

type Op string

func (binaryOp BinaryOp) Accept(visitor Visitor) interface{} {
	return visitor.VisitBinaryOp(binaryOp)
}

func (num Num) Accept(visitor Visitor) interface{} {
	return visitor.VisitNum(num)
}

func (node GenericNode) Accept(visitor Visitor) interface{} {
	fmt.Println("Do Nothing")
	return "This should not be printed ever"
}
