package data

import (
	"strings"
)

// stack style LIFO
type Stack struct {
  Elems []string
}

func (s *Stack) PushStack(data string) {
  if data == "" {
    println("data null")
  }

  s.Elems = append(s.Elems, data)
}

func (s *Stack) TopStack() []string {
  for len(s.Elems) > 0 {
    n := len(s.Elems) - 1
    s.Elems = s.Elems[:n]

    return s.Elems

  }

  return nil
}

func (s *Stack) PopStack(index int) {
  s.Elems = append(s.Elems[:index], s.Elems[index+1:]...)
}

func (s *Stack) JoinStack(sep string) string {
  return strings.Join(s.Elems, sep)
}
