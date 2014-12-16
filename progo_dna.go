package main

import "fmt"

type ProgoDNA struct {
	*Program
}

func (program *ProgoDNA) EqGenes(a Gene, b Gene) bool {
	return fmt.Sprintf("%s", a.(Codepoint).Instruction) ==
		fmt.Sprintf("%s", b.(Codepoint).Instruction)
}

func (program *ProgoDNA) Eval(state *ProgramState) error {
	state.Memory.Zero()

	err := program.Program.Eval(state)
	if err != nil {
		return err
	}

	return nil
}

func (program *ProgoDNA) GetCode() []Gene {
	genes := make([]Gene, 0)

	for _, codepoint := range *program.Program {
		genes = append(genes,
			Codepoint{
				Label:       codepoint.Label,
				Instruction: codepoint.Instruction,
			},
		)
	}

	return genes
}

func (program *ProgoDNA) Replace(offset int, newCode []Gene) {
	Log(Debug,
		"DNA<%p> replace at offset %d: %s",
		program,
		offset,
		newCode,
	)

	for i, gene := range newCode {
		if offset+i >= len(*program.Program) {
			break
		}

		instruction := gene.(Codepoint).Instruction

		(*program.Program)[offset+i].Instruction = instruction
	}
}

func (program *ProgoDNA) GetGene(offset int) Gene {
	return (*program.Program)[offset]
}

func (program *ProgoDNA) GetLength() int {
	return len(*program.Program)
}

func (program *ProgoDNA) Copy() DNA {
	newProgram := make(Program, len(*program.Program))

	for i := 0; i < len(newProgram); i++ {
		newProgram[i] = Codepoint{
			Label:       (*program.Program)[i].Label,
			Instruction: (*program.Program)[i].Instruction.Copy(),
		}
	}

	return &ProgoDNA{
		&newProgram,
	}
}
