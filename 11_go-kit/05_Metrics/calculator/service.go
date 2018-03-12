package calculator

type Service interface {
	Plus(a int, b int) int
	Minus(a int, b int) int
	Multiply(a int, b int) int
	Divide(a int, b int) int
}

type Calculator struct {
}

func (Calculator) Plus(a, b int) int {
	return a + b
}

func (Calculator) Minus(a, b int) int {
	return a - b
}

func (Calculator) Multiply(a, b int) int {
	return a * b
}

func (Calculator) Divide(a, b int) int {
	return a / b
}

// create type that return function.
// this will be needed in main.go
type ServiceMiddleware func(Service) Service
