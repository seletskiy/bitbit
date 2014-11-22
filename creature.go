package main

import (
	"fmt"
	"strings"
)

type Creature interface {
	GetEnergy() float64
	SetEnergy(float64)
	Kill(Creature) bool
	// thx @a.baranov
	Consume(Creature)
	Died() bool
	GetChromosome() Chromosome
	GetAge() int
	Simulate()
}

type Population []Creature

func (population Population) String() string {
	result := make([]string, 0)

	for _, creature := range population {
		result = append(result, fmt.Sprint(creature))
	}

	return strings.Join(result, "\n\n")
}
