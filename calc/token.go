package calc

type Token struct {
	kind  TokenType
	value TokenVal
}

type TokenVal interface{}
type TokenType string

const (
	INTEGER TokenType = "Integer"
	ADD     TokenType = "Plus"
	SUB     TokenType = "Minus"
	MUL     TokenType = "Multiply"
	DIV     TokenType = "Divide"
	EOF     TokenType = "End_Of_File"
	LPAREN  TokenType = "Left_Parentheses"
	RPAREN  TokenType = "Right_Parentheses"
)

func NewToken(kind TokenType, val interface{}) Token {
	return Token{kind, TokenVal(val)}
}

func Eof() Token {
	return Token{EOF, nil}
}

func Plus() Token {
	return Token{ADD, "+"}
}

func Minus() Token {
	return Token{SUB, "-"}
}

func Mul() Token {
	return Token{MUL, "*"}
}

func Div() Token {
	return Token{DIV, "/"}
}

func LParen() Token {
	return Token{LPAREN, "("}
}

func RParen() Token {
	return Token{RPAREN, ")"}
}
