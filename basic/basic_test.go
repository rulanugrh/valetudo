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

func TestFizzBuzz(t *testing.T) {
  a := FizzBuzz(5)
  b := FizzBuzz(3)
  c := FizzBuzz(15)
  d := FizzBuzz(19)

  assert.Equal(t, "Fizz", b)
  assert.Equal(t, "Buzz", a)
  assert.Equal(t, "FizzBuzz", c)
  assert.Equal(t, 19, d)
}
