package main

type DataSource interface {
	Get(index int) float64
	Set(index int, value float64)
}
