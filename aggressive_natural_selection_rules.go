package main

import (
	"log"
	"math"
)

type AggressiveNaturalSelectionRules struct {
	DiePercentile      float64
	MinAge             int
	BasePopulationSize int
}

func (rules AggressiveNaturalSelectionRules) Apply(
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

	log.Printf("SELECTION: adults %d", len(adults))
	log.Printf("SELECTION: children %d", len(children))

	if len(adults) == 0 {
		Log(Debug, "SELECTION: population too young")
		return
	}

	adultsCount := len(adults)
	deadIndex := int(math.Min(
		float64(adultsCount)*rules.DiePercentile,
		float64(adultsCount-1),
	))

	percentileValue := adults[deadIndex].GetEnergy().(EloBasedEnergy).GetEloScore()

	Log(Debug,
		"SELECTION: killing percentile: %d <%p>",
		percentileValue,
		adults[deadIndex],
	)

	//for _, creature := range children {
	//    score := creature.GetEnergy().(EloBasedEnergy).GetEloScore()

	//    if creature.GetAge() <= rules.MinAge {
	//        continue
	//    }

	//    if score >= percentileValue {
	//        continue
	//    }

	//    creature.Kill()

	//    Log(Debug, "SELECTION: CREATURE<%p> child is killed (%d < %d)",
	//        creature, score, percentileValue,
	//    )
	//}

	minAliveCount := int(math.Max(
		1.0,
		float64(rules.BasePopulationSize)*(1-rules.DiePercentile),
	))

	for creatureIndex, creature := range adults {
		score := creature.GetEnergy().(EloBasedEnergy).GetEloScore()

		if creatureIndex > len(adults)-minAliveCount-1 {
			break
		}

		if score > percentileValue {
			continue
		}

		creature.Kill()

		Log(Debug, "SELECTION: CREATURE<%p> adult is killed (%d < %d)",
			creature, score, percentileValue,
		)
	}
}
