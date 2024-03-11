package basic

func Factorial(n int) int {
  if n == 1 {
    return 1
  } else {
    return n * Factorial(n - 1)
  }
}
