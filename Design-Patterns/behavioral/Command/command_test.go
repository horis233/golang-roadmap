package Command

import "testing"

func TestCommand_Execute(t *testing.T) {
	A := NewPerson("A", NewCommand(nil, nil))
	B := NewPerson("B", NewCommand(&A, A.Listen))
	C := NewPerson("C", NewCommand(&B, B.Buy))
	D := NewPerson("D", NewCommand(&C, C.Cook))
	E := NewPerson("E", NewCommand(&D, D.Wash))

	E.Talk()

}

