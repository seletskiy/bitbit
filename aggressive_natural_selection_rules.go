package main

import (
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
	//maxValue := selected[length-1].GetEnergy().GetFloat64()

	//sorted := *population
	//sort.Sort(ByEnergy(sorted))

	logger.Log(Debug, "SELECTION: killing percentile: %f", percentileValue)

	for _, creature := range *population {
		energy := creature.GetEnergy().GetFloat64()
		if creature.GetAge() >= rules.MinAge {
			continue
		}

		if energy >= percentileValue {
			continue
		}

		logger.Log(Debug, "SELECTION: CREATURE<%p> child is killed (%f <  %f)",
			creature, energy, percentileValue,
		)

		creature.Kill()
	}

	minAlive := int(math.Max(
		1.0,
		float64(rules.BasePopulationSize)*(1-rules.DiePercentile),
	))

	for creatureIndex, creature := range selected {
		energy := creature.GetEnergy().GetFloat64()

		if creatureIndex > len(selected)-minAlive-1 {
			break
		}

		logger.Log(Debug, "SELECTION: CREATURE<%p> adult is killed (%f <= %f)",
			creature, energy, percentileValue,
		)

		creature.Kill()
	}

	//matureIndex := 0
	//for _, creature := range sorted {
	//    energy := creature.GetEnergy().GetFloat64()
	//    if creature.GetAge() >= rules.MinAge {
	//        if matureIndex >= deadIndex {
	//            if energy >= maxValue {
	//                break
	//            }
	//        }
	//        matureIndex++
	//    }

	//    if energy > percentileValue {
	//        break
	//    }

	//    logger.Log(Debug, "SELECTION: CREATURE<%p> is killed (%f <= %f)",
	//        creature, creature.GetEnergy().GetFloat64(), percentileValue,
	//    )

	//    creature.Kill()
	//}
}
