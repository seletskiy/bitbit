package main

type ProgramInstructionNop struct{}

func (nop *ProgramInstructionNop) Eval(s *ProgramState) error {
	return nil
}

func (nop *ProgramInstructionNop) String() string {
	return "NOP"
}

func (nop *ProgramInstructionNop) GetArgsCount() int {
	return 0
}

func (nop *ProgramInstructionNop) GetArg(index int) ProgramInstructionArg {
	return nil
}

func (nop *ProgramInstructionNop) SetArg(
	index int, arg ProgramInstructionArg,
) {
	// noop
}

func (nop *ProgramInstructionNop) Init() {

}

func (nop *ProgramInstructionNop) Copy() ProgramInstruction {
	instructionCopy := *nop
	return &instructionCopy
}
