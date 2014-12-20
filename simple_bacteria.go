package main

import "fmt"

type SimpleBacteria struct {
	Parents    []Creature
	Energy     Energy
	Age        int
	Chromosome Chromosome
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
	bacteria.Energy.Simulate()
	bacteria.Age++
}

func (bacteria *SimpleBacteria) String() string {
	return fmt.Sprintf("%s", bacteria)
}

func (bacteria *SimpleBacteria) Reproduce() Creature {
	newEnergy := bacteria.Energy.Split()
	if newEnergy == nil {
		return nil
	}

	return &SimpleBacteria{
		Energy:     newEnergy,
		Chromosome: bacteria.Chromosome.Clone(),
		//Parents:    []Creature{bacteria},
	}
}

func (bacteria *SimpleBacteria) Died() bool {
	return bacteria.Energy == nil || bacteria.Energy.Void()
}

func (bacteria SimpleBacteria) CanKill(opponent Creature) bool {
	return bacteria.Energy.GetFloat64() > opponent.GetEnergy().GetFloat64()
}

func (bacteria *SimpleBacteria) Consume(opponent Creature) {
	opponent.GetEnergy().TransferTo(bacteria.Energy)
}

func (bacteria *SimpleBacteria) Kill() {
	bacteria.GetEnergy().Free()
}

func (bacteria *SimpleBacteria) GetEnergy() Energy {
	return bacteria.Energy
}

func (bacteria *SimpleBacteria) GetParents() []Creature {
	return bacteria.Parents
}

func (bacteria *SimpleBacteria) SetEnergy(energy Energy) {
	bacteria.Energy = energy
}
