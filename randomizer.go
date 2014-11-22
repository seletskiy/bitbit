package main

import (
	"fmt"
	"math/rand"
)

func RandProgramInstructionToValueTuple(
	valueVariance float64,
	referenceOccurenceRate float64,
	maxRegisterNumber int,
) (Register, Reference) {
	var to Register
	var from Reference

	to = FloatRegister(rand.Intn(maxRegisterNumber))
	if rand.Float64() < referenceOccurenceRate {
		from = FloatRegister(rand.Intn(maxRegisterNumber))
	} else {
		from = FloatValue((rand.Float64() - 0.5) * valueVariance)
	}

	return to, from
}

func RandProgramInstruction(maxRegisterNumber int) ProgramInstruction {
	variants := []struct {
		Instruction ProgramInstruction
		Weight      float64
	}{
		{&ProgramInstructionAdd{}, 1.0},
		{&ProgramInstructionMov{}, 1.0},
		{&ProgramInstructionDiv{}, 0.3},
		{&ProgramInstructionFun{}, 0.1},
		{&ProgramInstructionNop{}, 1.0},
	}

	sum := 0.0
	for _, variant := range variants {
		sum += variant.Weight
	}

	random := rand.Float64()

	var chosenOne ProgramInstruction

	for _, variant := range variants {
		if random >= variant.Weight/sum {
			random -= variant.Weight / sum
			continue
		} else {
			chosenOne = variant.Instruction
			break
		}
	}

	switch instruction := chosenOne.(type) {
	case ProgramInstructionArgsToValueSetter:
		to, value := RandProgramInstructionToValueTuple(
			10.0,
			0.5,
			maxRegisterNumber,
		)
		instruction.SetTo(to)
		instruction.SetValue(value)
	case ProgramInstructionArgsValueSetter:
		_, value := RandProgramInstructionToValueTuple(
			10.0,
			0.5,
			maxRegisterNumber,
		)
		instruction.SetValue(value)
	case *ProgramInstructionNop:
		// empty
	}

	return chosenOne
}

func RandProgram(layout *Program, maxRegisterNumber int) *Program {
	newProg := make(Program, len(*layout))

	for i, point := range *layout {
		newProg[i] = Codepoint{
			Label:       point.Label,
			Instruction: RandProgramInstruction(maxRegisterNumber),
		}
	}

	return &newProg
}

func RandInstructionLabel(length int) string {
	var label string

	for i := 0; i < length; i++ {
		label += fmt.Sprintf("%1x", rand.Int63()&0xf)
	}

	return label
}

func RandProgramLayout(length int) *Program {
	program := make(Program, length)

	for i := 0; i < length; i++ {
		program[i] = Codepoint{
			Label:       RandInstructionLabel(length),
			Instruction: &ProgramInstructionNop{},
		}
	}

	return &program
}

func RandProgoBact(
	layout *Program,
	memSize int,
	initialEnergy float64,
	externalData interface{},
) *ProgoBact {
	mem := NewProgramMemory(memSize)

	return &ProgoBact{
		&SimpleBacteria{
			Energy: initialEnergy,
			Chromosome: &SimpleChromosome{
				&ProgDNA{
					RandProgram(layout, memSize),
				},
			},
			Plasmids: nil,
		},
		&ProgramState{
			IPS:          0,
			Memory:       mem,
			ExternalData: externalData,
		},
	}
}
