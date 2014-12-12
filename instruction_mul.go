package main

import "fmt"

type ProgramInstructionMul struct {
	Out ProgramArgRegister
	In  ProgramArgReference
}

func (mul *ProgramInstructionMul) Eval(state *ProgramState) error {
	a := mul.Out.GetValue(state).GetFloat64()
	b := mul.In.GetValue(state).GetFloat64()

	mul.Out.SetValue(state, FloatValue(a*b))

	return nil
}

func (mul *ProgramInstructionMul) String() string {
	return fmt.Sprintf("MUL %s %s", mul.Out, mul.In)
}

func (mul *ProgramInstructionMul) GetArgsCount() int {
	return 2
}

func (mul *ProgramInstructionMul) GetArg(index int) ProgramInstructionArg {
	switch index {
	case 0:
		return mul.Out
	case 1:
		return mul.In
	}

	return nil
}

func (mul *ProgramInstructionMul) SetArg(
	index int, arg ProgramInstructionArg,
) {
	switch index {
	case 0:
		mul.Out = arg.(ProgramArgRegister)
	case 1:
		mul.In = arg.(ProgramArgReference)
	}
}

func (mul *ProgramInstructionMul) Init() {
	mul.Out = FloatRegister(0)
	mul.In = FloatValue(0)
}

func (mul *ProgramInstructionMul) Copy() ProgramInstruction {
	instructionCopy := *mul
	return &instructionCopy
}
