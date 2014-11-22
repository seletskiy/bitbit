package main

import (
	"fmt"
	"math/rand"
	"strings"
)

type Plasmid struct {
	Applied   bool
	Exchanged bool
	Prefix    []Gene
	Code      []Gene
}

func (p *Plasmid) Apply(chromosome *SimpleChromosome) bool {
	dna := chromosome.DNA
	if variants := matchPrefix(dna, p.Prefix); len(variants) > 0 {
		dna.Replace(variants[rand.Intn(len(variants))], p.Code)
		p.Applied = true
		return true
	} else {
		return false
	}
}

func (p *Plasmid) String() string {
	result := make([]string, 0)

	foreign := ' '
	if p.Exchanged {
		foreign = 'E'
	}

	applied := ' '
	if p.Applied {
		applied = '+'
	}

	for _, g := range p.Prefix {
		result = append(result, fmt.Sprintf("  #  %s", g))
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
		for j, gene2 := range prefix {
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
