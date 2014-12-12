package main

import "fmt"

type ProgramMemory struct {
	Cells []float64
}

func NewProgramMemory(size int) *ProgramMemory {
	return &ProgramMemory{Cells: make([]float64, size)}
}

func (m *ProgramMemory) GetFloat(key int) float64 {
	return m.Cells[key]
}

func (m *ProgramMemory) SetFloat(key int, value float64) {
	m.Cells[key] = value
}

func (m *ProgramMemory) GetSize() int {
	return len(m.Cells)
}

func (m *ProgramMemory) String() string {
	result := ""

	for i, cell := range m.Cells {
		result += fmt.Sprintf("%03d  % 10.7g\n", i, cell)
	}

	return result
}

func (memory *ProgramMemory) Clone() *ProgramMemory {
	destinationCells := make([]float64, len(memory.Cells))

	copy(destinationCells, memory.Cells)

	return &ProgramMemory{
		Cells: destinationCells,
	}
}

func (memory *ProgramMemory) Zero() {
	for index, _ := range memory.Cells {
		memory.Cells[index] = 0
	}
}
