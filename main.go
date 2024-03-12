package main

import (
	"fmt"
	"os"

	"github.com/rulanugrh/veletudo/algorithm"
	"github.com/rulanugrh/veletudo/tree"
)

func main() {
	t := &tree.BinaryTree{}
	t.InsertData(100).InsertData(-20).InsertData(-50).InsertData(-15).InsertData(-60).InsertData(50).InsertData(60).InsertData(55).InsertData(75).InsertData(15).InsertData(5).InsertData(-10)

	tree.Printout(os.Stdout, t.Root, 0, 'M')

	gra := [][]algorithm.Graph{
		1: {{2, 3}, {3, 8}, {5, -4}},
		2: {{4, 1}, {5, 7}},
		3: {{2, 4}},
		4: {{1, 2}, {3, -5}},
		5: {{4, 6}},
	}

	dist := algorithm.FloydWarshall(gra)
	for _, d := range dist {
		fmt.Printf("%4g\n", d)
	}

  var ts algorithm.Hanoi
  ts = new(algorithm.Towers)
  ts.PlayHanoi(3)
}
