package main

import (
	"fmt"
	"math"
	"sort"
	"testing"
)

type TestEnergy struct {
	*ReproductiveEnergy
	TargetValue  float64
	CurrentValue float64
	LastError    float64
	Set          bool
}

func (origin *TestEnergy) GetError() float64 {
	return origin.TargetValue - origin.CurrentValue
}

func (origin *TestEnergy) SetValue(value float64) {
	origin.Set = true
	origin.CurrentValue = value
}

func (origin TestEnergy) GetFloat64() float64 {
	return 1 / math.Abs(origin.GetError())
}

func (origin *TestEnergy) Scatter(n int) []Energy {
	scattered := []Energy{}
	for _, part := range origin.ReproductiveEnergy.Scatter(n) {
		scattered = append(scattered,
			&TestEnergy{
				ReproductiveEnergy: part,
				TargetValue:        origin.TargetValue,
				LastError:          origin.LastError,
			},
		)
	}

	return scattered
}

func (origin *TestEnergy) Split() Energy {
	scattered := origin.Scatter(2)

	if len(scattered) < 1 {
		return nil
	}

	scattered[0].TransferTo(origin)

	if len(scattered) < 2 {
		return nil
	}

	return scattered[1]
}

func (origin TestEnergy) String() string {
	return fmt.Sprintf("reproduce potential: %d; error: %f; score: %f; last err: %f",
		origin.Potential, origin.GetError(), origin.GetFloat64(),
		origin.LastError,
	)
}

func (origin *TestEnergy) Simulate() {
	if origin.Set {
		origin.LastError = origin.GetError()
	}

	origin.CurrentValue = 0
}

type TestInstruction struct {
	In ProgramArgReference
}

func (instruction *TestInstruction) Eval(state *ProgramState) error {
	state.ExternalData.(*TestEnergy).SetValue(
		instruction.In.GetValue(state).GetFloat64(),
	)
	return nil
}

func (instruction *TestInstruction) String() string {
	return fmt.Sprintf("TEST %s", instruction.In)
}

func (instruction *TestInstruction) GetArgsCount() int {
	return 1
}

func (instruction *TestInstruction) GetArg(index int) ProgramInstructionArg {
	return instruction.In
}

func (instruction *TestInstruction) SetArg(
	index int, arg ProgramInstructionArg,
) {
	instruction.In = arg.(ProgramArgReference)
}

func (instruction *TestInstruction) Init() {
	instruction.In = FloatRegister(0)
}

func (instruction *TestInstruction) Copy() ProgramInstruction {
	instructionCopy := *instruction
	return &instructionCopy
}

type LastErrorInstruction struct {
	Out ProgramArgRegister
}

func (instruction *LastErrorInstruction) Eval(state *ProgramState) error {
	instruction.Out.SetValue(state,
		FloatValue(state.ExternalData.(*TestEnergy).LastError),
	)

	return nil
}

func (instruction *LastErrorInstruction) String() string {
	return fmt.Sprintf("LASTERR %s", instruction.Out)
}

func (instruction *LastErrorInstruction) GetArgsCount() int {
	return 1
}

func (instruction *LastErrorInstruction) GetArg(index int) ProgramInstructionArg {
	return instruction.Out
}

func (instruction *LastErrorInstruction) SetArg(
	index int, arg ProgramInstructionArg,
) {
	instruction.Out = arg.(ProgramArgRegister)
}

func (instruction *LastErrorInstruction) Init() {
	instruction.Out = FloatRegister(0)
}

func (instruction *LastErrorInstruction) Copy() ProgramInstruction {
	instructionCopy := *instruction
	return &instructionCopy
}

func TestCanConvergeToNumberValue(t *testing.T) {
	targetValue := 543612.0
	targetError := 0.01

	programLayout := RandProgramLayout(programLength)

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

	population := make(Population, initialPopulationSize)
	for i := 0; i < initialPopulationSize; i++ {
		program := RandProgram(
			programLayout,
			programReferenceProbability,
			defaultVarianceGenerator,
			programMemorySize,
			programInstructionVariants,
		)

		energy := &TestEnergy{
			TargetValue: targetValue,
			Potential:   1,
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
		environment.Simulate(&population)

		populationErrors := []float64{}
		totalError := 0.0
		minError := math.Abs(targetValue)
		mature := 0
		for _, creature := range population {
			if creature.GetAge() <= 0 {
				continue
			}

			mature++

			energy := creature.GetEnergy().(*TestEnergy)

			creatureError := math.Abs(energy.GetError())

			totalError += creatureError
			minError = math.Min(minError, creatureError)
			populationErrors = append(populationErrors, creatureError)
		}

		sort.Float64s(populationErrors)

		medianError := populationErrors[len(populationErrors)/2]

		fmt.Printf(
			"[%4d]: avg err: %10.4f med: %10.4f min: %10.4f (%3d mature/%3d total)\n",
			tick,
			totalError/float64(len(populationErrors)),
			medianError,
			minError,
			mature,
			len(population),
		)

		if medianError <= targetError {
			break
		}

		tick++
	}

	sort.Sort(ByEnergy(population))

	fmt.Printf("TICKS: %d\n", tick)
	fmt.Printf("BEST:\n%s", population[len(population)-1])
}
