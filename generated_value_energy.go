package main

import (
	"fmt"
	"math"
)

type ValueGenerator interface {
	Generate() float64
	Simulate()
	GetData(tableIndex, cellIndex int) float64
}

type ErrorBasedEnergy interface {
	GetCurrentError() float64
	GetTotalError() float64
	GetAvgError() float64
}

type GeneratedValueEnergy struct {
	Base                 Energy
	TargetValueGenerator ValueGenerator
	TargetValue          float64
	ConsiderZero         float64
	CurrentValue         float64
	Age                  int
	Errors               []float64
	AvgPeriod            int
}

func (origin *GeneratedValueEnergy) GetCurrentError() float64 {
	return math.Abs(origin.TargetValue - origin.CurrentValue)
}

func (origin *GeneratedValueEnergy) GetAvgError() float64 {
	return origin.GetTotalError() / float64(len(origin.Errors)+1)

}

func (origin *GeneratedValueEnergy) GetTotalError() float64 {
	total := origin.GetCurrentError()
	for _, value := range origin.Errors {
		total += value
	}

	return total
}

func (origin GeneratedValueEnergy) GetFloat64() float64 {
	return 1 / origin.GetAvgError()
}

func (origin GeneratedValueEnergy) Void() bool {
	if math.IsNaN(origin.GetCurrentError()) {
		return true
	}

	if origin.Base.Void() {
		return true
	} else {
		return math.Abs(origin.GetFloat64()) <= origin.ConsiderZero
	}
}

func (origin *GeneratedValueEnergy) Scatter(n int) []Energy {
	scattered := []Energy{}
	for _, base := range origin.Base.Scatter(n) {
		scattered = append(scattered,
			&GeneratedValueEnergy{
				Base:                 base,
				TargetValueGenerator: origin.TargetValueGenerator,
				TargetValue:          origin.TargetValue,
				ConsiderZero:         origin.ConsiderZero,
				AvgPeriod:            origin.AvgPeriod,
			},
		)
	}

	return scattered
}

func (origin *GeneratedValueEnergy) TransferTo(energy Energy) {
	origin.Base.TransferTo(energy.(*GeneratedValueEnergy).Base)
}

func (origin *GeneratedValueEnergy) Split() Energy {
	scattered := origin.Scatter(2)

	if len(scattered) < 1 {
		return nil
	}

	rest := scattered[0].(*GeneratedValueEnergy)

	rest.TransferTo(origin)

	if len(scattered) < 2 {
		return nil
	}

	return scattered[1]
}

func (origin GeneratedValueEnergy) String() string {
	return fmt.Sprintf("error: %f~%f; score: %f; base: %s\ngenerator: %s",
		origin.GetCurrentError(),
		origin.GetAvgError(),
		origin.GetFloat64(),
		origin.Base,
		origin.TargetValueGenerator,
	)
}

func (origin *GeneratedValueEnergy) Simulate() {
	if origin.Age > 0 {
		origin.Errors = append(origin.Errors, origin.GetCurrentError())
		errorsCount := len(origin.Errors)
		if errorsCount > origin.AvgPeriod {
			origin.Errors = origin.Errors[errorsCount-origin.AvgPeriod-1 : errorsCount-1]
		}
	}

	origin.CurrentValue = 0
	origin.TargetValue = origin.TargetValueGenerator.Generate()
	origin.Age++
}

func (origin *GeneratedValueEnergy) Free() {
	origin.Base.Free()
}

func (origin *GeneratedValueEnergy) SetCurrentValue(value float64) {
	origin.CurrentValue = value
}

func (origin *GeneratedValueEnergy) GetData(
	tableIndex, cellIndex int,
) float64 {
	return origin.TargetValueGenerator.GetData(tableIndex, cellIndex)
}
