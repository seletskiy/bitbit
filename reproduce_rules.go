package main

import "math/rand"

type ReproduceRules struct {
	ReproduceProbability float64
	MinReproduceAge      int
}

func (rules ReproduceRules) Apply(population *Population) {
	for _, creature := range *population {
		if creature.Died() {
			continue
		}

		if creature.GetAge() < rules.MinReproduceAge {
			continue
		}

		if rand.Float64() > rules.ReproduceProbability {
			continue
		}

		child := creature.Reproduce()
		if child == nil {
			Log(Debug, "CREATURE<%p> reproduce to nil", creature)
			continue
		}

		Log(Debug, "CREATURE<%p> reproduce to CREATURE<%p>", creature, child)

		*population = append(*population, child)
	}
}
