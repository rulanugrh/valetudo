package sorting

func CountingSort(arr []int) []int {
	var max = arr[0]

	var i = 1
	for i < len(arr) {
		if arr[i] > max {
			max = arr[i]
		}

		i++
	}

	var indix = make([]int, max+1)
	var j = 0

	for j < len(arr) {
		indix[arr[j]]++
		j++
	}

	var k = 1
	for k < len(indix) {
		indix[k] += indix[k-1]
		k++
	}

	var result = make([]int, len(arr))

	var m = 0
	for m < len(arr) {
		result[indix[arr[m]]-1] = arr[m]
		indix[arr[m]]--

		m++
	}

	return result
}
