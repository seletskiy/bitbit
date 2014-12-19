package main

import (
	"fmt"
	"log"
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
		p.ReplaceIndex = variants[rand.Intn(len(variants))]
		dna.Replace(p.ReplaceIndex+len(p.Prefix), p.Code)
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
	for startIndex, _ := range dnaCode {
		matched := true
		for checkIndex, gene2 := range prefix {
			if startIndex+checkIndex >= len(dnaCode) {
				matched = false
				break
			}

			gene := dnaCode[startIndex+checkIndex]
			if !dna.EqGenes(gene, gene2) {
				matched = false
				break
			}
		}

		if matched {
			result = append(result, startIndex)
		}
	}

	log.Printf("%#v", result)

	return result
}
