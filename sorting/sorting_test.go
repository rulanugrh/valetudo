package sorting

import (
	"testing"

  "github.com/stretchr/testify/assert"
)

var arr = []int{1, 4, 5, 7, 2, 10, 11}

func TestMain(m *testing.M) {
  println("runnint unit test")

  m.Run()

  println("test success")
}

func TestBubbleSort(t *testing.T) {
  valid := []int{1, 2, 4, 5, 7, 10, 11}
  result := BubbleSort(arr)

  assert.Equal(t, valid, result)
}

func TestNotValidBubbleSort(t *testing.T) {
  notEqual := []int{1, 5, 2, 4, 10, 11}
  result := BubbleSort(arr)

  assert.NotEqual(t, notEqual, result)
}

func TestSelectionSort(t *testing.T) {
  valid := []int{1, 2, 4, 5, 7, 10, 11}
  result := SelectionSort(arr)

  if !assert.Equal(t, valid, result) {
    println("selection sort not valid")
  }

  print("success test selection sort")

}

func TestInsertionSort(t *testing.T) {
  valid := []int{1, 2, 4, 5, 7, 10, 11}
  result := InsertionSort(arr)

  if !assert.Equal(t, valid, result) {
    println("insertion sort not valid")
  }

  println("success test insertion sort")
}
