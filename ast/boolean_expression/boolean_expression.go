package boolean_expression

type BooleanExpression interface {
	Interpret() bool
	Stringify() string
}