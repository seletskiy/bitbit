package main

import (
	"fmt"
	"math/rand"
)

type Bacteria interface {
	Creature
	GetPlasmids() []*Plasmid
	SetPlasmids([]*Plasmid)
	Reproduce() Bacteria
}

type SimpleBacteria struct {
	Energy     float64
	Age        int
	Chromosome *SimpleChromosome
	Plasmids   []*Plasmid
}

func (b *SimpleBacteria) GetChromosome() Chromosome {
	return b.Chromosome
}

func (b *SimpleBacteria) GetAge() int {
	return b.Age
}

func (bacteria *SimpleBacteria) GetPlasmids() []*Plasmid {
	return bacteria.Plasmids
}

func (bacteria *SimpleBacteria) SetPlasmids(plasmids []*Plasmid) {
	bacteria.Plasmids = plasmids
}

func (bacteria *SimpleBacteria) Simulate() {
	// noop, must be overriden in child struct
}

func (bacteria *SimpleBacteria) String() string {
	// noop, must be overriden in child struct
	return fmt.Sprintf("%s", bacteria)
}

func (bacteria *SimpleBacteria) Reproduce() Bacteria {
	bacteria.Energy /= 2.0

	newBacteria := &SimpleBacteria{
		Energy:     bacteria.Energy,
		Age:        0,
		Plasmids:   make([]*Plasmid, 0),
		Chromosome: bacteria.Chromosome.Clone().(*SimpleChromosome),
	}

	keepPlasmids := make([]*Plasmid, 0)
	for _, plasmid := range bacteria.GetPlasmids() {
		if rand.Float64() < 0.8 {
			continue
		}

		plasmidCopy := *plasmid

		newBacteria.Plasmids = append(newBacteria.Plasmids, &plasmidCopy)

		if rand.Float64() > 0.5 {
			keepPlasmids = append(keepPlasmids, plasmid)
		}
	}

	bacteria.SetPlasmids(keepPlasmids)

	return newBacteria
}

func (bacteria *SimpleBacteria) Died() bool {
	return bacteria.Energy <= 0
}

func (bacteria SimpleBacteria) Kill(opponent Creature) bool {
	if bacteria.GetEnergy() > opponent.GetEnergy() {
		return true
	}

	return false
}

func (bacteria *SimpleBacteria) Consume(opponent Creature) {
	bacteria.Energy += opponent.GetEnergy()
	opponent.SetEnergy(0)
}

func (bacteria *SimpleBacteria) SetEnergy(energy float64) {
	bacteria.Energy = energy
}

func (bacteria *SimpleBacteria) GetEnergy() float64 {
	return bacteria.Energy
}
