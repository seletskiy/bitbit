package main

import "sort"

type AggressiveNaturalSelectionRules struct {
	DiePercentile float64
}

func (rules AggressiveNaturalSelectionRules) Apply(
	population *Population,
) {
	sortedPopulation := *population
	sort.Sort(ByEnergy(sortedPopulation))

	populationSize := len(sortedPopulation)
	deadIndex := int(float64(populationSize) * rules.DiePercentile)
	percentileValue := sortedPopulation[deadIndex].GetEnergy().GetFloat64()
	maxValue := sortedPopulation[populationSize-1].GetEnergy().GetFloat64()

	logger.Log(Debug, "SELECTION: killing percentile: %f", percentileValue)

	for creatureIndex, creature := range sortedPopulation {
		energy := creature.GetEnergy().GetFloat64()
		if creatureIndex > deadIndex {
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
