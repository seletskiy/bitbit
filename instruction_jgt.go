package main

import (
	"fmt"
)

type ProgramInstructionJumpGreaterThan struct {
	Jump   ProgramArgJump
	A      ProgramArgReference
	B      ProgramArgReference
	Jumped bool
}

func (jump *ProgramInstructionJumpGreaterThan) Eval(state *ProgramState) error {
	a := jump.A.GetValue(state).GetFloat64()
	b := jump.B.GetValue(state).GetFloat64()
	if a > b {
		jump.Jump.Apply(state)
		jump.Jumped = true
	} else {
		jump.Jumped = false
	}

	return nil
}

func (jump *ProgramInstructionJumpGreaterThan) String() string {
	return fmt.Sprintf(
		"JGT %d %s %s",
		jump.Jump,
		jump.A, jump.B,
	)
}

func (jump *ProgramInstructionJumpGreaterThan) GetArg(
	index int,
) ProgramInstructionArg {
	switch index {
	case 0:
		return jump.Jump
	case 1:
		return jump.A
	case 2:
		return jump.B
	}

	return nil
}

func (jump *ProgramInstructionJumpGreaterThan) GetArgsCount() int {
	return 3
}

func (jump *ProgramInstructionJumpGreaterThan) SetArg(
	index int, arg ProgramInstructionArg,
) {
	switch index {
	case 0:
		jump.Jump = arg.(ProgramArgJump)
	case 1:
		jump.A = arg.(ProgramArgReference)
	case 2:
		jump.B = arg.(ProgramArgReference)
	}
}

func (jump *ProgramInstructionJumpGreaterThan) Init() {
	jump.Jump = ForwardJump(0)
	jump.A = FloatValue(0)
	jump.B = FloatValue(0)
}

func (jump *ProgramInstructionJumpGreaterThan) Copy() ProgramInstruction {
	instructionCopy := *jump
	return &instructionCopy
}
