package calculator

import "golang.org/x/net/context"

type Service interface {
	Plus(ctx context.Context, a int32, b int32) (int32, error)
	Minus(ctx context.Context, a int32, b int32) (int32, error)
	Multiply(ctx context.Context, a int32, b int32) (int32, error)
	Divide(ctx context.Context, a int32, b int32) (int32, error)
}

type CalculatorService struct {
}

func (CalculatorService) Plus(ctx context.Context, a int32, b int32) (int32, error) {
	return a + b, nil
}
func (CalculatorService) Minus(ctx context.Context, a int32, b int32) (int32, error) {
	return a - b, nil
}
func (CalculatorService) Multiply(ctx context.Context, a int32, b int32) (int32, error) {
	return a * b, nil
}
func (CalculatorService) Divide(ctx context.Context, a int32, b int32) (int32, error) {
	return a / b, nil
}
