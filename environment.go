package main

import "fmt"

type Environment interface {
	Simulate(tick int, population Population)
	Reap(tick int, population Population)
}

type SimpleEnvironment struct{}

func (environment *SimpleEnvironment) Simulate(
	tick int,
	population Population,
) Population {
	for _, individual := range population {
		individual.Simulate()

		fmt.Printf("TICK %d %s\n", tick, individual)
		fmt.Println()
	}

	return population
}

func (environment *SimpleEnvironment) Reap(
	tick int,
	population Population,
) Population {
	alive := Population{}

	for _, individual := range population {
		if individual.Died() {
			continue
		}

		alive = append(alive, individual)
	}

	return alive
}
