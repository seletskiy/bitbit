package main

import (
	"fmt"
	"math/rand"
	"strings"
)

type Plasmid struct {
	Id           int64
	Self         bool
	Applied      bool
	Exchanged    bool
	Prefix       []Gene
	Code         []Gene
	ReplaceIndex int
	AppliedCount int
	Age          int
}

func (p *Plasmid) Apply(chromosome *SimpleChromosome) (bool, int) {
	dna := chromosome.DNA
	if variants := matchPrefix(dna, p.Prefix); len(variants) > 0 {
		p.ReplaceIndex = rand.Intn(len(variants))
		dna.Replace(variants[p.ReplaceIndex], p.Code)
		p.Applied = true
		return true, p.ReplaceIndex
	} else {
		return false, 0
	}
}

func (p *Plasmid) String() string {
	result := make([]string, 0)

	result = append(result,
		fmt.Sprintf(
			"~ AGE %d APPLIED %d",
			p.Age, p.AppliedCount,
		),
	)

	foreign := ' '
	if p.Exchanged {
		foreign = 'E'
	}

	if p.Self {
		foreign = 'S'
	}

	applied := ' '
	if p.Applied {
		applied = '+'
	}

	for _, g := range p.Prefix {
		result = append(result, fmt.Sprintf("  ?  %s", g))
	}

	for _, g := range p.Code {
		result = append(result, fmt.Sprintf(" %c %c %s", foreign, applied, g))
	}

	return strings.Join(result, "\n")
}

func matchPrefix(dna DNA, prefix []Gene) []int {
	result := make([]int, 0)

	dnaCode := dna.GetCode()
	for j, _ := range dnaCode {
		i := 0
		matched := true
		for _, gene2 := range prefix {
			if j+i >= len(dnaCode) {
				matched = false
				break
			}

			gene := dnaCode[j+i]
			if dna.EqGenes(gene, gene2) {
				continue
			} else {
				matched = false
				break
			}

			j += 1
		}

		if matched {
			result = append(result, j)
		}
	}

	return result
}
