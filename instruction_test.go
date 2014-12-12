package main

import "fmt"

type TestInstruction struct {
	In ProgramArgReference
}

func (instruction *TestInstruction) Eval(state *ProgramState) error {
	state.ExternalData.(ValueSetterEnergy).SetValue(
		instruction.In.GetValue(state).GetFloat64(),
	)
	return nil
}

func (instruction *TestInstruction) String() string {
	return fmt.Sprintf("TEST %s", instruction.In)
}

func (instruction *TestInstruction) GetArgsCount() int {
	return 1
}

func (instruction *TestInstruction) GetArg(index int) ProgramInstructionArg {
	return instruction.In
}

func (instruction *TestInstruction) SetArg(
	index int, arg ProgramInstructionArg,
) {
	instruction.In = arg.(ProgramArgReference)
}

func (instruction *TestInstruction) Init() {
	instruction.In = FloatRegister(0)
}

func (instruction *TestInstruction) Copy() ProgramInstruction {
	instructionCopy := *instruction
	return &instructionCopy
}
