package main

import (
	"math"
)

type EloSelectionRules struct {
	DiePercentile      float64
	MinAge             int
	BasePopulationSize int
}

func (rules EloSelectionRules) Apply(
	population *Population,
) {
	adults := []Creature{}
	children := []Creature{}

	for _, creature := range *population {
		if creature.GetAge() < rules.MinAge {
			children = append(children, creature)
		} else {
			adults = append(adults, creature)
		}
	}

	if len(adults) == 0 {
		Log(Debug, "SELECTION: population too young")
		return
	}

	//deadIndex := int(float64(len(*population)) * (1 - rules.DiePercentile))
	deadIndex := int(float64(rules.BasePopulationSize) * (1 - rules.DiePercentile))

	percentileValue := (*population)[deadIndex].GetEnergy().(EloBasedEnergy).GetEloScore()

	Log(Debug,
		"SELECTION: killing percentile: %d <%p>",
		percentileValue,
		(*population)[deadIndex],
	)

	minAliveCount := int(math.Max(
		1.0,
		float64(rules.BasePopulationSize)*(1-rules.DiePercentile),
	))

	for creatureIndex, creature := range adults {
		energy := creature.GetEnergy().(EloBasedEnergy)

		if energy.GetWinsInRow() > 0 {
			continue
		}

		score := energy.GetEloScore()

		if creatureIndex > len(adults)-minAliveCount-1 {
			break
		}

		if score >= percentileValue {
			continue
		}

		creature.Kill()

		Log(Debug, "SELECTION: CREATURE<%p> adult is killed (%d < %d)",
			creature, score, percentileValue,
		)
	}
}
