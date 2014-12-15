package main

import (
	"log"
	"math"
	"sort"
)

type AggressiveNaturalSelectionRules struct {
	DiePercentile      float64
	MinAge             int
	BasePopulationSize int
}

func (rules AggressiveNaturalSelectionRules) Apply(
	population *Population,
) {
	matures := Population{}
	children := Population{}

	for _, creature := range *population {
		if creature.GetAge() < rules.MinAge {
			children = append(children, creature)
			continue
		}

		matures = append(matures, creature)
	}

	if len(matures) == 0 {
		Log(Debug, "SELECTION: population too young")
		return
		//matures = *population
		//children = Population{}
	}

	sort.Sort(ByEnergy(matures))

	length := len(matures)
	deadIndex := int(math.Min(
		float64(length)*rules.DiePercentile,
		float64(length-1),
	))
	percentileValue := matures[deadIndex].GetEnergy().GetFloat64()

	Log(Debug,
		"SELECTION: killing percentile: %f <%p>",
		percentileValue,
		matures[deadIndex],
	)

	for _, creature := range children {
		energy := creature.GetEnergy().GetFloat64()
		if creature.GetAge() >= rules.MinAge {
			continue
		}

		if energy >= percentileValue {
			continue
		}

		Log(Debug, "SELECTION: CREATURE<%p> child is killed (%f < %f)",
			creature, energy, percentileValue,
		)

		creature.Kill()
	}

	minAlive := int(math.Max(
		1.0,
		float64(rules.BasePopulationSize)*(1-rules.DiePercentile),
	))

	log.Printf("SELECTION: matures %d", len(matures))
	log.Printf("SELECTION: children %d", len(children))

	for creatureIndex, creature := range matures {
		energy := creature.GetEnergy().GetFloat64()

		if creatureIndex > len(matures)-minAlive-1 {
			break
		}

		Log(Debug, "SELECTION: CREATURE<%p> adult is killed (%f <= %f)",
			creature, energy, percentileValue,
		)

		creature.Kill()
	}
}
