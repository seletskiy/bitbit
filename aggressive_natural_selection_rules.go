package main

import "sort"

type AggressiveNaturalSelectionRules struct {
	DiePercentile float64
}

func (rules AggressiveNaturalSelectionRules) Apply(
	population *Population,
) {
	sort.Sort(ByEnergy(*population))

	dieIndex := int(float64(len(*population)) * rules.DiePercentile)

	for _, creature := range (*population)[:dieIndex] {
		creature.Kill()
	}
}
