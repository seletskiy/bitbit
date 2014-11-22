package main

import (
	"fmt"
)

type ProgramInstructionFun struct {
	ProgramInstructionArgsValue
}

func (op *ProgramInstructionFun) Eval(state *ProgramState) {
	state.ExternalData.(*DataStorage).FunValue = float64(op.Value.GetValue(state))
}

func (op *ProgramInstructionFun) String() string {
	return fmt.Sprintf("FUN %s", op.Value)
}
