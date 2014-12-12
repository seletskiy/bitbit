package main

type ValueStorer interface {
	Set(value float64)
	Get() float64
}

type ExternalValueEnergy struct {
	*ErrorBasedEnergy
	Store ValueStorer
}

func (origin *ExternalValueEnergy) Simulate() {
	origin.ErrorBasedEnergy.Simulate()

	origin.SetTargetValue(origin.Store.Get())
}

func (origin *ExternalValueEnergy) TransferTo(energy Energy) {
	origin.ErrorBasedEnergy.TransferTo(
		energy.(*ExternalValueEnergy).ErrorBasedEnergy,
	)
}

func (origin *ExternalValueEnergy) Scatter(n int) []Energy {
	scattered := []Energy{}
	for _, part := range origin.ErrorBasedEnergy.Scatter(n) {
		scattered = append(scattered,
			&ExternalValueEnergy{
				ErrorBasedEnergy: part.(*ErrorBasedEnergy),
				Store:            origin.Store,
			},
		)
	}

	return scattered
}

func (origin *ExternalValueEnergy) Split() Energy {
	child := origin.ErrorBasedEnergy.Split()
	if child == nil {
		return nil
	}

	return &ExternalValueEnergy{
		ErrorBasedEnergy: child.(*ErrorBasedEnergy),
		Store:            origin.Store,
	}
}
