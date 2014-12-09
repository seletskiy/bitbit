package main

type Environment interface {
	Simulate(population *Population)
}
