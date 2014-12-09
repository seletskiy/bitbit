package main

import (
	"fmt"
)

type ProgramInstructionPosition struct {
	Amount ProgramArgReference
	TTL    ProgramArgReference
}

func (op *ProgramInstructionPosition) Eval(state *ProgramState) error {
	connection := state.ExternalData.(*MarketEnergy)
	return connection.PlacePosition(
		op.Amount.GetValue(state).GetFloat64(),
		int(op.TTL.GetValue(state).GetFloat64()),
	)
}

func (op *ProgramInstructionPosition) String() string {
	return fmt.Sprintf("POS %s %s", op.Amount, op.TTL)
}

func (op *ProgramInstructionPosition) GetArgsCount() int {
	return 2
}

func (op *ProgramInstructionPosition) GetArg(
	index int,
) ProgramInstructionArg {
	switch index {
	case 0:
		return op.Amount
	case 1:
		return op.TTL
	}

	return nil
}

func (op *ProgramInstructionPosition) SetArg(
	index int, arg ProgramInstructionArg,
) {
	switch index {
	case 0:
		op.Amount = arg.(ProgramArgReference)
	case 1:
		op.TTL = arg.(ProgramArgReference)
	}
}

func (op *ProgramInstructionPosition) Init() {
	op.Amount = FloatValue(0)
	op.TTL = FloatValue(0)
}

func (op *ProgramInstructionPosition) Copy() ProgramInstruction {
	instructionCopy := *op
	return &instructionCopy
}
