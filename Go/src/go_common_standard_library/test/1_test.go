package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Add(a, b int) int {
	return a + b
}

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	expected := 5
	assert.Equal(t, expected, result, "Add(2, 3) result should be equal to expected")
}