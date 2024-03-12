package algorithm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompressLZW(t *testing.T) {
	str := CompressLZW("Australia")
	expected := []int{65, 117, 115, 116, 114, 97, 108, 105, 97}

	assert.Equal(t, expected, str)
}

func TestKMP(t *testing.T) {
  str := SearchString("syahara", "sy")

  assert.Equal(t, 0, str)
}

func TestKMPLast(t *testing.T) {
  str := SerchNext("syahara", "ha")
  assert.Equal(t, 3, str)
}

func TestKMPNone(t *testing.T) {
  str := SearchString("araa", "rul")
  assert.Equal(t, -1, str)
}
