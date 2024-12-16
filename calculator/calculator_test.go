package calculator

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAdd(t *testing.T) {
	calc := Calculator{}
	result := calc.Add(2, 4)
	require.Equal(t, 6, result)
	require.NotEmpty(t, result)

}

func TestSum(t *testing.T) {
	calc := Calculator{}
	result := calc.sum(2, 3)
	require.Equal(t, 5, result)
	require.NotEmpty(t, result)

	result = calc.sum(2, 3, 4, 5, 6, 9)
	require.Equal(t, 29, result)
	require.NotEmpty(t, result)

}

func TestSubtract(t *testing.T) {
	calc := Calculator{}
	result := calc.Subtract(2, 3)
	require.Equal(t, -1, result)
	require.NotEmpty(t, result)

}

func TestMultiply(t *testing.T) {
	calc := Calculator{}
	result := calc.Multiply(2, 3)
	require.Equal(t, 6, result)
	result = calc.Multiply(0, 10)
	require.Empty(t, result)
	require.Equal(t, 0, result)

}

func TestDivistion(t *testing.T) {
	calc := Calculator{}
	result, err := calc.Div(6, 3)
	require.NoError(t, err)
	require.Equal(t, 2, result)

	result, err = calc.Div(0, 10)
	require.NoError(t, err)
	require.Empty(t, result)
	require.Equal(t, 0, result)

	result, err = calc.Div(100, 0)
	require.Error(t, err)
	require.Equal(t, "denominator should not be zero", err.Error())
	require.Empty(t, result)

}
