package main

type EloBasedEnergy interface {
	GetEloScore() int
	SetEloScore(int)
}

type EloEnergy struct {
	Base         Energy
	Score        int
	BaseScore    int
	MinEloRating int
}

func (origin *EloEnergy) GetFloat64() float64 {
	return float64(origin.GetEloScore())
}

func (origin *EloEnergy) GetEloScore() int {
	return origin.Score
}

func (origin *EloEnergy) SetEloScore(score int) {
	origin.Score = score
}

func (origin *EloEnergy) Void() bool {
	if origin.Base.Void() {
		return true
	}

	return origin.Score < origin.MinEloRating
}

func (origin *EloEnergy) Scatter(amount int) []Energy {
	baseScattered := origin.Base.Scatter(amount)
	scattered := make([]Energy, len(baseScattered))
	for index, base := range baseScattered {
		scattered[index] = &EloEnergy{
			Base:         base,
			BaseScore:    origin.BaseScore,
			Score:        origin.BaseScore,
			MinEloRating: origin.MinEloRating,
		}
	}

	return scattered
}

func (origin *EloEnergy) TransferTo(energy Energy) {
	origin.Base.TransferTo(energy.(*EloEnergy).Base)
}

func (origin *EloEnergy) Split() Energy {
	scattered := origin.Scatter(2)

	if len(scattered) < 1 {
		return nil
	}

	rest := scattered[0].(*EloEnergy)

	rest.TransferTo(origin)

	if len(scattered) < 2 {
		return nil
	}

	return scattered[1]
}

func (origin *EloEnergy) Simulate() {
	origin.Base.Simulate()
}

func (origin *EloEnergy) Free() {
	origin.Base.Free()
}
