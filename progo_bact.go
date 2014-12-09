package main

import (
	"fmt"
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

	ipsRegexp := regexp.MustCompile(
		fmt.Sprintf("(?m) %03d ", bacteria.State.IPS))

	result := ""
	result += fmt.Sprintf(
		"BACTERIA<%p> AGE %d DNA<%p>\nState:\n%s\nChromosome:\n%s\n",
		bacteria,
		bacteria.GetAge(),
		bacteria.GetChromosome().GetDNAs()[0],
		bacteria.State,
		ipsRegexp.ReplaceAllString(chromosomeString, " ==> "),
	)

	if len(bacteria.Plasmids) > 0 {
		result += "Plasmids:"
		for i, plasmid := range bacteria.Plasmids {
			if plasmid.Applied {
				result += fmt.Sprintf("\n#%d (offset %d):\n%s", i,
					plasmid.ReplaceIndex,
					plasmid,
				)
			} else {
				result += fmt.Sprintf("\n#%d:\n%s", i, plasmid)
			}
		}
	}

	if len(bacteria.Plasmids) > 0 {
		result += "\n"
	}

	return result
}

func (bacteria *ProgoBact) Simulate() {
	bacteria.SimpleBacteria.Simulate()

	bacteria.Chromosome.GetDNAs()[0].(*ProgoDNA).Eval(bacteria.State)
}

func (progobact *ProgoBact) Reproduce() Creature {
	base := progobact.SimpleBacteria.Reproduce()
	if base == nil {
		return nil
	}

	simpleBacteria := base.(*SimpleBacteria)
	simpleBacteria.Parents = []Creature{progobact}

	newState := progobact.State.Clone()
	newState.ExternalData = base.GetEnergy()

	return &ProgoBact{
		SimpleBacteria: simpleBacteria,
		State:          newState,
	}
}

func (bacteria *ProgoBact) Died() bool {
	return bacteria.SimpleBacteria.Died() || bacteria.State.Crashed
}
