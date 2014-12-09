package main

import "fmt"

type ProgramInstructionMov struct {
	Out ProgramArgRegister
	In  ProgramArgReference
}

func (mov *ProgramInstructionMov) Eval(state *ProgramState) error {
	mov.Out.SetValue(state, mov.In.GetValue(state))
	return nil
}

func (mov *ProgramInstructionMov) String() string {
	return fmt.Sprintf("MOV %s %s", mov.Out, mov.In)
}

func (mov *ProgramInstructionMov) GetArgsCount() int {
	return 2
}

func (mov *ProgramInstructionMov) GetArg(index int) ProgramInstructionArg {
	switch index {
	case 0:
		return mov.Out
	case 1:
		return mov.In
	}

	return nil
}

func (mov *ProgramInstructionMov) SetArg(
	index int, arg ProgramInstructionArg,
) {
	switch index {
	case 0:
		mov.Out = arg.(ProgramArgRegister)
	case 1:
		mov.In = arg.(ProgramArgReference)
	}
}

func (mov *ProgramInstructionMov) Init() {
	mov.Out = FloatRegister(0)
	mov.In = FloatValue(0)
}

func (mov *ProgramInstructionMov) Copy() ProgramInstruction {
	instructionCopy := *mov
	return &instructionCopy
}
