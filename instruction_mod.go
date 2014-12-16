package main

import (
	"fmt"
	"math"
)

type ProgramInstructionMod struct {
	Out ProgramArgRegister
	In  ProgramArgReference
}

func (mod *ProgramInstructionMod) Eval(state *ProgramState) error {
	mod.Out.SetValue(state,
		FloatValue(math.Mod(
			mod.Out.GetValue(state).GetFloat64(),
			mod.In.GetValue(state).GetFloat64(),
		)),
	)

	return nil
}

func (mod *ProgramInstructionMod) String() string {
	return fmt.Sprintf("MOD %s %s", mod.Out, mod.In)
}

func (mod *ProgramInstructionMod) GetArgsCount() int {
	return 2
}

func (mod *ProgramInstructionMod) GetArg(index int) ProgramInstructionArg {
	switch index {
	case 0:
		return mod.Out
	case 1:
		return mod.In
	}

	return nil
}

func (mod *ProgramInstructionMod) SetArg(
	index int, arg ProgramInstructionArg,
) {
	switch index {
	case 0:
		mod.Out = arg.(ProgramArgRegister)
	case 1:
		mod.In = arg.(ProgramArgReference)
	}
}

func (mod *ProgramInstructionMod) Init() {
	mod.Out = FloatRegister(0)
	mod.In = FloatValue(0)
}

func (mod *ProgramInstructionMod) Copy() ProgramInstruction {
	instructionCopy := *mod
	return &instructionCopy
}
