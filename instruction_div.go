package main

import (
	"errors"
	"fmt"
	"math"
)

type ProgramInstructionDiv struct {
	Out ProgramArgRegister
	In  ProgramArgReference
}

func (div *ProgramInstructionDiv) Eval(state *ProgramState) error {
	a := div.Out.GetValue(state).GetFloat64()
	b := div.In.GetValue(state).GetFloat64()
	if math.Abs(float64(b)) < 1e-12 {
		return errors.New("float64 divide by zero")
	}

	div.Out.SetValue(state, FloatValue(a/b))

	return nil
}

func (div *ProgramInstructionDiv) String() string {
	return fmt.Sprintf("DIV %s %s", div.Out, div.In)
}

func (div *ProgramInstructionDiv) GetArgsCount() int {
	return 2
}

func (div *ProgramInstructionDiv) GetArg(index int) ProgramInstructionArg {
	switch index {
	case 0:
		return div.Out
	case 1:
		return div.In
	}

	return nil
}

func (div *ProgramInstructionDiv) SetArg(
	index int, arg ProgramInstructionArg,
) {
	switch index {
	case 0:
		div.Out = arg.(ProgramArgRegister)
	case 1:
		div.In = arg.(ProgramArgReference)
	}
}

func (div *ProgramInstructionDiv) Init() {
	div.Out = FloatRegister(0)
	div.In = FloatValue(0)
}

func (div *ProgramInstructionDiv) Copy() ProgramInstruction {
	instructionCopy := *div
	return &instructionCopy
}
