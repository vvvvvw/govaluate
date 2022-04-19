package govaluate

/*
	Represents a single parsed token.
*/
type ExpressionToken struct {
	Kind  TokenKind   //token的类型
	Value interface{} //token的值
}
