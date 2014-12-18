package main

import "math/rand"

type BacterialGeneTransferRules struct {
	BirthTransferProbability      float64
	ReproduceLossProbability      float64
	TransferLossProbability       float64
	ApplyLossProbability          float64
	ApplyPlasmidProbability       float64
	MaxPlasmidsNumber             int
	PlasmidPerAge                 int
	ExchangePlasmidsProbability   float64
	MinAgeForExchange             int
	PlasmidPrefixLengthProportion float64
	MinPlasmidPrefixLength        int
	MaxPlasmidLength              int
}

func (rules BacterialGeneTransferRules) Apply(
	population *Population,
) {
	rules.transferPlasmidsOnBirth(population)
	rules.applyPlasmids(population)
	rules.extractPlasmids(population)
	rules.exchangePlasmids(population)
}

func (rules BacterialGeneTransferRules) transferPlasmidsOnBirth(
	population *Population,
) {
	for _, creature := range *population {
		if creature.GetAge() > 0 {
			continue
		}

		parents := creature.GetParents()
		if len(parents) == 0 {
			continue
		}

		parent := parents[0]

		keepPlasmids := make([]*Plasmid, 0)
		for plasmidIndex, plasmid := range parent.(Bacteria).GetPlasmids() {
			if rand.Float64() > rules.BirthTransferProbability {
				continue
			}

			plasmidCopy := *plasmid
			plasmidCopy.Applied = false
			plasmidCopy.Self = false
			plasmidCopy.Age = plasmidCopy.Age + 1

			creature.(Bacteria).SetPlasmids(
				append(creature.(Bacteria).GetPlasmids(), &plasmidCopy),
			)

			Log(Debug, "CREATURE<%p> -> CREATURE<%p> plasmid #%d transfer on birth",
				parent,
				creature,
				plasmidIndex,
			)

			if rand.Float64() < rules.ReproduceLossProbability {
				keepPlasmids = append(keepPlasmids, plasmid)

				Log(Debug, "CREATURE<%p> plasmid #%d lost",
					parent,
					plasmidIndex,
				)
			}
		}

		parent.(Bacteria).SetPlasmids(keepPlasmids)
	}
}

func (rules BacterialGeneTransferRules) extractPlasmids(
	population *Population,
) {
	for _, creature := range *population {
		plasmids := creature.(Bacteria).GetPlasmids()
		if len(plasmids) >= rules.MaxPlasmidsNumber {
			return
		}

		if creature.GetAge() < rules.PlasmidPerAge*(len(plasmids)+1) {
			return
		}

		chromosome := creature.GetChromosome()

		codeStart := rand.Intn(chromosome.GetLength() - 1)
		codeLength := rand.Intn(chromosome.GetLength() - codeStart)
		if codeLength > rules.MaxPlasmidLength {
			codeLength = rules.MaxPlasmidLength
		}

		if codeLength == 0 {
			return
		}

		code := make([]Gene, codeLength)
		for i := 0; i < codeLength; i++ {
			code[i] = chromosome.GetDominantGene(codeStart + i)
		}

		prefixLength := rules.MinPlasmidPrefixLength + rand.Intn(
			int(float64(codeLength)*rules.PlasmidPrefixLengthProportion)+1)

		newPlasmid := &Plasmid{
			Id:           rand.Int63(),
			Applied:      true,
			Exchanged:    false,
			Prefix:       code[:prefixLength],
			Code:         code[prefixLength:],
			Self:         true,
			ReplaceIndex: codeStart,
		}

		Log(Debug, "CREATURE<%p> plasmid extract: prefix %s code %s",
			creature, code[:prefixLength], code[prefixLength:],
		)

		creature.(Bacteria).SetPlasmids(append(plasmids, newPlasmid))
	}
}

func (rules BacterialGeneTransferRules) exchangePlasmids(
	population *Population,
) {
	for _, creature := range *population {
		plasmids := creature.(Bacteria).GetPlasmids()
		if len(plasmids) == 0 {
			return
		}

		for _, target := range *population {
			if creature == target {
				continue
			}

			if rand.Float64() > rules.ExchangePlasmidsProbability {
				continue
			}

			if len(target.(Bacteria).GetPlasmids()) >= rules.MaxPlasmidsNumber {
				continue
			}

			if target.GetAge() < rules.MinAgeForExchange {
				continue
			}

			plasmidIndex := rand.Intn(len(plasmids))
			plasmid := *plasmids[plasmidIndex]
			targetAlreadyHavePlasmid := false
			for _, targetPlasmid := range target.(Bacteria).GetPlasmids() {
				if targetPlasmid.Id == plasmid.Id {
					targetAlreadyHavePlasmid = true
				}
			}

			if targetAlreadyHavePlasmid {
				break
			}

			Log(Debug,
				"CREATURE<%p> -> CREATURE<%p> plasmid #%d exchange",
				creature, target, plasmidIndex,
			)

			plasmid.Applied = false
			plasmid.Exchanged = true
			plasmid.Self = false

			target.(Bacteria).SetPlasmids(
				append(target.(Bacteria).GetPlasmids(), &plasmid))

			if rand.Float64() < rules.TransferLossProbability {
				creature.(Bacteria).SetPlasmids(
					append(
						plasmids[:plasmidIndex],
						plasmids[plasmidIndex+1:]...,
					),
				)
			}

			break
		}
	}
}

func (rules BacterialGeneTransferRules) applyPlasmids(
	population *Population,
) {
	for _, creature := range *population {
		if rand.Float64() > rules.ApplyPlasmidProbability {
			continue
		}

		plasmids := creature.(Bacteria).GetPlasmids()
		if len(plasmids) == 0 {
			continue
		}

		plasmidIndex := rand.Intn(len(plasmids))
		plasmid := plasmids[plasmidIndex]

		if plasmid.Applied {
			continue
		}

		chromosome := creature.GetChromosome().(*SimpleChromosome)
		applied, applyIndex := plasmid.Apply(chromosome)
		if applied {
			plasmid.AppliedCount++
			Log(Debug,
				"CREATURE<%p> plasmid #%d applied to offset %d",
				creature,
				plasmidIndex,
				applyIndex,
			)
		} else {
			Log(Debug,
				"CREATURE<%p> plasmid #%d does not match",
				creature,
				plasmidIndex,
			)
		}

		if rand.Float64() < rules.ApplyLossProbability {
			Log(Debug,
				"CREATURE<%p> plasmid #%d lost",
				creature, plasmidIndex,
			)

			creature.(Bacteria).SetPlasmids(
				append(
					plasmids[:plasmidIndex],
					plasmids[plasmidIndex+1:]...,
				),
			)
		}
	}
}
