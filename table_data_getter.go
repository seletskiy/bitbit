package main

type TableDataGetter interface {
	GetData(tableIndex, cellIndex int) float64
}
