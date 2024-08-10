package calculator

import (
	"errors"
	"fmt"
	"math"
)

func Add(a, b float64) float64 {
	return a + b
}

func Substract(a, b float64) float64 {
	return a - b
}

func Multiply(a, b float64) float64 {
	return a * b
}

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero not allowed")
	}
	return a / b, nil
}

func Sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, fmt.Errorf("square root of negative number not allowed: %f", a)
	}
	return math.Sqrt(a), nil
}
