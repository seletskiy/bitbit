package main

type ReapRules struct{}

func (rules ReapRules) Apply(
	population *Population,
) {
	alive := Population{}

	bonus := []Energy{}

	dead := 0
	for _, creature := range *population {
		if creature.Died() {
			bonus = append(bonus, creature.GetEnergy())
			dead++
			continue
		}

		alive = append(alive, creature)
	}

	scattered := []Energy{}
	for _, energy := range bonus {
		scattered = append(scattered, energy.Scatter(len(alive))...)
	}

	Log(Debug, "REAPER: scattering %d energy items to %d creatures",
		len(scattered),
		len(alive),
	)

	for bonusIndex, singleBonus := range scattered {
		creatureIndex := bonusIndex % len(alive)
		singleBonus.TransferTo(alive[creatureIndex].GetEnergy())
	}

	Log(Debug, "REAPER: %d dead", dead)

	*population = alive
}
