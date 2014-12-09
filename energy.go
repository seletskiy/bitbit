package main

type Energy interface {
	TransferTo(target Energy)
	Split() Energy
	Scatter(n int) []Energy
	Void() bool
	Simulate()
	GetFloat64() float64
	Free()
}
