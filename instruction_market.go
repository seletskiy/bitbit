package main

import (
	"fmt"
)

type ProgramInstructionMarket struct {
	High ProgramArgRegister
	Low  ProgramArgRegister
}

func (market *ProgramInstructionMarket) Eval(state *ProgramState) error {
	connection := state.ExternalData.(*MarketEnergy).MarketConnection
	market.High.SetValue(state, FloatValue(connection.Market.GetHigh()))
	market.Low.SetValue(state, FloatValue(connection.Market.GetLow()))
	return nil
}

func (market *ProgramInstructionMarket) String() string {
	return fmt.Sprintf("MRKT %s %s", market.High, market.Low)
}

func (market *ProgramInstructionMarket) GetArgsCount() int {
	return 2
}

func (market *ProgramInstructionMarket) GetArg(
	index int,
) ProgramInstructionArg {
	switch index {
	case 0:
		return market.High
	case 1:
		return market.Low
	}

	return nil
}

func (market *ProgramInstructionMarket) SetArg(
	index int, arg ProgramInstructionArg,
) {
	switch index {
	case 0:
		market.High = arg.(ProgramArgRegister)
	case 1:
		market.Low = arg.(ProgramArgRegister)
	}
}

func (market *ProgramInstructionMarket) Init() {
	market.High = FloatRegister(0)
	market.Low = FloatRegister(0)
}

func (market *ProgramInstructionMarket) Copy() ProgramInstruction {
	instructionCopy := *market
	return &instructionCopy
}
