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

func TestInsertStack(t *testing.T) {
  arr := Stack{}
  arr.PushStack("ruls")
  arr.PushStack("kyoo")

  value := arr.JoinStack(", ")
  assert.Equal(t, "ruls, kyoo", value)
}

func TestDeleteStack(t *testing.T) {
  arr := Stack{}
  arr.PushStack("ruls")
  arr.PushStack("kyoo")
  arr.PushStack("ryoo")

  value := arr.TopStack()
  asserts := []string{"ruls", "kyoo"}

  assert.Equal(t, asserts, value)
}

func TestDeleteByIndexStack(t *testing.T) {
  arr := Stack{}
  arr.PushStack("yooo")
  arr.PushStack("kyoo")
  arr.PushStack("ryoo")

  arr.PopStack(1)
  value := arr.JoinStack(", ")
  assert.Equal(t, "yooo, ryoo", value)
}
