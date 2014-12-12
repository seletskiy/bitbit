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

type ErrorGetterEnergy interface {
	GetError() float64
}

type GeneratedValueEnergy struct {
	Base                 Energy
	TargetValueGenerator ValueGenerator
	TargetValue          float64
	ConsiderZero         float64
	CurrentValue         float64
	Set                  bool
	TotalError           float64
}

func (origin *GeneratedValueEnergy) GetError() float64 {
	return origin.TotalError + math.Abs(
		origin.TargetValue-origin.CurrentValue,
	)
}

func (origin GeneratedValueEnergy) GetFloat64() float64 {
	return 1 / origin.GetError()
}

func (origin GeneratedValueEnergy) Void() bool {
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
				ConsiderZero:         origin.ConsiderZero,
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

	scattered[0].TransferTo(origin)

	if len(scattered) < 2 {
		return nil
	}

	return scattered[1]
}

func (origin GeneratedValueEnergy) String() string {
	return fmt.Sprintf("error: %f; score: %f; base: %s\ngenerator: %s",
		origin.GetError(),
		origin.GetFloat64(),
		origin.Base,
		origin.TargetValueGenerator,
	)
}

func (origin *GeneratedValueEnergy) Simulate() {
	if origin.Set {
		origin.TotalError = origin.GetError()
	}

	origin.CurrentValue = 0
	origin.TargetValue = origin.TargetValueGenerator.Generate()
	origin.Set = true
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
