package main

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
)

type ConstGenerator struct {
	Value float64
}

func (generator ConstGenerator) Generate() float64 {
	return generator.Value
}

func (generator ConstGenerator) GetData(_, _ int) float64 {
	return 0
}

func (generator ConstGenerator) Simulate() {

}

type LinearGenerator struct {
	A float64
	B float64
	X float64
}

func (generator *LinearGenerator) Generate() float64 {
	return generator.A*generator.X + generator.B
}

func (generator LinearGenerator) GetData(tableIndex, cellIndex int) float64 {
	if tableIndex == 0 && cellIndex == 0 {
		return generator.X
	} else {
		return 0
	}
}

func (generator *LinearGenerator) Simulate() {
	//generator.X += 0.1
	generator.X = rand.Float64() * 100.0
}

func (generator LinearGenerator) String() string {
	return fmt.Sprintf("Y = %.3f X + %3.f = %.3f, X = %.3f",
		generator.A,
		generator.B,
		generator.Generate(),
		generator.X,
	)
}

type PopulationValidator interface {
	Validate(step string, population Population) bool
}

func converge(
	valueGenerator ValueGenerator,
	validator PopulationValidator,
	additionalInstructions []RandInstructionVariant,
	stabilityCheckLength int,
) (population Population, tick int) {
	programInstructionVariants := []RandInstructionVariant{
		{&ProgramInstructionAdd{},
			addInstructionProbability},
		{&ProgramInstructionMov{},
			movInstructionProbability},
		{&ProgramInstructionDiv{},
			divInstructionProbability},
		{&ProgramInstructionMul{},
			mulInstructionProbability},
		{&ProgramInstructionNop{},
			nopInstructionProbability},
		{&ProgramInstructionJumpGreaterThan{},
			jumpGreaterThanInstructionProbability},
		{&ProgramInstructionCls{},
			clsInstructionProbability},

		{&TestInstruction{}, 0.5},
		{&DataInstruction{}, 0.5},
	}

	programInstructionVariants = append(
		programInstructionVariants,
		additionalInstructions...,
	)

	programLayout := RandProgramLayout(programLength)

	//externalValueStore := ValueStore{}

	population = make(Population, initialPopulationSize)
	for i := 0; i < initialPopulationSize; i++ {
		program := RandProgram(
			programLayout,
			programReferenceProbability,
			defaultVarianceGenerator,
			programMemorySize,
			maxDataIndex,
			programInstructionVariants,
		)

		energy := &GeneratedValueEnergy{
			Base: &ReproductiveEnergy{
				Potential: 1,
			},

			ConsiderZero: 1e-6,

			TargetValueGenerator: valueGenerator,
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

	tick = 0
	for len(population) > 0 {
		valueGenerator.Simulate()
		environment.Simulate(&population)

		validated := false
		if validator.Validate("evolve", population) {
			validated = validate(
				valueGenerator,
				population, validator, stabilityCheckLength)
		}

		if validated {
			break
		}

		tick++
	}

	return population, tick
}

func validate(
	valueGenerator ValueGenerator,
	population Population, validator PopulationValidator, ticks int,
) bool {
	environment := SimpleEnvironment{
		Rules: []Rules{
			defaultSumulationRules,
			//defaultBacterialRules,
			defaultReapRules,
		},
	}

	for ticks > 0 {
		valueGenerator.Simulate()
		environment.Simulate(&population)

		ticks--
	}

	return validator.Validate("validate", population)
}

func getBest(population Population) Creature {
	sort.Sort(ByEnergy(population))

	return population[len(population)-1]
}

func TestCanConvergeToConstantValue(t *testing.T) {
	validator := MedianErrorValidator{Threshold: 0.001}

	_, _ = converge(
		ConstGenerator{12345.6},
		validator,
		[]RandInstructionVariant{},
		100,
	)
}

func TestCanEstimateLinearFunction(t *testing.T) {
	validator := MedianErrorValidator{Threshold: 0.001}

	population, _ := converge(
		&LinearGenerator{A: 3.0, B: 0.0},
		validator,
		[]RandInstructionVariant{},
		100,
	)

	fmt.Printf("BEST:\n%s", getBest(population))
}
