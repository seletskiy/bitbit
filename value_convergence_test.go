package main

import "testing"

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

func converge(
	valueGenerator ValueGenerator,
	validator PopulationValidator,
	additionalInstructions []RandInstructionVariant,
) (population Population, tick int) {
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

	population = make(Population, initialPopulationSize)
	for i := 0; i < initialPopulationSize; i++ {
		program := RandProgram(
			programLayout,
			programReferenceProbability,
			defaultVarianceGenerator,
			programMemorySize,
			programInstructionVariants,
		)

		energy := &ErrorBasedEnergy{
			ReproductiveEnergy: &ReproductiveEnergy{
				Potential: 1,
			},

			ConsiderZero: 1e-6,

			TargetValue: externalValueStore,
		}

		population[i] = RandProgoBact(
			programMemorySize, program, energy,
			dataSource,
		)
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
		externalValueStore.Set(valueGenerator.Generate())

		environment.Simulate(&population)

		if validator.Validate(population) {
			break
		}

		tick++
	}

	return population, tick
}

func validate(
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
		environment.Simulate(&population)

		ticks--
	}

	return validator.Validate(population)
}

func TestCanConvergeToConstantValue(t *testing.T) {
	validator := MedianErrorValidator{Threshold: 0.001}

	validator.LogPrefix = "[evolve] "
	population, _ := converge(
		ConstGenerator(12345.6),
		validator,
		[]RandInstructionVariant{},
	)

	validator.LogPrefix = "[check]  "
	if !validate(population, validator, 100) {
		t.Fatalf("population evolve, but failed to continue it's life")
	}
}

//func TestCanEstimateLinearFunction(t *testing.T) {
//    converge(
//        &LinearGenerator{A: 10.0, B: 5.0},
//        0.001,
//        []RandInstructionVariant{},
//    )
//}
