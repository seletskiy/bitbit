package main

import (
	"fmt"
	"math"
)

type ProgramInstructionPow struct {
	Out    ProgramArgRegister
	Degree ProgramArgReference
}

func (pow *ProgramInstructionPow) Eval(state *ProgramState) error {
	base := pow.Out.GetValue(state).GetFloat64()
	degree := pow.Degree.GetValue(state).GetFloat64()

	pow.Out.SetValue(state, FloatValue(
		math.Pow(base, degree),
	))

	return nil
}

func (pow *ProgramInstructionPow) String() string {
	return fmt.Sprintf("POW %s %s", pow.Out, pow.Degree)
}

func (pow *ProgramInstructionPow) GetArgsCount() int {
	return 2
}

func (pow *ProgramInstructionPow) GetArg(index int) ProgramInstructionArg {
	switch index {
	case 0:
		return pow.Out
	case 1:
		return pow.Degree
	}

	return nil
}

func (pow *ProgramInstructionPow) SetArg(
	index int, arg ProgramInstructionArg,
) {
	switch index {
	case 0:
		pow.Out = arg.(ProgramArgRegister)
	case 1:
		pow.Degree = arg.(ProgramArgReference)
	}
}

func (pow *ProgramInstructionPow) Init() {
	pow.Out = FloatRegister(0)
	pow.Degree = FloatValue(0)
}

func (pow *ProgramInstructionPow) Copy() ProgramInstruction {
	instructionCopy := *pow
	return &instructionCopy
}
