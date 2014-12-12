package main

import (
	"fmt"
	"math"
)

type PotentialEnergy interface {
	TransferPotential(target ReproductiveEnergy)
}

type ReproductiveEnergy struct {
	Potential int
	Dead      bool
}

func (energy *ReproductiveEnergy) GetFloat64() float64 {
	if energy.Dead {
		return 0.0
	} else {
		return float64(energy.Potential)
	}
}

func (energy *ReproductiveEnergy) Free() {
	energy.Dead = true
}

func (energy *ReproductiveEnergy) TransferTo(target Energy) {
	energy.TransferPotential(target.(*ReproductiveEnergy))
}

func (energy *ReproductiveEnergy) TransferPotential(target *ReproductiveEnergy) {
	target.Potential += energy.Potential
	energy.Potential = 0
	energy.Dead = true
}

func (energy *ReproductiveEnergy) Scatter(n int) []Energy {
	amount := int(math.Ceil(float64(energy.Potential) / float64(n)))
	scattered := []Energy{}
	for i := 0; i < n; i++ {
		transfer := amount
		if energy.Potential < amount {
			transfer = energy.Potential
		}

		if transfer <= 0 {
			break
		}

		scattered = append(scattered, &ReproductiveEnergy{
			Potential: transfer,
		})

		energy.Potential -= transfer
	}

	return scattered
}

func (energy *ReproductiveEnergy) Split() Energy {
	return nil
}

func (energy *ReproductiveEnergy) Void() bool {
	return energy.Dead
}

func (energy *ReproductiveEnergy) Simulate() {
}

func (energy *ReproductiveEnergy) String() string {
	return fmt.Sprintf("can split to %d; dead: %v",
		energy.Potential-1,
		energy.Dead,
	)
}
