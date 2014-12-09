package main

import "fmt"

type ProgramInstructionAdd struct {
	Out ProgramArgRegister
	In  ProgramArgReference
}

func (add *ProgramInstructionAdd) Eval(state *ProgramState) error {
	a := add.Out.GetValue(state).GetFloat64()
	b := add.In.GetValue(state).GetFloat64()
	add.Out.SetValue(state, FloatValue(a+b))

	return nil
}

func (add *ProgramInstructionAdd) String() string {
	return fmt.Sprintf("ADD %s %s", add.Out, add.In)
}

func (add *ProgramInstructionAdd) GetArgsCount() int {
	return 2
}

func (add *ProgramInstructionAdd) GetArg(index int) ProgramInstructionArg {
	switch index {
	case 0:
		return add.Out
	case 1:
		return add.In
	}

	return nil
}

func (add *ProgramInstructionAdd) SetArg(index int, arg ProgramInstructionArg) {
	switch index {
	case 0:
		add.Out = arg.(ProgramArgRegister)
	case 1:
		add.In = arg.(ProgramArgReference)
	}
}

func (add *ProgramInstructionAdd) Init() {
	add.Out = FloatRegister(0)
	add.In = FloatValue(0)
}

func (add *ProgramInstructionAdd) Copy() ProgramInstruction {
	instructionCopy := *add
	return &instructionCopy
}
