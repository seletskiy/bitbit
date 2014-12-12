package main

import (
	"fmt"
	"math"
	"sort"
	"testing"
)

type ValueStore struct {
	Data float64
}

func (store *ValueStore) Set(value float64) {
	store.Data = value
}

func (store ValueStore) Get() float64 {
	return store.Data
}

type ValueGenerator interface {
	Generate() float64
}

type ConstGenerator float64

func (value ConstGenerator) Generate() float64 {
	return float64(value)
}

type LinearGenerator struct {
	A float64
	B float64
	X float64
}

func (generator LinearGenerator) Generate() float64 {
	generator.X += 1
	return generator.A + generator.B*generator.X
}

type PopulationValidator interface {
	Validate(Population) bool
}

type MedianErrorValidator struct {
	Threshold float64
}

func (validator MedianErrorValidator) Validate(population Population) bool {
	populationErrors := []float64{}
	for _, creature := range population {
		if creature.GetAge() <= 0 {
			continue
		}

		populationErrors = append(populationErrors,
			math.Abs(
				creature.GetEnergy().(ErrorGetterEnergy).GetError(),
			),
		)
	}
	sort.Float64s(populationErrors)

	minError := populationErrors[0]
	medianError := populationErrors[len(populationErrors)/2]
	totalError := 0.0
	for _, val := range populationErrors {
		totalError += val
	}

	fmt.Printf(
		"avg err: %10.4f med: %10.4f min: %10.4f (%3d/%3d)\n",
		totalError/float64(len(populationErrors)),
		medianError,
		minError,
		len(populationErrors),
		len(population),
	)

	if medianError <= validator.Threshold {
		return true
	} else {
		return false
	}
}

func converge(
	valueGenerator ValueGenerator,
	validator PopulationValidator,
	additionalInstructions []RandInstructionVariant,
) Population {
	programInstructionVariants := []RandInstructionVariant{
		{&ProgramInstructionAdd{},
			addInstructionProbability},
		{&ProgramInstructionMov{},
			movInstructionProbability},
		{&ProgramInstructionDiv{},
			divInstructionProbability},
		{&ProgramInstructionNop{},
			nopInstructionProbability},
		{&ProgramInstructionJumpGreaterThan{},
			jumpGreaterThanInstructionProbability},

		{&TestInstruction{}, 0.5},
		{&LastErrorInstruction{}, 0.5},
	}

	programInstructionVariants = append(
		programInstructionVariants,
		additionalInstructions...,
	)

	programLayout := RandProgramLayout(programLength)

	externalValueStore := ValueStore{}

	population := make(Population, initialPopulationSize)
	for i := 0; i < initialPopulationSize; i++ {
		program := RandProgram(
			programLayout,
			programReferenceProbability,
			defaultVarianceGenerator,
			programMemorySize,
			programInstructionVariants,
		)

		energy := &ExternalValueEnergy{
			ErrorBasedEnergy: &ErrorBasedEnergy{
				ReproductiveEnergy: &ReproductiveEnergy{
					Potential: 1,
				},
			},

			Store: &externalValueStore,
		}

		population[i] = RandProgoBact(programMemorySize, program, energy)
	}

	environment := SimpleEnvironment{
		Rules: []Rules{
			defaultSumulationRules,
			defaultAggressiveSelectionRules,
			defaultReapRules,
			defaultAggressiveReproduceRules,
			//defaultBacterialRules,
			defaultMutateRules(programInstructionVariants),
			//defaultReapRules,
		},
	}

	tick := 0
	for len(population) > 0 {
		externalValueStore.Set(valueGenerator.Generate())

		environment.Simulate(&population)

		if validator.Validate(population) {
			break
		}

		tick++
	}

	return population
}

func TestCanConvergeToConstantValue(t *testing.T) {
	converge(
		ConstGenerator(12345.6),
		MedianErrorValidator{0.001},
		[]RandInstructionVariant{},
	)
}

//func TestCanEstimateLinearFunction(t *testing.T) {
//    converge(
//        &LinearGenerator{A: 10.0, B: 5.0},
//        0.001,
//        []RandInstructionVariant{},
//    )
//}
