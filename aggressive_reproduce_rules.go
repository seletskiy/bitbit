package main

type AggressiveReproduceRules struct {
}

func (rules AggressiveReproduceRules) Apply(population *Population) {
	reproduced := true
	for reproduced {
		reproduced = false
		for _, creature := range *population {
			if creature.Died() {
				continue
			}

			for {
				child := creature.Reproduce()
				if child == nil {
					break
				}

				logger.Log(Debug, "CREATURE<%p> reproduce to CREATURE<%p>", creature, child)
				*population = append(*population, child)

				reproduced = true
			}
		}
	}
}
