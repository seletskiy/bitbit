package main

import "fmt"

type Value float64

type Reference interface {
	GetValue(*ProgramState) Value
}

type Register interface {
	SetValue(*ProgramState, Value)
	GetValue(*ProgramState) Value
}

type FloatValue float64

func (value FloatValue) GetValue(state *ProgramState) Value {
	return Value(value)
}

func (value FloatValue) String() string {
	return fmt.Sprintf("%.7f", float64(value))
}

type FloatRegister int

func (register FloatRegister) GetValue(state *ProgramState) Value {
	return Value(state.Memory.GetFloat(int(register)))
}

func (register FloatRegister) SetValue(state *ProgramState, value Value) {
	state.Memory.SetFloat(int(register), float64(value))
}

func (register FloatRegister) String() string {
	return fmt.Sprintf("[%d]", int(register))
}
