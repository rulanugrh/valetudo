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
