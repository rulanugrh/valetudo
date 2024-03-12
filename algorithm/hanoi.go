package algorithm

import "fmt"

type Hanoi interface {
	PlayHanoi(n int)
}

type Towers struct{}

func (t *Towers) PlayHanoi(n int) {
  t.MoveN(n, 1, 2, 3)
}

func (t *Towers) MoveN(n, from, to, via int) {
	if n > 0 {
		t.MoveN(n-1, from, via, to)
    t.MoveM(from, to)
    t.MoveN(n-1, via, to, from)
	}
}

func (t *Towers) MoveM(from, to int) {
	fmt.Println("Move from rod", from, "to rod", to)
}
