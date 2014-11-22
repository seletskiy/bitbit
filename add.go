package main

import (
	"fmt"
)

type ProgramInstructionAdd struct {
	ProgramInstructionArgsToValue
}

func (op *ProgramInstructionAdd) Eval(state *ProgramState) {
	op.To.SetValue(state, op.To.GetValue(state)+op.Value.GetValue(state))
}

func (op *ProgramInstructionAdd) String() string {
	return fmt.Sprintf("ADD %s %s", op.To, op.Value)
}
