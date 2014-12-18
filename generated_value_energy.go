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
}

type GeneratedValueEnergy struct {
	*EloEnergy
	TargetValueGenerator ValueGenerator
	TargetValue          float64
	ZeroThreshold        float64
	CurrentValue         float64
	Age                  int
}

func (origin *GeneratedValueEnergy) GetCurrentError() float64 {
	return math.Abs(origin.TargetValue - origin.CurrentValue)
}

func (origin GeneratedValueEnergy) GetFloat64() float64 {
	return -origin.GetCurrentError()
}

func (origin GeneratedValueEnergy) Void() bool {
	if math.IsNaN(origin.GetCurrentError()) {
		return true
	}

	if origin.EloEnergy.Void() {
		return true
	} else {
		return math.Abs(origin.GetFloat64()) <= origin.ZeroThreshold
	}
}

func (origin *GeneratedValueEnergy) Scatter(n int) []Energy {
	scattered := []Energy{}
	for _, base := range origin.EloEnergy.Scatter(n) {
		scattered = append(scattered,
			&GeneratedValueEnergy{
				EloEnergy:            base.(*EloEnergy),
				TargetValueGenerator: origin.TargetValueGenerator,
				TargetValue:          origin.TargetValue,
				ZeroThreshold:        origin.ZeroThreshold,
			},
		)
	}

	return scattered
}

func (origin *GeneratedValueEnergy) TransferTo(energy Energy) {
	origin.EloEnergy.TransferTo(energy.(*GeneratedValueEnergy).EloEnergy)
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
	return fmt.Sprintf("error: %f; energy: %f; elo: %d;\ngenerator: %s",
		origin.GetCurrentError(),
		origin.GetFloat64(),
		origin.GetEloScore(),
		origin.TargetValueGenerator,
	)
}

func (origin *GeneratedValueEnergy) Simulate() {
	origin.CurrentValue = 0
	origin.TargetValue = origin.TargetValueGenerator.Generate()
	origin.Age++
}

func (origin *GeneratedValueEnergy) Free() {
	origin.EloEnergy.Free()
}

func (origin *GeneratedValueEnergy) SetCurrentValue(value float64) {
	origin.CurrentValue = value
}

func (origin *GeneratedValueEnergy) GetData(
	tableIndex, cellIndex int,
) float64 {
	return origin.TargetValueGenerator.GetData(tableIndex, cellIndex)
}
