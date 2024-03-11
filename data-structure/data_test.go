package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.M) {
  println("running test data structure")

  t.Run()

  println("running test done")
}

func TestArray(t *testing.T) {
  arr := Array{}
  arr.InsertArray("ruls")
  arr.InsertArray("kyoo")

  value := arr.JoinArray(", ")
  assert.Equal(t, "ruls, kyoo", value)
}

func TestDeleteArray(t *testing.T) {
  arr := Array{}
  arr.InsertArray("ruls")
  arr.InsertArray("kyoo")

  value := arr.DeletedArray(1)
  assert.Equal(t, "success deleted", value)
}
