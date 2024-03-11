package sorting

func InsertionSort(arr []int) []int {
	var i = 1
	for i < len(arr) {
		var j = i
		for j >= 1 && arr[j] < arr[j-1] {
			arr[j], arr[j-1] = arr[j-1], arr[j]

			j--
		}
		i++
	}

	return arr
}
