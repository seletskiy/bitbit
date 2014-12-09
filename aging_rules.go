package main

type AgingRules struct {
	DieAge int
}

func (rules AgingRules) Apply(
	population *Population,
) {
	alive := Population{}

	for _, creature := range *population {
		if creature.GetAge() >= rules.DieAge {
			continue
		}

		alive = append(alive, creature)
	}

	*population = alive
}
