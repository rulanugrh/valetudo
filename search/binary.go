package search

func BinarySearch(data []int, item int) string {
  var low = 0

  var high = len(data) - 1

  for low <= high {
    var mid = low + high
    var guess = data[mid]

    if guess == item {
      return "Data found"
    } else if guess > item {
      high = mid - 1
    } else {
      low = mid + 1
    }
  }

  return "Data not found"
}
