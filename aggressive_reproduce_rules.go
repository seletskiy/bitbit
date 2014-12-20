package main

type AggressiveReproduceRules struct {
	MutateRules
}

func (rules AggressiveReproduceRules) Apply(population *Population) {
	toReproduce := []Creature{}
	for _, creature := range *population {
		if creature.Died() {
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

				Log(Debug,
					"CREATURE<%p> reproduce to CREATURE<%p>",
					creature, child,
				)

				rules.applyMutation(child)

				toReproduce = append(toReproduce, child)
				*population = append(*population, child)

				reproduced = true
			}
		}
	}
}
