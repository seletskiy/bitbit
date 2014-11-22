package main

import "math/rand"

func CellFission(population []Creature) []Creature {
	for _, c := range population {
		bacteria := c.(Bacteria)
		if rand.Float64() > 0.98 {
			newBacteria := bacteria.Reproduce()

			if newBacteria == nil {
				continue
			}

			population = append(population, newBacteria)
		}
	}

	return population
}
