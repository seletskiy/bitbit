package main

import (
	"fmt"
	"log"
)

type ProgDNA struct {
	*Program
}

func (program *ProgDNA) EqGenes(a Gene, bacteria Gene) bool {
	return fmt.Sprintf("%s", a.(Codepoint).Instruction) ==
		fmt.Sprintf("%s", bacteria.(Codepoint).Instruction)
}

func (program *ProgDNA) GetCode() []Gene {
	genes := make([]Gene, 0)

	for _, cp := range *program.Program {
		genes = append(genes, Codepoint{
			Label:       cp.Label,
			Instruction: cp.Instruction,
		})
	}

	return genes
}

func (program *ProgDNA) Replace(offset int, newCode []Gene) {
	log.Printf("%program replace offset %d", program.Program, offset)
	for i, cp := range newCode {
		if offset+i >= len(*program.Program) {
			break
		}

		(*program.Program)[offset+i].Instruction = cp.(Codepoint).Instruction
	}
}

func (program *ProgDNA) GetGene(offset int) Gene {
	return (*program.Program)[offset]
}

func (program *ProgDNA) GetLength() int {
	return len(*program.Program)
}

func (program *ProgDNA) Copy() DNA {
	newProgram := make(Program, len(*program.Program))

	for i := 0; i < len(newProgram); i++ {
		newProgram[i] = Codepoint{
			Label:       (*program.Program)[i].Label,
			Instruction: (*program.Program)[i].Instruction,
		}
	}

	return &ProgDNA{
		&newProgram,
	}
}
