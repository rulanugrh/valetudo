package sorting_test

import (
	"testing"

	"github.com/rulanugrh/veletudo/sorting"
	"github.com/stretchr/testify/assert"
)

var arr = []int{1, 4, 5, 7, 2, 10, 11}
var arr32 = []int32{200, 55, -12, 58, 100}

func TestMain(m *testing.M) {
	println("runnint unit test")

	m.Run()

	println("test success")
}

func TestBubbleSort(t *testing.T) {
	valid := []int{1, 2, 4, 5, 7, 10, 11}
	result := sorting.BubbleSort(arr)

	assert.Equal(t, valid, result)
}

func TestNotValidBubbleSort(t *testing.T) {
	notEqual := []int{1, 5, 2, 4, 10, 11}
	result := sorting.BubbleSort(arr)

	assert.NotEqual(t, notEqual, result)
}

func TestSelectionSort(t *testing.T) {
	valid := []int{1, 2, 4, 5, 7, 10, 11}
	result := sorting.SelectionSort(arr)

	if !assert.Equal(t, valid, result) {
		println("selection sort not valid")
	}

	print("success test selection sort")

}

func TestInsertionSort(t *testing.T) {
	valid := []int{1, 2, 4, 5, 7, 10, 11}
	result := sorting.InsertionSort(arr)

	if !assert.Equal(t, valid, result) {
		println("insertion sort not valid")
	}

	println("success test insertion sort")
}

func TestMergeSort(t *testing.T) {
	valid := []int{1, 2, 4, 5, 7, 10, 11}
	result := sorting.MergeSort(arr)

	if !assert.Equal(t, valid, result) {
		println("merge sort not valid")
	}

	println("success test merge sort")
}

func TestCountingSort(t *testing.T) {
	valid := []int{1, 2, 4, 5, 7, 10, 11}
	result := sorting.CountingSort(arr)

	if !assert.Equal(t, valid, result) {
		println("counting sort not valid")
	}

	println("success test counting sort")
}

func TestRadixSort(t *testing.T) {
	valid := []int32{-12, 55, 58, 100, 200}
	result := sorting.RadixSort(arr32)

	if !assert.Equal(t, valid, result) {
		println("radix sort invalid test")
	}

	println("success test radix sort")
}
