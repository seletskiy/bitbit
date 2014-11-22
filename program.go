package main

import (
	"fmt"
	"log"
	"strings"
)

type ProgramInstruction interface {
	Eval(*ProgramState)
}

type ProgramInstructionArgsValueSetter interface {
	SetValue(Reference)
}

type ProgramInstructionArgsToValueSetter interface {
	ProgramInstructionArgsValueSetter
	SetTo(Register)
}

type ProgramInstructionArgsValue struct {
	Value Reference
}

type ProgramInstructionArgsToValue struct {
	To Register
	ProgramInstructionArgsValue
}

func (op *ProgramInstructionArgsToValue) SetTo(register Register) {
	op.To = register
}

func (op *ProgramInstructionArgsValue) SetValue(reference Reference) {
	op.Value = reference
}

func (op *ProgramInstructionArgsValue) Eval(state *ProgramState) {
	panic(`must be reimplemented`)
}

type ProgramState struct {
	IPS          int
	Crashed      bool
	Memory       *ProgramMemory
	ExternalData interface{}
}

type Codepoint struct {
	Label       string
	Instruction ProgramInstruction
}

type Program []Codepoint

func (program *Program) String() string {
	result := []string{}

	for i, cp := range *program {
		result = append(result, fmt.Sprintf("%03d   %s", i, cp.Instruction))
	}

	result = append(result, fmt.Sprintf("%03d   END", len(*program)))

	return strings.Join(result, "\n")
}

func (c Codepoint) String() string {
	return fmt.Sprintf("%s", c.Instruction)
}

func (program *Program) Eval(state *ProgramState) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("crash: %s", err)
			state.Crashed = true
		}
	}()

	state.Crashed = false

	state.IPS = 0

	for {
		ips := state.IPS

		(*program)[ips].Instruction.Eval(state)

		if ips == state.IPS {
			state.IPS += 1
		}

		if state.IPS < 0 {
			return
		}

		if state.IPS >= len(*program) {
			return
		}
	}
}

func (state *ProgramState) String() string {
	result := ""

	result += fmt.Sprintf("%s\n\n", state.Memory)

	result += fmt.Sprintf("IPS @ %03d", state.IPS)
	if state.Crashed {
		result += " !!! CRASHED\n"
	} else {
		result += "\n"
	}

	result += "\n"
	result += fmt.Sprintf("External:\n%s", state.ExternalData)
	result += "\n"

	return result
}
