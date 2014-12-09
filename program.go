package main

import (
	"errors"
	"fmt"
)

type ProgramInstruction interface {
	Eval(state *ProgramState) error
	GetArgsCount() int
	SetArg(index int, arg ProgramInstructionArg)
	GetArg(index int) ProgramInstructionArg
	Init()
	Copy() ProgramInstruction
}

type ProgramInstructionValue interface {
	GetFloat64() float64
	GetInt() int
}

// @TODO: consider fill it with methods
type ProgramInstructionArg interface{}

type ProgramState struct {
	IPS          int
	Crashed      bool
	Memory       *ProgramMemory
	ExternalData interface{}
	CrashReport  error
}

type Codepoint struct {
	Label       string
	Instruction ProgramInstruction
}

type Program []Codepoint

func (program Program) String() string {
	result := ""

	jumpPath := -1
	for i, codepoint := range program {
		jumpPrefix := " "
		if jumpPath >= 0 {
			if jumpPath == 0 {
				jumpPrefix = "X"
			} else {
				jumpPrefix = "↓"
			}
			jumpPath--
		}

		switch jump := codepoint.Instruction.(type) {
		case *ProgramInstructionJumpGreaterThan:
			if jump.Jumped {
				jumpPath = int(jump.Jump.(ForwardJump))
				jumpPrefix = "*"
			}
		}

		result += fmt.Sprintf(
			"OP<%012p> %s %03d   %s\n",
			codepoint.Instruction,
			jumpPrefix,
			i, codepoint,
		)
	}

	result += fmt.Sprintf("OP<%#012x>   %03d   END\n", 0, len(program))

	return result
}

func (codepoint Codepoint) String() string {
	return fmt.Sprint(codepoint.Instruction)
}

func (program Program) Eval(state *ProgramState) error {
	// @TODO: move to separate method
	for _, codepoint := range program {
		switch jump := codepoint.Instruction.(type) {
		case *ProgramInstructionJumpGreaterThan:
			jump.Jumped = false
		}
	}

	state.Crashed = false

	state.IPS = 0

	var err error

	for {
		ips := state.IPS

		err = program[ips].Instruction.Eval(state)
		if err != nil {
			state.Crashed = true
			break
		}

		if ips == state.IPS {
			state.IPS++
		}

		if state.IPS == len(program) {
			return nil
		}

		if state.IPS < 0 {
			state.Crashed = true
			state.IPS = ips
			err = errors.New(`IPS is out of bounds`)
			break
		}

		if state.IPS > len(program) {
			state.Crashed = true
			state.IPS = ips
			err = errors.New(`IPS is out of bounds`)
			break
		}
	}

	state.CrashReport = err

	return err
}

func (state *ProgramState) String() string {
	result := ""

	result += fmt.Sprintf("%s\n\n", state.Memory)

	result += "\n"
	result += fmt.Sprintf("External:\n%s", state.ExternalData)
	result += "\n"
	result += "\n"
	result += fmt.Sprintf("IPS @ %03d", state.IPS)
	if state.Crashed {
		result += fmt.Sprintf(" !!! CRASHED: %s\n", state.CrashReport)
	} else {
		result += "\n"
	}

	return result
}

func (state *ProgramState) Clone() *ProgramState {
	return &ProgramState{
		Memory: state.Memory.Clone(),
	}
}
