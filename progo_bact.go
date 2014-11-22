package main

import (
	"fmt"
	"log"
	"math/rand"
	"regexp"
)

type ProgoBact struct {
	*SimpleBacteria
	State *ProgramState
}

func (bacteria *ProgoBact) String() string {
	plasmids := make([]string, 0)
	for i, program := range bacteria.GetPlasmids() {
		plasmids = append(plasmids,
			fmt.Sprintf("  %3d plasmid:\n%s", i+1, program))
	}

	chromosomeString := fmt.Sprint(bacteria.Chromosome)

	ipsRegexp := regexp.MustCompile(fmt.Sprintf("(?m)^%03d", bacteria.State.IPS))

	result := ""
	result += fmt.Sprintf(
		"BACTERIA %4.1fE (%p):\nState:\n%s\nChromosome:\n%s\n",
		bacteria.Energy,
		bacteria.SimpleBacteria,
		bacteria.State,
		ipsRegexp.ReplaceAllString(chromosomeString, ">>>"),
	)

	result += "Plasmids:"
	for i, plasmid := range bacteria.Plasmids {
		result += fmt.Sprintf("\nN%d:\n%s", i, plasmid)
	}

	if len(bacteria.Plasmids) > 0 {
		result += "\n"
	}

	return result
}

func (bacteria *ProgoBact) Simulate() {
	bacteria.Age++
	bacteria.Chromosome.DNA.(*ProgDNA).Eval(bacteria.State)
	bacteria.Energy = bacteria.State.ExternalData.(*DataStorage).FunValue
}

func (progobact *ProgoBact) Reproduce() Bacteria {
	memSize := progobact.State.Memory.GetSize()
	mem := NewProgramMemory(memSize)

	data := progobact.State.ExternalData.(*DataStorage)
	data.FunValue /= 2

	newBacteria := &ProgoBact{
		progobact.SimpleBacteria.Reproduce().(*SimpleBacteria),
		&ProgramState{
			IPS:    0,
			Memory: mem,
			ExternalData: &DataStorage{
				FunValue: data.FunValue,
			},
		},
	}

	newDna := newBacteria.GetChromosome().(*SimpleChromosome).DNA.(*ProgDNA)

	for i, cp := range *newDna.Program {
		if rand.Float64() > 0.95 {
			(*newDna.Program)[i].Instruction = RandProgramInstruction(memSize)
			log.Printf("mutation at %d: '%s' -> '%s'", i, cp.Instruction,
				(*newDna.Program)[i].Instruction)
		}
	}

	return newBacteria
}

func (bacteria *ProgoBact) Died() bool {
	return bacteria.SimpleBacteria.Died() || bacteria.State.Crashed
}
