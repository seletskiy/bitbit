package main

import (
	"fmt"
)

type ProgramInstructionMov struct {
	ProgramInstructionArgsToValue
}

func (op *ProgramInstructionMov) Eval(state *ProgramState) {
	op.To.SetValue(state, op.Value.GetValue(state))
}

func (op *ProgramInstructionMov) String() string {
	return fmt.Sprintf("MOV %s %s", op.To, op.Value)
}
