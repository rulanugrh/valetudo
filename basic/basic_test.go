package basic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.M) {
  t.Run()
}

func TestFactorial(t *testing.T) {
  num := Factorial(5)

  assert.Equal(t, 120, num)
}

func TestRecursive(t *testing.T) {
  num := Recursive(7)

  assert.Equal(t, 112, num)
}
