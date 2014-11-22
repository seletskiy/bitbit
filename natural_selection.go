package main

import "math/rand"

type NaturalSelectionEnvironment struct {
	SimpleEnvironment
	KillPossibility float64
}

func (environment *NaturalSelectionEnvironment) Simulate(
	tick int,
	population Population,
) Population {
	population = environment.SimpleEnvironment.Simulate(tick, population)

	for _, creature := range population {
		opponent := population[rand.Intn(len(population))]

		if creature.Kill(opponent) {
			if rand.Float64() < environment.KillPossibility {
				creature.Consume(opponent)
			}
		}
	}

	population = environment.Reap(tick, population)

	return population
}
