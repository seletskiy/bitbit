package main

import (
	"fmt"
	"math/rand"
)

type RandInstructionVariant struct {
	Instruction ProgramInstruction
	Weight      float64
}

func RandProgramInstructionJumpValue(
	maxInstructionNumber int,
) ProgramArgJump {
	return ForwardJump(rand.Intn(maxInstructionNumber) + 1)
}

func RandProgramInstructionOutValue(
	maxRegisterNumber int,
) ProgramArgRegister {
	return FloatRegister(rand.Intn(maxRegisterNumber))
}

func RandProgramInstructionInValue(
	valueVarianceGenerator func() float64,
	referenceProbability float64,
	maxRegisterNumber int,
) ProgramArgReference {
	var value ProgramArgReference

	if rand.Float64() < referenceProbability {
		value = RandProgramInstructionOutValue(maxRegisterNumber)
	} else {
		value = FloatValue(valueVarianceGenerator())
	}

	return value
}
func ChooseWeighted(variants []float64) int {
	sum := 0.0
	for _, variant := range variants {
		sum += variant
	}

	random := rand.Float64()

	for variantIndex, variant := range variants {
		if random >= variant/sum {
			random -= variant / sum
			continue
		} else {
			return variantIndex
		}
	}

	return 0
}

func RandProgramInstruction(
	referenceProbability float64,
	valueVarianceGenerator func() float64,
	maxRegisterNumber int,
	maxInstructionNumber int,
	variants []RandInstructionVariant,
) ProgramInstruction {
	weights := make([]float64, len(variants))
	for i, variant := range variants {
		weights[i] = variant.Weight
	}

	chosenOne := variants[ChooseWeighted(weights)].Instruction.Copy()
	chosenOne.Init()

	argsNumber := chosenOne.GetArgsCount()
	for argIndex := 0; argIndex < argsNumber; argIndex++ {
		switch chosenOne.GetArg(argIndex).(type) {
		case ProgramArgRegister:
			chosenOne.SetArg(argIndex,
				RandProgramInstructionOutValue(maxRegisterNumber),
			)
		case ProgramArgReference:
			chosenOne.SetArg(argIndex, RandProgramInstructionInValue(
				valueVarianceGenerator,
				referenceProbability,
				maxRegisterNumber,
			))
		case ProgramArgJump:
			chosenOne.SetArg(argIndex,
				RandProgramInstructionJumpValue(maxInstructionNumber),
			)
		default:
			panic("unknown arg type")
		}
	}

	return chosenOne
}

func RandProgramInstructionSet(
	programLength int,
	referenceProbability float64,
	valueVarianceGenerator func() float64,
	maxRegisterNumber int,
	variants []RandInstructionVariant,
) []ProgramInstruction {
	instructions := make([]ProgramInstruction, programLength)

	for index, _ := range instructions {
		instructions[index] = RandProgramInstruction(
			referenceProbability,
			valueVarianceGenerator,
			maxRegisterNumber,
			programLength,
			variants,
		)
	}

	return instructions
}

func RandProgram(
	layout *Program,
	referenceProbability float64,
	valueVarianceGenerator func() float64,
	maxRegisterNumber int,
	instructionVariants []RandInstructionVariant,
) *Program {
	program := make(Program, len(*layout))

	instructions := RandProgramInstructionSet(
		len(program),
		referenceProbability,
		valueVarianceGenerator,
		maxRegisterNumber,
		instructionVariants,
	)

	for index, codepoint := range *layout {
		program[index] = Codepoint{
			Label:       codepoint.Label,
			Instruction: instructions[index],
		}
	}

	return &program
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
	memorySize int,
	program *Program,
	initialEnergy Energy,
	externalData interface{},
) *ProgoBact {
	bacteria := &ProgoBact{
		&SimpleBacteria{
			Energy: initialEnergy,
			Chromosome: &SimpleChromosome{
				&ProgoDNA{program},
			},
			Plasmids: nil,
		},
		&ProgramState{
			IPS:          0,
			Memory:       NewProgramMemory(memorySize),
			ExternalData: externalData,
		},
	}

	return bacteria
}
