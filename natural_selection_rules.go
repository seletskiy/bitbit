package main

import "math/rand"

type NaturalSelectionRules struct {
	KillPossibility   float64
	MinPopulationSize int
	MinKillAge        int
}

func (environment NaturalSelectionRules) Apply(
	population *Population,
) {
	if environment.MinPopulationSize > len(*population) {
		return
	}

	for _, creature := range *population {
		opponent := (*population)[rand.Intn(len(*population))]

		if creature.GetAge() < environment.MinKillAge {
			continue
		}

		if opponent == creature {
			continue
		}

		if opponent.Died() {
			continue
		}

		if creature.Died() {
			continue
		}

		if creature.CanKill(opponent) {
			r := rand.Float64()
			if r < environment.KillPossibility {
				Log(Debug,
					"CONSUME: CREATURE<%p> eats CREATURE<%p>",
					creature,
					opponent,
				)

				creature.Consume(opponent)
			}
		}
	}
}
