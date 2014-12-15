package main

import (
	"math"
	"math/rand"
)

const (
	initialPopulationSize = 1000
	initialFunds          = 100000.0
	initialGoods          = 10000.0

	programLength               = 30
	programMemorySize           = 2
	programReferenceProbability = 0.4
	programValueVariance        = 100.0

	bootstrapFunds = 10.0
	bootstrapGoods = 1.0

	addInstructionProbability             = 0.8
	movInstructionProbability             = 0.8
	divInstructionProbability             = 0.1
	mulInstructionProbability             = 0.5
	nopInstructionProbability             = 1.0
	jumpGreaterThanInstructionProbability = 0.2
	clsInstructionProbability             = 0.05
	powInstructionProbability             = 0.2

	floatValueMutationProbability = 0.9
	indexMutationProbability      = 0.5
	referenceMutationProbability  = 0.4
	registerMutationProbability   = 0.3
	jumpMutationProbability       = 0.2

	smallChangeValueMutationProbability = 0.6

	dnaMutationProbability = 0.15
	dnaMutationMaxSize     = 2
	dnaMutationCount       = 1

	geneMutationProbability = 0.7

	diePercentile = 0.96

	minSelectionAge = 5
	minReproduceAge = 5

	maxDataIndex = 1
)

func defaultAggressiveReproduceRules(
	programInstructionVariants []RandInstructionVariant,
) AggressiveReproduceRules {
	return AggressiveReproduceRules{
		MutateRules: defaultMutateRules(programInstructionVariants),
		MinAge:      minReproduceAge,
	}
}

var defaultBacterialRules = BacterialGeneTransferRules{
	BirthTransferProbability:      0.5,
	ReproduceLossProbability:      0.01,
	TransferLossProbability:       0.02,
	ApplyPlasmidProbability:       0.3,
	ApplyLossProbability:          0.01,
	MaxPlasmidsNumber:             2,
	PlasmidPerAge:                 10,
	ExchangePlasmidsProbability:   0.3,
	MinAgeForExchange:             5,
	PlasmidPrefixLengthProportion: 0.3,
}

var defaultAggressiveSelectionRules = AggressiveNaturalSelectionRules{
	DiePercentile:      diePercentile,
	MinAge:             minSelectionAge,
	BasePopulationSize: initialPopulationSize,
}

func defaultVarianceGenerator() float64 {
	return (rand.Float64() - 0.5) * programValueVariance
}

func defaultMutateRules(
	programInstructionVariants []RandInstructionVariant,
) MutateRules {
	return MutateRules{
		DNAMutationProbability: dnaMutationProbability,
		DNAMutationMaxSize:     dnaMutationMaxSize,
		DNAMutationCount:       dnaMutationCount,

		GeneMutationProbability: geneMutationProbability,

		GeneGenerator: func(amount int) []Gene {
			return defaultGeneGenerator(amount, programInstructionVariants)
		},

		GeneMutator: defaultGeneMutator,
	}
}

func defaultGeneMutator(gene Gene) Gene {
	instruction := gene.(Codepoint).Instruction

	if instruction.GetArgsCount() == 0 {
		return nil
	}

	argsCount := instruction.GetArgsCount()
	weights := make([]float64, argsCount)
	for index, _ := range weights {
		switch instruction.GetArg(index).(type) {
		case ProgramArgRegister:
			weights[index] = registerMutationProbability
		case Index:
			weights[index] = indexMutationProbability
		case FloatValue:
			weights[index] = floatValueMutationProbability
		case ProgramArgReference:
			weights[index] = referenceMutationProbability
		case ProgramArgJump:
			weights[index] = jumpMutationProbability
		}
	}

	operandIndex := ChooseWeighted(weights)

	var mutatedArg ProgramInstructionArg
	switch concreteOperand := instruction.GetArg(operandIndex).(type) {
	case ProgramArgRegister:
		//Log(Debug, "MUTATE: GENE<%p> mutate as register", instruction)
		if rand.Float64() < smallChangeValueMutationProbability {
			mutatedArg = concreteOperand
			break
		}
		mutatedArg = RandProgramInstructionOutValue(
			programMemorySize,
		)
	case Index:
		//Log(Debug, "MUTATE: GENE<%p> mutate as index", instruction)
		currentValue := concreteOperand.GetValue(nil).GetInt()
		variance := currentValue
		if rand.Float64() < smallChangeValueMutationProbability {
			variance = 2
		}

		mutatedArg = Index(math.Abs(
			float64(rand.Intn(variance+1) - variance/2 + currentValue),
		))
	case FloatValue:
		//Log(Debug, "MUTATE: GENE<%p> mutate as float value", instruction)
		currentValue := concreteOperand.GetValue(nil).GetFloat64()

		variance := currentValue
		if rand.Float64() < smallChangeValueMutationProbability {
			variance = currentValue / 1000.0
		}

		mutatedArg = FloatValue(
			2*(rand.Float64()-0.5)*variance + currentValue,
		)
	case ProgramArgReference:
		//Log(Debug, "MUTATE: GENE<%p> mutate as reference", instruction)
		if rand.Float64() < smallChangeValueMutationProbability {
			mutatedArg = concreteOperand
			break
		}
		mutatedArg = RandProgramInstructionInValue(
			defaultVarianceGenerator,
			programReferenceProbability,
			programMemorySize,
		)
	case ProgramArgJump:
		//Log(Debug, "MUTATE: GENE<%p> mutate as jump", instruction)
		mutatedArg = RandProgramInstructionJumpValue(
			programLength,
		)
	}

	instruction.SetArg(operandIndex, mutatedArg)

	return gene
}

var defaultReapRules = ReapRules{}

func defaultGeneGenerator(
	amount int,
	programInstructionVariants []RandInstructionVariant,
) []Gene {
	instructions := RandProgramInstructionSet(
		amount,
		programReferenceProbability,
		defaultVarianceGenerator,
		programMemorySize,
		maxDataIndex,
		programInstructionVariants,
	)

	result := make([]Gene, amount)
	for index, instruction := range instructions {
		result[index] = Codepoint{
			Instruction: instruction,
		}
	}

	return result
}

var defaultSumulationRules = SimulationRules{}
