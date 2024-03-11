package search

func LinearSearch(data []string, item string) string {
  for _, dt := range data {
    if dt == item {
      return "Data found"
    }
  }

  return "Not Found"
}
