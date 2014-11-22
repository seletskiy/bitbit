package main

type MarketEnvironment struct {
	NaturalSelectionEnvironment
	FreeFundsPool float64
}

func (environment *MarketEnvironment) Reap(
	tick int,
	population Population,
) Population {
	for _, individual := range population {
		bacteria := individual.(*ProgoBact)

		if bacteria.Died() {
			environment.FreeFundsPool += bacteria.Energy
		}
	}

	return environment.NaturalSelectionEnvironment.Reap(tick, population)
}
