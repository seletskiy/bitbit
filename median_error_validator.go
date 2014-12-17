package main

import (
	"fmt"
	"sort"
)

type MedianErrorValidator struct {
	Threshold float64
}

func (validator MedianErrorValidator) Validate(
	step string, population Population,
) bool {
	nonEggs := []Creature{}

	for _, creature := range population {
		if creature.GetAge() <= 0 {
			continue
		}

		nonEggs = append(nonEggs, creature)
	}

	if len(nonEggs) == 0 {
		return false
	}

	sort.Sort(ByMedianError(nonEggs))

	best := nonEggs[0].GetEnergy().(ErrorBasedEnergy)
	median := nonEggs[len(nonEggs)/2].GetEnergy().(ErrorBasedEnergy)
	totalError := 0.0
	for _, creature := range nonEggs {
		totalError += creature.GetEnergy().(ErrorBasedEnergy).GetMedianError()
	}

	fmt.Printf(
		"[%10s] avg err: %10.4g med: %10.5g min: %10.5g (%4d/%4d) <%p>\n",
		step,
		totalError/float64(len(nonEggs)),
		median.GetMedianError(),
		best.GetMedianError(),
		len(nonEggs),
		len(population),
		nonEggs[0],
	)

	if best.GetMedianError() <= validator.Threshold {
		return true
	} else {
		return false
	}
}
