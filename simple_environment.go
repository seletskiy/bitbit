package main

import "time"

type SimpleEnvironment struct {
	Rules []Rules
}

func (environment *SimpleEnvironment) Simulate(population *Population) {
	for _, rules := range environment.Rules {
		startTime := time.Now()
		rules.Apply(population)
		elapsedTime := time.Duration(
			int(time.Since(startTime)) / len(*population))
		logger.Log(Debug, "TIME: %T took %s per creature", rules, elapsedTime)
	}
}
