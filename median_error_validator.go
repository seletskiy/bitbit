package main

import (
	"fmt"
	"math"
	"sort"
)

type MedianErrorValidator struct {
	Threshold float64
	LogPrefix string
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
		"%savg err: %10.4g med: %10.4f min: %10.4f (%3d/%3d)\n",
		validator.LogPrefix,
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
