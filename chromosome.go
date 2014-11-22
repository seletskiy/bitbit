package main

type SimpleChromosome struct {
	DNA DNA
}

func (chromosome *SimpleChromosome) GetDominantGene(offset int) Gene {
	return chromosome.DNA.GetGene(offset)
}

func (chromosome *SimpleChromosome) GetRecessiveGene(offset int) Gene {
	return chromosome.DNA.GetGene(offset)
}

func (chromosome *SimpleChromosome) String() string {
	return chromosome.DNA.String()
}

func (chromosome *SimpleChromosome) GetLength() int {
	return chromosome.DNA.GetLength()
}

func (chromosome *SimpleChromosome) Clone() Chromosome {
	return &SimpleChromosome{
		DNA: chromosome.DNA.Copy(),
	}
}
