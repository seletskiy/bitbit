package main

type SimulationRules struct{}

func (rules SimulationRules) Apply(
	population *Population,
) {
	for _, creature := range *population {
		// @TODO: remove this?
		if creature.GetEnergy() == nil && creature.GetAge() == 0 {
			continue
		}

		creature.Simulate()
	}

	Log(Debug, "SIMULATION COMPLETE")
	Log(Debug, "POPULATION (%d)\n%s", len(*population), *population)
}
