package tree

import (
	"fmt"
	"io"
)

type BinaryNode struct {
	left  *BinaryNode
	right *BinaryNode
	data  int64
}

type BinaryTree struct {
	Root *BinaryNode
}

func (t *BinaryTree) InsertData(data int64) *BinaryTree {
	if t.Root == nil {
		t.Root = &BinaryNode{data: data, left: nil, right: nil}
	} else {
		t.Root.InsertNode(data)
	}

	return t
}

func (n *BinaryNode) InsertNode(data int64) {
	if n == nil {
		return
	} else if data <= n.data {
		if n.left == nil {
			n.left = &BinaryNode{data: data, left: nil, right: nil}
		} else {
			n.left.InsertNode(data)
		}
	} else {
		if n.right == nil {
			n.right = &BinaryNode{data: data, left: nil, right: nil}
		} else {
			n.right.InsertNode(data)
		}
	}
}

func Printout(w io.Writer, node *BinaryNode, ns int, ch rune) {
  if node == nil {
    return
  }

  for i := 0; i < ns; i++ {
    fmt.Fprintf(w, " ")
  }

  fmt.Fprintf(w, "%c:%v\n", ch, node.data)
  Printout(w, node.left, ns+2, 'L')
  Printout(w, node.right, ns+2, 'R')
}
