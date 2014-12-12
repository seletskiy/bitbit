package main

import (
	"fmt"
	"math"
)

type ValueSetterEnergy interface {
	SetValue(float64)
}

type ErrorGetterEnergy interface {
	GetError() float64
	GetLastError() float64
}

type ErrorBasedEnergy struct {
	*ReproductiveEnergy
	TargetValue  float64
	CurrentValue float64
	LastError    float64
	Set          bool
}

func (origin *ErrorBasedEnergy) GetError() float64 {
	return origin.TargetValue - origin.CurrentValue
}

func (origin *ErrorBasedEnergy) GetLastError() float64 {
	return origin.LastError
}

func (origin *ErrorBasedEnergy) SetValue(value float64) {
	origin.Set = true
	origin.CurrentValue = value
}

func (origin *ErrorBasedEnergy) SetTargetValue(value float64) {
	origin.TargetValue = value
}

func (origin ErrorBasedEnergy) GetFloat64() float64 {
	return 1 / math.Abs(origin.GetError())
}

func (origin *ErrorBasedEnergy) Scatter(n int) []Energy {
	scattered := []Energy{}
	for _, part := range origin.ReproductiveEnergy.Scatter(n) {
		scattered = append(scattered,
			&ErrorBasedEnergy{
				ReproductiveEnergy: part.(*ReproductiveEnergy),
				TargetValue:        origin.TargetValue,
				LastError:          origin.LastError,
			},
		)
	}

	return scattered
}

func (origin *ErrorBasedEnergy) TransferTo(energy Energy) {
	origin.ReproductiveEnergy.TransferTo(
		energy.(*ErrorBasedEnergy).ReproductiveEnergy,
	)
}

func (origin *ErrorBasedEnergy) Split() Energy {
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

func (origin ErrorBasedEnergy) String() string {
	return fmt.Sprintf("reproduce potential: %d; error: %f; score: %f; last err: %f",
		origin.Potential, origin.GetError(), origin.GetFloat64(),
		origin.LastError,
	)
}

func (origin *ErrorBasedEnergy) Simulate() {
	if origin.Set {
		origin.LastError = origin.GetError()
	}

	origin.CurrentValue = 0
}
