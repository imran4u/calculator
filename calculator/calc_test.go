package calculator

import (
    "testing"
)

func TestAdd(t *testing.T) {
    c := Calculator{}
    result := c.Add(3, 2)
    expected := 5
    if result != expected {
        t.Errorf("Add(3, 2) = %d; want %d", result, expected)
    }
}

func TestSum(t *testing.T) {
    c := Calculator{}
    result := c.sum(1, 2, 3, 4, 5)
    expected := 15
    if result != expected {
        t.Errorf("sum(1, 2, 3, 4, 5) = %d; want %d", result, expected)
    }
}

func TestSubtract(t *testing.T) {
    c := Calculator{}
    result := c.Subtract(5, 3)
    expected := 2
    if result != expected {
        t.Errorf("Subtract(5, 3) = %d; want %d", result, expected)
    }
}

func TestMultiply(t *testing.T) {
    c := Calculator{}
    result := c.Multiply(2, 3)
    expected := 6
    if result != expected {
        t.Errorf("Multiply(2, 3) = %d; want %d", result, expected)
    }
}

func TestDiv(t *testing.T) {
    c := Calculator{}

    t.Run("Division by non-zero", func(t *testing.T) {
        result, err := c.Div(6, 2)
        expected := 3
        if err != nil {
            t.Errorf("Div(6, 2) returned error: %v", err)
        }
        if result != expected {
            t.Errorf("Div(6, 2) = %d; want %d", result, expected)
        }
    })

    t.Run("Division by zero", func(t *testing.T) {
        _, err := c.Div(6, 0)
        if err == nil {
            t.Error("Div(6, 0) did not return expected error")
        }
    })
}
