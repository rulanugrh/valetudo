package search

// Interpolation sort ialah metode searching
// yang improvisasi dari Binary Search, berbeda dengan binary search
// yang mencari nilai ke nilai mid nya
// interpolation mencari kemana saja hingga nilai nya sama dengan keynya

func InterPolationSort(array []int, key int) int {
  min, max := array[0], array[len(array) - 1]
  low, high := 0, len(array)-1

  for {
    if key < min {
      return low
    }

    if key > max {
      return high + 1
    }

    var guess int
    if high == low {
      guess = high
    } else {
      size := high - low
      offset := int(float64(size - 1) * (float64(key-min) / float64(max-min)))
      guess = low + offset
    }

    if array[guess] == key {
      for guess > 0 && array[guess - 1] == key {
        guess--
      }

      return guess
    }

    if array[guess] > key {
      high = guess - 1
      max = array[high]
    } else {
      low = guess + 1
      min = array[low]
    }
  }
}
