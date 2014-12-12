package main

type ProgramInstructionCls struct{}

func (cls *ProgramInstructionCls) Eval(s *ProgramState) error {
	s.Memory.Zero()
	return nil
}

func (cls *ProgramInstructionCls) String() string {
	return "CLS"
}

func (cls *ProgramInstructionCls) GetArgsCount() int {
	return 0
}

func (cls *ProgramInstructionCls) GetArg(index int) ProgramInstructionArg {
	return nil
}

func (cls *ProgramInstructionCls) SetArg(
	index int, arg ProgramInstructionArg,
) {
	// noop
}

func (cls *ProgramInstructionCls) Init() {

}

func (cls *ProgramInstructionCls) Copy() ProgramInstruction {
	instructionCopy := *cls
	return &instructionCopy
}
