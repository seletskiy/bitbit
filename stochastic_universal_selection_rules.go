package main

import (
	"log"
	"math"
	"math/rand"
	"sort"
)

type StochasticUnversalSelectionRules struct {
	DiePercentile float64
}

func (rules StochasticUnversalSelectionRules) Apply(
	population *Population,
) {
	sort.Sort(ByEnergyReverse(*population))

	totalFitness := 0.0
	for _, creature := range *population {
		if math.IsInf(creature.GetEnergy().GetFloat64(), 0) {
			continue
		}

		totalFitness += creature.GetEnergy().GetFloat64()
	}

	numberToKeep := int(float64(len(*population)) * (1 - rules.DiePercentile))
	pointersDistance := totalFitness / float64(numberToKeep)

	startPoint := rand.Float64() * pointersDistance

	pointers := make([]float64, numberToKeep-1)
	for i, _ := range pointers {
		pointers[i] = startPoint + float64(i)*pointersDistance
	}

	log.Printf("%#v", pointers)

	alive := rules.selectByRoulette(population, pointers)
	for _, creature := range alive {
		log.Printf("SUS: CREATURE<%p> energy: %f", creature,
			creature.GetEnergy().GetFloat64(),
		)
	}

	for _, creature := range *population {
		kept := false
		for _, lucky := range alive {
			if creature == lucky {
				kept = true
				break
			}
		}

		if kept {
			continue
		}

		creature.Kill()
	}
}

func (rules StochasticUnversalSelectionRules) selectByRoulette(
	population *Population, pointers []float64,
) []Creature {
	keep := []Creature{}

	index := 0
	fitnessSum := 0.0
	for _, point := range pointers {
		log.Printf("POINT %#v", point)
		for {
			energy := (*population)[index].GetEnergy().GetFloat64()
			fitnessSum += energy
			if fitnessSum >= point {
				break
			}

			log.Printf("%p %f %f", (*population)[index], 1/energy, fitnessSum)
			index++
		}

		keep = append(keep, (*population)[index])
	}

	return keep
}
