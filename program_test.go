package main

import (
	"fmt"
	"testing"
)

func TestNewRandLabel(t *testing.T) {
	fmt.Println(NewRandProgram(NewProgramLayout(10), 10))
	t.Fatal(1)
}
