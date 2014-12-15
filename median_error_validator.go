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
		if creature.GetAge() <= 1 {
			continue
		}

		absError := math.Abs(
			1 / creature.GetEnergy().GetFloat64(),
		)

		populationErrors = append(populationErrors, absError)
	}

	if len(populationErrors) == 0 {
		return false
	}

	sort.Float64s(populationErrors)

	minError := populationErrors[0]
	medianError := populationErrors[len(populationErrors)/2]
	totalError := 0.0
	for _, val := range populationErrors {
		totalError += val
	}

	sorted := make(Population, len(population))
	copy(sorted, population)
	sort.Sort(ByEnergy(sorted))

	fmt.Printf(
		"[%10s] avg err: %10.4g med: %10.5g min: %10.5g (%4d/%4d) <%p>\n",
		step,
		totalError/float64(len(populationErrors)),
		medianError,
		minError,
		len(populationErrors),
		len(population),
		sorted[len(sorted)-1],
	)

	if minError <= validator.Threshold {
		return true
	} else {
		return false
	}
}
