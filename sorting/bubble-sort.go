package sorting

func BubbleSort(arr []int) []int {
	var isDone = false

	for !isDone {
		isDone = true
		var i = 0

		for i < len(arr)-1 {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				isDone = false
			}

			i++
		}
	}

	return arr
}
