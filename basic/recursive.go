package basic

func Recursive(n int) int {
  if n < 100 {
    return Recursive(n + n)
  } else if n >= 100 {
    return n
  } else {
    return 0
  }
}
