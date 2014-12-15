package main

import "math/rand"

type MutateRules struct {
	GeneGenerator           func(amount int) []Gene
	GeneMutator             func(gene Gene) Gene
	DNAMutationMaxSize      int
	DNAMutationCount        int
	DNAMutationProbability  float64
	GeneMutationProbability float64
}

func (rules MutateRules) Apply(population *Population) {
	for _, creature := range *population {
		if creature.GetAge() > 0 {
			continue
		}

		rules.applyMutation(creature)
	}
}

func (rules MutateRules) applyMutation(creature Creature) {
	for _, dna := range creature.GetChromosome().GetDNAs() {
		if rand.Float64() < rules.DNAMutationProbability {
			Log(Debug,
				"CREATURE<%p> DNA<%p> mutate dna", creature, dna,
			)
			rules.mutateDNA(dna)
		}

		rules.mutateGenes(dna)
	}
}

func (rules MutateRules) mutateDNA(dna DNA) {
	mutiesCount := rules.DNAMutationCount
	for mutiesCount > 0 {
		offset := dna.GetLength() - rules.DNAMutationMaxSize
		length := rand.Intn(rules.DNAMutationMaxSize)
		if offset <= 0 {
			offset = dna.GetLength()
			length = dna.GetLength()
		}

		dna.Replace(
			rand.Intn(offset),
			rules.GeneGenerator(length+1),
		)

		mutiesCount--
	}
}

func (rules MutateRules) mutateGenes(dna DNA) {
	for geneIndex := 0; geneIndex < dna.GetLength(); geneIndex++ {
		originGene := dna.GetGene(geneIndex)

		if rand.Float64() > rules.GeneMutationProbability {
			continue
		}

		Log(Debug,
			"DNA<%p> gene chosen to mutate (%d): %v",
			dna, geneIndex, originGene,
		)

		mutateGene := rules.GeneMutator(originGene)

		if mutateGene == nil {
			continue
		}

		dna.Replace(geneIndex, []Gene{mutateGene})
	}
}
