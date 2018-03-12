package calculator

type Service interface {
	plus(a int, b int) int
	minus(a int, b int) int
	multiple(a int, b int) int
	divide(a int, b int) int
}

type CalculatorService struct {
}

func (CalculatorService) plus(a int, b int) int {
	return a + b
}
func (CalculatorService) minus(a int, b int) int {
	return a - b
}
func (CalculatorService) multiple(a int, b int) int {
	return a * b
}
func (CalculatorService) divide(a int, b int) int {
	return a / b
}
