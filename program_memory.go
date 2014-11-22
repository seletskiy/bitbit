package main

import (
	"fmt"
	"strings"
)

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
	result := make([]string, 0)

	for i, cell := range m.Cells {
		result = append(result, fmt.Sprintf("%03d  %10.7f", i, cell))
	}

	return strings.Join(result, "\n")
}
