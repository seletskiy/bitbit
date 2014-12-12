package main

import (
	"fmt"
	"math"
	"sort"
)

type MedianErrorValidator struct {
	Threshold float64
}

func (validator MedianErrorValidator) Validate(
	step string, population Population,
) bool {
	populationErrors := []float64{}
	for _, creature := range population {
		if creature.GetAge() <= 0 {
			continue
		}

		absError := math.Abs(
			creature.GetEnergy().(ErrorGetterEnergy).GetError(),
		)
		//log.Printf("aaa %p %10.5f", creature, absError)
		populationErrors = append(populationErrors, absError)
	}
	sort.Float64s(populationErrors)

	minError := populationErrors[0]
	medianError := populationErrors[len(populationErrors)/2]
	totalError := 0.0
	for _, val := range populationErrors {
		totalError += val
	}

	fmt.Printf(
		"[%10s] avg err: %10.4g med: %10.5f min: %10.5f (%4d/%4d)\n",
		step,
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
