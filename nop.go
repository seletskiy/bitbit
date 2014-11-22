package main

type ProgramInstructionNop struct{}

func (op *ProgramInstructionNop) Eval(s *ProgramState) {
	// no op
}

func (op *ProgramInstructionNop) String() string {
	return "NOP"
}
