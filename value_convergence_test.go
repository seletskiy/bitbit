package main

import (
	"fmt"
	"math"
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

func (generator ConstGenerator) String() string {
	return fmt.Sprintf("%f", generator)
}

type CeilGenerator struct {
	X float64
}

func (generator *CeilGenerator) Generate() float64 {
	return math.Ceil(generator.X)
}

func (generator CeilGenerator) GetData(tableIndex, cellIndex int) float64 {
	return 0
}

func (generator *CeilGenerator) Simulate() {
	generator.X = rand.Float64() * 1000.0
}

func (generator CeilGenerator) String() string {
	return fmt.Sprintf("Y = CEIL( X ) = %.3f, X = %.3f",
		generator.Generate(),
		generator.X,
	)
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
	sign := -1.0
	if generator.X < generator.B {
		sign = 1.0
	}
	generator.X = sign * rand.Float64() * 1000.0
}

func (generator LinearGenerator) String() string {
	return fmt.Sprintf("Y = %.3f X + %3.f = %.3f, X = %.3f",
		generator.A,
		generator.B,
		generator.Generate(),
		generator.X,
	)
}

type SinGenerator struct {
	History      []float64
	HistoryIndex int
	Variation    float64
	Amplitude    float64
	Period       float64
	X            float64
}

func (generator *SinGenerator) Generate() float64 {
	return generator.Amplitude * math.Sin(generator.Period*generator.X)
}

func (generator SinGenerator) GetData(tableIndex, cellIndex int) float64 {
	switch tableIndex {
	case 0:
		if cellIndex == 0 {
			return generator.X
		}
	case 1:
		if cellIndex < len(generator.History) {
			getIndex := generator.HistoryIndex - cellIndex
			if getIndex < 0 {
				getIndex = 0
			}

			return generator.History[getIndex%cap(generator.History)]
		}
	}

	return 0
}

func (generator *SinGenerator) Simulate() {
	generator.History[generator.HistoryIndex%cap(generator.History)] = generator.Generate()
	generator.HistoryIndex++

	generator.X += 0.05
}

func (generator SinGenerator) String() string {
	return fmt.Sprintf("Y = %.3f SIN( %.3f X ) = %.3f, X = %.3f",
		generator.Amplitude,
		generator.Period,
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
		//{&ProgramInstructionZero{},
		//    zeroInstructionProbability},
		{&ProgramInstructionMod{},
			modInstructionProbability},
		{&ProgramInstructionDiv{},
			divInstructionProbability},
		{&ProgramInstructionMul{},
			mulInstructionProbability},
		{&ProgramInstructionNop{},
			nopInstructionProbability},
		{&ProgramInstructionJumpGreaterThan{},
			jumpGreaterThanInstructionProbability},
		{&ProgramInstructionPow{},
			powInstructionProbability},
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
			EloEnergy: &EloEnergy{
				Base: &ReproductiveEnergy{
					Potential: 1,
				},

				Score:     1500,
				BaseScore: 1500,
			},

			ZeroThreshold: 1e-12,

			TargetValueGenerator: valueGenerator,
		}

		population[i] = RandProgoBact(programMemorySize, program, energy)
	}

	environment := SimpleEnvironment{
		Rules: []Rules{
			defaultSumulationRules,
			defaultReapRules,
			defaultEloRatingsRules,
			//defaultAggressiveSelectionRules,
			defaultEloSelectionRules,
			defaultReapRules,
			defaultAggressiveReproduceRules(programInstructionVariants),
			defaultBacterialRules,
			//defaultMutateRules(programInstructionVariants),
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
				&population, validator, stabilityCheckLength)
		}

		fmt.Printf("BEST:\n%s", getBest(population))

		if validated {
			break
		}

		tick++
	}

	return population, tick
}

func validate(
	valueGenerator ValueGenerator,
	population *Population, validator PopulationValidator, ticks int,
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
		environment.Simulate(population)

		ticks--
	}

	return validator.Validate("validate", *population)
}

func getBest(population Population) Creature {
	sort.Sort(ByEnergyReverse(population))

	return population[0]
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
	validator := MedianErrorValidator{
		Threshold: 0.1,
	}

	_, _ = converge(
		&LinearGenerator{
			A: 32123.0,
			B: 1656789.0,
		},
		validator,
		[]RandInstructionVariant{},
		10,
	)
}

func TestCanEstimateSinFunction(t *testing.T) {
	validator := MedianErrorValidator{
		Threshold: 0.001,
	}

	_, _ = converge(
		&SinGenerator{
			Variation: 0.5,
			Amplitude: 67.0,
			Period:    1.0,
			History:   make([]float64, 1000),
		},
		validator,
		[]RandInstructionVariant{},
		50,
	)
}

func TestCanEstimateCeilFunction(t *testing.T) {
	validator := MedianErrorValidator{
		Threshold: 0.001,
	}

	_, _ = converge(
		&CeilGenerator{},
		validator,
		[]RandInstructionVariant{},
		50,
	)
}
