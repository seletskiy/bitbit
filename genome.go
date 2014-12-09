package main

type Gene interface{}

type DNA interface {
	GetGene(int) Gene
	GetCode() []Gene
	EqGenes(Gene, Gene) bool
	Replace(int, []Gene)
	String() string
	GetLength() int
	Copy() DNA
}

type Chromosome interface {
	GetDominantGene(int) Gene
	GetRecessiveGene(int) Gene
	GetLength() int
	Clone() Chromosome
	GetDNAs() []DNA
}
