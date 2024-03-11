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
