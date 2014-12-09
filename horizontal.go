package main

import "math/rand"

func HorizontalGeneTransfer(population Population) Population {
	for _, c := range population {
		extractPlasmid(c.(Bacteria))
		exchangePlasmids(c.(Bacteria), population)
		applyPlasmids(c.(Bacteria))
	}

	return population
}

func extractPlasmid(bacteria Bacteria) {
	plasmids := bacteria.GetPlasmids()
	if len(plasmids) >= 4 {
		return
	}

	if bacteria.GetAge() < 30*(len(plasmids)+1) {
		return
	}

	chromosome := bacteria.GetChromosome()

	codeStart := rand.Intn(
		int(float64(chromosome.GetLength()) / 1.2))
	codeLength := rand.Intn(
		int(float64(chromosome.GetLength() - codeStart)))

	if codeLength == 0 {
		return
	}

	code := make([]Gene, codeLength)
	for i := 0; i < codeLength; i++ {
		code[i] = chromosome.GetDominantGene(codeStart + i)
	}

	prefixLength := rand.Intn(int(codeLength/2) + 1)

	newPlasmid := &Plasmid{
		Applied:   true,
		Exchanged: false,
		Prefix:    code[:prefixLength],
		Code:      code[prefixLength:],
	}

	bacteria.SetPlasmids(append(plasmids, newPlasmid))
}

func exchangePlasmids(bacteria Bacteria, population []Creature) {
	plasmids := bacteria.GetPlasmids()
	if len(plasmids) == 0 {
		return
	}

	for _, c := range population {
		target := c.(Bacteria)

		if rand.Float64() < 0.9 {
			continue
		}

		if len(target.GetPlasmids()) >= 3 {
			continue
		}

		if target.GetAge() < 50 {
			continue
		}

		pn := rand.Intn(len(plasmids))

		logger.Log(Debug, "plasmid #%d exchange %p -> %p", pn, bacteria, target)

		plasmid := *plasmids[pn]
		plasmid.Applied = false
		plasmid.Exchanged = true

		target.SetPlasmids(append(target.GetPlasmids(), &plasmid))

		if rand.Float64() > 0.8 {
			bacteria.SetPlasmids(append(plasmids[:pn], plasmids[pn+1:]...))
		}

		break
	}
}

func applyPlasmids(bacteria Bacteria) {
	if rand.Float64() < 0.8 {
		return
	}

	plasmids := bacteria.GetPlasmids()
	if len(plasmids) == 0 {
		return
	}

	pn := rand.Intn(len(plasmids))
	plasmid := plasmids[pn]

	if plasmid.Applied {
		return
	}

	logger.Log(Debug, "plasmid #%d applied to %p", pn+1, bacteria)

	plasmid.Apply(bacteria.GetChromosome().(*SimpleChromosome))

	if rand.Float64() > 0.8 {
		bacteria.SetPlasmids(append(plasmids[:pn], plasmids[pn+1:]...))
	}
}
