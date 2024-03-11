package sorting

func SelectionSort(arr []int) []int {
	var i = 1

	for i < len(arr)-1 {
		var j = i + 1
		var minIndex = i

		if j < len(arr) {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
			j++
		}

		if minIndex != i {
			var temp = arr[i]
			arr[i] = arr[minIndex]
			arr[minIndex] = temp
		}

		i++
	}

	return arr
}
