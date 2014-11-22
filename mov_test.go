package main

import (
	"bitbit/prog/state"
	"fmt"
	"math/rand"
	"testing"
)

var _ = fmt.Println

func TestMOVOnIntRegister(t *testing.T) {
	s := &state.ProgramState{ProgramMemory: state.NewProgramMemory(3)}

	intVal := IntValue(42)
	reg0 := &IntRegister{Number: 0}
	reg1 := &IntRegister{Number: 1}

	(&MOV{reg0, intVal}).Eval(s)

	if s.ProgramMemory.GetInt(0) != intVal.GetValue(s) {
		t.Fatal("I[0] != 42")
	}

	(&MOV{reg1, reg0}).Eval(s)

	if s.ProgramMemory.GetInt(1) != intVal.GetValue(s) {
		t.Fatal("I[0] != I[1]")
	}

	if (&MOV{reg0, intVal}).String() != "MOV I[0] 42" {
		t.Fatal("string MOV{reg0, intVal}")
	}
}

func TestMOVPanic(t *testing.T) {
	defer func() {
		e := recover()
		if fmt.Sprintf("%s", e) != "interface conversion: interface "+
			"is float64, not int" {
			t.Fatalf("unexpected panic: %s", e)
		}
	}()

	s := &state.ProgramState{ProgramMemory: state.NewProgramMemory(3)}

	(&MOV{&IntRegister{Number: 0}, &FltRegister{Number: 0}}).Eval(s)
}

func TestMOVPanic2(t *testing.T) {
	defer func() {
		e := recover()
		if fmt.Sprintf("%s", e) != "interface conversion: interface "+
			"is int, not float64" {
			t.Fatalf("unexpected panic: %s", e)
		}
	}()

	s := &state.ProgramState{ProgramMemory: state.NewProgramMemory(3)}

	(&MOV{&FltRegister{Number: 0}, &IntRegister{Number: 0}}).Eval(s)
}

func ExampleRandMOV() {
	rand.Seed(2)

	fmt.Println(NewRandMOV(1, 0.5, 0.5, 5))
	fmt.Println(NewRandMOV(9, 1.0, 0.0, 5))
	fmt.Println(NewRandMOV(1, 0.0, 0.0, 5))
	fmt.Println(NewRandMOV(1, 0.0, 1.0, 10))

	// Output:
	// MOV I[1] I[0]
	// MOV I[4] 2
	// MOV F[1] -0.1539457
	// MOV F[2] F[7]
}
