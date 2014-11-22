package main

import "fmt"

const (
	programLength         = 10
	programMemorySize     = 10
	initialPopulationSize = 10
	initialFunds          = 100.0
	killPossibility       = 0.05
)

func main() {
	//logFile, _ := os.Create("log")
	//log.SetOutput(logFile)

	//rand.Seed(time.Now().UnixNano())

	population := Population{}

	layout := RandProgramLayout(programLength)

	for i := 0; i < initialPopulationSize; i++ {
		externalData := &DataStorage{
			FunValue: 1.0,
		}

		population = append(
			population,
			RandProgoBact(
				layout,
				programMemorySize,
				initialFunds,
				externalData,
			),
		)
	}

	tick := 0
	environment := MarketEnvironment{
		NaturalSelectionEnvironment: NaturalSelectionEnvironment{
			KillPossibility: killPossibility,
		},
	}

	for len(population) > 0 {
		if tick > 1000 {
			return
		}

		fmt.Printf("POP SIZE: %d\n", len(population))
		population = environment.Simulate(tick, population)
		population = HorizontalGeneTransfer(population)
		population = CellFission(population)

		tick++
	}
}
