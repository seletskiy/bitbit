package main

import (
	"fmt"
	"math"
)

type ProgramInstructionDiv struct {
	ProgramInstructionArgsToValue
}

func (op *ProgramInstructionDiv) Eval(state *ProgramState) {
	a := op.To.GetValue(state)
	b := op.Value.GetValue(state)
	if math.Abs(float64(b)) < 1e-12 {
		panic("float64 divide by zero")
	}
	op.To.SetValue(state, a/b)
}

func (op *ProgramInstructionDiv) String() string {
	return fmt.Sprintf("DIV %s %s", op.To, op.Value)
}
