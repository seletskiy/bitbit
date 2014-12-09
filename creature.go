package main

import (
	"fmt"
	"strings"
)

type Creature interface {
	GetEnergy() Energy
	SetEnergy(Energy)
	CanKill(Creature) bool
	// thx @a.baranov
	Consume(Creature)
	Kill()
	Died() bool
	GetChromosome() Chromosome
	GetAge() int
	Simulate()
	GetParents() []Creature
	Reproduce() Creature
}

type Population []Creature

func (population Population) String() string {
	result := make([]string, 0)

	for _, creature := range population {
		result = append(result, fmt.Sprint(creature))
	}

	return strings.Join(result, "\n\n")
}
