package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {
	data := []int{1, 2, 3, 7, 20, 22}
	result := BinarySearch(data, 7)

	assert.Equal(t, "Data found", result)
}

func TestNotFoundBinarySearch(t *testing.T) {
	data := []int{1, 2, 3, 7, 20, 22}
	result := BinarySearch(data, 11)

	assert.NotEqual(t, "Data found", result)
}

func TestLinearSearch(t *testing.T) {
	data := []string{"ruls", "raa", "kyoo"}
	result := LinearSearch(data, "raa")

	assert.Equal(t, "Data found", result)
}

func TestIntepolationSearch(t *testing.T) {
  data := []int{-10, 2, 6, -11, 44, 20, 12, 55}
  result := InterPolationSort(data, 44)

  assert.Equal(t, 4, result)
}
