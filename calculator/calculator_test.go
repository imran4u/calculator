package calculator

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAdd(t *testing.T) {
	calc := Calculator{}
	result := calc.Add(2, 3)
	require.Equal(t, result, 5)
	require.NotEmpty(t, result)

}

func TestSubtract(t *testing.T) {
	calc := Calculator{}
	result := calc.Subtract(2, 3)
	require.Equal(t, -1, result)
	require.NotEmpty(t, result)

}
