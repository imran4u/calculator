package calculator

import "fmt"

type Calculator struct{}

func (c *Calculator) Add(a, b int) int {
	return a + b
}

func (c *Calculator) sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func (c *Calculator) Subtract(a, b int) int {
	return a - b
}

func (c *Calculator) Multiply(a, b int) int {
	return a * b
}

func (c *Calculator) Div(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("denominator should not be zero")
	}
	return a / b, nil
}
