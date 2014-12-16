package main

import "fmt"

type ProgramInstructionZero struct {
	Out ProgramArgRegister
}

func (zero *ProgramInstructionZero) Eval(state *ProgramState) error {
	zero.Out.SetValue(state, FloatValue(0))
	return nil
}

func (zero *ProgramInstructionZero) String() string {
	return fmt.Sprintf("ZERO %s", zero.Out)
}

func (zero *ProgramInstructionZero) GetArgsCount() int {
	return 1
}

func (zero *ProgramInstructionZero) GetArg(index int) ProgramInstructionArg {
	switch index {
	case 0:
		return zero.Out
	}

	return nil
}

func (zero *ProgramInstructionZero) SetArg(
	index int, arg ProgramInstructionArg,
) {
	switch index {
	case 0:
		zero.Out = arg.(ProgramArgRegister)
	}
}

func (zero *ProgramInstructionZero) Init() {
	zero.Out = FloatRegister(0)
}

func (zero *ProgramInstructionZero) Copy() ProgramInstruction {
	instructionCopy := *zero
	return &instructionCopy
}
