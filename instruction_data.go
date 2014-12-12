package main

import "fmt"

type DataInstruction struct {
	Out        ProgramArgRegister
	TableIndex Index
	CellIndex  Index
}

func (instruction *DataInstruction) Eval(state *ProgramState) error {
	value := state.ExternalData.(TableDataGetter).GetData(
		instruction.TableIndex.GetInt(),
		instruction.CellIndex.GetInt(),
	)

	instruction.Out.SetValue(state, FloatValue(value))

	return nil
}

func (instruction *DataInstruction) String() string {
	return fmt.Sprintf("DATA %s %s %s",
		instruction.Out,
		instruction.TableIndex,
		instruction.CellIndex,
	)
}

func (instruction *DataInstruction) GetArgsCount() int {
	return 3
}

func (instruction *DataInstruction) GetArg(index int) ProgramInstructionArg {
	switch index {
	case 0:
		return instruction.Out
	case 1:
		return instruction.TableIndex
	case 2:
		return instruction.CellIndex
	}

	return nil
}

func (instruction *DataInstruction) SetArg(
	index int, arg ProgramInstructionArg,
) {
	switch index {
	case 0:
		instruction.Out = arg.(ProgramArgRegister)
	case 1:
		instruction.TableIndex = arg.(Index)
	case 2:
		instruction.CellIndex = arg.(Index)
	}
}

func (instruction *DataInstruction) Init() {
	instruction.Out = FloatRegister(0)
	instruction.TableIndex = Index(0)
	instruction.CellIndex = Index(0)
}

func (instruction *DataInstruction) Copy() ProgramInstruction {
	instructionCopy := *instruction
	return &instructionCopy
}
