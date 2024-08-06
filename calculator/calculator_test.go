package calculator_test

import (
	"calculator"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Parallel()
	var want float64 = 4
	got := calculator.Add(2, 2)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}

func TestSubtract(t *testing.T) { //testing function should start with 'Test' and accept a parameter of type '*testing/T', not have any type of output type
	t.Parallel() //run concurrently with other tests
	var want float64 = 2
	got := calculator.Substract(4, 2)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}
