package calculator

import (
	"testing"
)

func TestAdd(t *testing.T) {
	calc := Calculator{}
	result := calc.Add(2, 3)
	if result != 5 {
		t.Errorf("Expected 5, but got %d", result)
	}

}
