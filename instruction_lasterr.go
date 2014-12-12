package main

import "fmt"

type LastErrorInstruction struct {
	Out ProgramArgRegister
}

func (instruction *LastErrorInstruction) Eval(state *ProgramState) error {
	instruction.Out.SetValue(state,
		FloatValue(state.ExternalData.(ErrorGetterEnergy).GetLastError()),
	)

	return nil
}

func (instruction *LastErrorInstruction) String() string {
	return fmt.Sprintf("LASTERR %s", instruction.Out)
}

func (instruction *LastErrorInstruction) GetArgsCount() int {
	return 1
}

func (instruction *LastErrorInstruction) GetArg(index int) ProgramInstructionArg {
	return instruction.Out
}

func (instruction *LastErrorInstruction) SetArg(
	index int, arg ProgramInstructionArg,
) {
	instruction.Out = arg.(ProgramArgRegister)
}

func (instruction *LastErrorInstruction) Init() {
	instruction.Out = FloatRegister(0)
}

func (instruction *LastErrorInstruction) Copy() ProgramInstruction {
	instructionCopy := *instruction
	return &instructionCopy
}
