package main

type ProgramInstance struct {
	Code  Program
	State ProgramState
}

func (instance *ProgramInstance) Run() error {
	for {
		ips := instance.State.IPS
		instance.Code[ips].Instruction.Eval(&instance.State)

		if ips == instance.State.IPS {
			instance.State.IPS += 1
		}

		if instance.State.IPS < 0 {
			break
		}

		if instance.State.IPS >= len(instance.Code) {
			return nil // @TODO error
		}
	}

	return nil
}
