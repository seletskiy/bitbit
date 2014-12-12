package main

type AggressiveReproduceRules struct {
	MinAge int
}

func (rules AggressiveReproduceRules) Apply(population *Population) {
	toReproduce := []Creature{}
	for _, creature := range *population {
		if creature.Died() {
			continue
		}

		if creature.GetAge() < rules.MinAge {
			continue
		}

		toReproduce = append(toReproduce, creature)
	}

	reproduced := true
	for reproduced {
		reproduced = false
		for _, creature := range toReproduce {
			for {
				child := creature.Reproduce()
				if child == nil {
					break
				}

				logger.Log(Debug, "CREATURE<%p> reproduce to CREATURE<%p>", creature, child)
				toReproduce = append(toReproduce, child)
				*population = append(*population, child)

				reproduced = true
			}
		}
	}
}
