package main

import (
	"fmt"
)

type ProgramInstructionFree struct {
	Funds ProgramArgRegister
	Goods ProgramArgRegister
}

func (free *ProgramInstructionFree) Eval(state *ProgramState) error {
	connection := state.ExternalData.(*MarketEnergy).MarketConnection
	free.Funds.SetValue(state, FloatValue(connection.GetFreeFunds()))
	free.Goods.SetValue(state, FloatValue(connection.GetFreeGoods()))
	return nil
}

func (free *ProgramInstructionFree) String() string {
	return fmt.Sprintf("FREE %s %s", free.Funds, free.Goods)
}

func (free *ProgramInstructionFree) GetArgsCount() int {
	return 2
}

func (free *ProgramInstructionFree) GetArg(index int) ProgramInstructionArg {
	switch index {
	case 0:
		return free.Funds
	case 1:
		return free.Goods
	}

	return nil
}

func (free *ProgramInstructionFree) SetArg(
	index int, arg ProgramInstructionArg,
) {
	switch index {
	case 0:
		free.Funds = arg.(ProgramArgRegister)
	case 1:
		free.Goods = arg.(ProgramArgRegister)
	}
}

func (free *ProgramInstructionFree) Init() {
	free.Funds = FloatRegister(0)
	free.Goods = FloatRegister(0)
}

func (free *ProgramInstructionFree) Copy() ProgramInstruction {
	instructionCopy := *free
	return &instructionCopy
}
