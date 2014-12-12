package main

import (
	"math"
	"sort"
)

type AggressiveNaturalSelectionRules struct {
	DiePercentile float64
	MinAge        int
}

func (rules AggressiveNaturalSelectionRules) Apply(
	population *Population,
) {
	selected := []Creature{}

	for _, creature := range *population {
		if creature.GetAge() < rules.MinAge {
			continue
		}

		selected = append(selected, creature)
	}

	if len(selected) == 0 {
		logger.Log(Debug, "SELECTION: population too young")
		return
	}

	sort.Sort(ByEnergy(selected))

	length := len(selected)
	deadIndex := int(math.Min(
		float64(length)*rules.DiePercentile,
		float64(length-1),
	))

	percentileValue := selected[deadIndex].GetEnergy().GetFloat64()
	maxValue := selected[length-1].GetEnergy().GetFloat64()

	logger.Log(Debug, "SELECTION: killing percentile: %f", percentileValue)

	for creatureIndex, creature := range selected {
		energy := creature.GetEnergy().GetFloat64()
		if creatureIndex >= deadIndex {
			if energy >= maxValue {
				continue
			}

			if energy > percentileValue {
				continue
			}
		}

		logger.Log(Debug, "SELECTION: CREATURE<%p> is killed (%f <= %f)",
			creature, creature.GetEnergy().GetFloat64(), percentileValue,
		)

		creature.Kill()
	}
}

type ByEnergy []Creature

func (creatures ByEnergy) Len() int {
	return len(creatures)
}

func (creatures ByEnergy) Swap(i, j int) {
	creatures[i], creatures[j] = creatures[j], creatures[i]
}

func (creatures ByEnergy) Less(i, j int) bool {
	if creatures[i].GetEnergy().Void() {
		return true
	}

	if creatures[j].GetEnergy().Void() {
		return false
	}

	return creatures[i].GetEnergy().GetFloat64() <
		creatures[j].GetEnergy().GetFloat64()
}

type ByAge []Creature

func (creatures ByAge) Len() int {
	return len(creatures)
}

func (creatures ByAge) Swap(i, j int) {
	creatures[i], creatures[j] = creatures[j], creatures[i]
}

func (creatures ByAge) Less(i, j int) bool {
	return creatures[i].GetAge() < creatures[j].GetAge()
}
