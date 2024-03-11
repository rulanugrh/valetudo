package main

import (
	"os"

	"github.com/rulanugrh/veletudo/tree"
)

func main() {
	t := &tree.BinaryTree{}
  t.InsertData(100).InsertData(-20).InsertData(-50).InsertData(-15).InsertData(-60).InsertData(50).InsertData(60).InsertData(55).InsertData(75).InsertData(15).InsertData(5).InsertData(-10)

  tree.Printout(os.Stdout, t.Root, 0, 'M')
}
