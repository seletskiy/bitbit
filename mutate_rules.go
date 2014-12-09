package main

import "math/rand"

type MutateRules struct {
	GeneGenerator           func(amount int) []Gene
	GeneMutator             func(gene Gene) Gene
	DNAMutationMaxSize      int
	DNAMutationProbability  float64
	GeneMutationProbability float64
}

func (rules MutateRules) Apply(population *Population) {
	for _, creature := range *population {
		if creature.GetAge() > 0 {
			continue
		}

		for _, dna := range creature.GetChromosome().GetDNAs() {
			if rand.Float64() < rules.DNAMutationProbability {
				logger.Log(Debug,
					"CREATURE<%p> DNA<%p> mutate dna", creature, dna,
				)
				rules.mutateDNA(dna)
			}

			if rand.Float64() < rules.GeneMutationProbability {
				logger.Log(Debug,
					"CREATURE<%p> DNA<%p> mutate single gene", creature, dna,
				)

				rules.mutateGene(dna)
			}
		}
	}
}

func (rules MutateRules) mutateDNA(dna DNA) {
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
}

func (rules MutateRules) mutateGene(dna DNA) {
	offset := rand.Intn(dna.GetLength())

	originGene := dna.GetGene(offset)

	logger.Log(Debug,
		"DNA<%p> gene chosen to mutate: %v", dna, originGene,
	)

	mutateGene := rules.GeneMutator(originGene)

	if mutateGene == nil {
		return
	}

	dna.Replace(offset, []Gene{mutateGene})
}
