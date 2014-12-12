package main

import "fmt"

type ProgramArgReference interface {
	GetValue(*ProgramState) ProgramInstructionValue
}

type ProgramArgRegister interface {
	SetValue(*ProgramState, ProgramInstructionValue)
	GetValue(*ProgramState) ProgramInstructionValue
}

type ProgramArgJump interface {
	Apply(*ProgramState)
}

type ForwardJump uint

type Index uint

func (index Index) GetValue(state *ProgramState) ProgramInstructionValue {
	return ProgramInstructionValue(index)
}

func (index Index) GetInt() int {
	return int(index)
}

func (index Index) GetFloat64() float64 {
	panic("can't get float64 value")
}

func (index Index) String() string {
	return fmt.Sprintf("*%d", index)
}

func (jump ForwardJump) Apply(state *ProgramState) {
	state.IPS += int(jump)
}

func (jump ForwardJump) String() string {
	return fmt.Sprintf("+%d", int(jump))
}

type FloatValue float64

func (value FloatValue) GetValue(state *ProgramState) ProgramInstructionValue {
	return ProgramInstructionValue(value)
}

func (value FloatValue) String() string {
	return fmt.Sprintf("%.7f", float64(value))
}

func (value FloatValue) GetFloat64() float64 {
	return float64(value)
}

func (value FloatValue) GetInt() int {
	panic("can't get int from float value")
}

type FloatRegister int

func (register FloatRegister) GetValue(state *ProgramState) ProgramInstructionValue {
	return FloatValue(state.Memory.GetFloat(int(register)))
}

func (register FloatRegister) SetValue(state *ProgramState, value ProgramInstructionValue) {
	state.Memory.SetFloat(int(register), value.GetFloat64())
}

func (register FloatRegister) String() string {
	return fmt.Sprintf("[%d]", int(register))
}
