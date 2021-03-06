package main

type ByEnergy []Creature

func (creatures ByEnergy) Len() int {
	return len(creatures)
}

func (creatures ByEnergy) Swap(i, j int) {
	creatures[i], creatures[j] = creatures[j], creatures[i]
}

func (creatures ByEnergy) Less(i, j int) bool {
	if creatures[i].GetEnergy().Void() {
		return true
	}

	if creatures[j].GetEnergy().Void() {
		return false
	}

	return creatures[i].GetEnergy().GetFloat64() <
		creatures[j].GetEnergy().GetFloat64()
}

type ByEnergyReverse []Creature

func (creatures ByEnergyReverse) Len() int {
	return len(creatures)
}

func (creatures ByEnergyReverse) Swap(i, j int) {
	creatures[i], creatures[j] = creatures[j], creatures[i]
}

func (creatures ByEnergyReverse) Less(i, j int) bool {
	if creatures[i].GetEnergy().Void() {
		return false
	}

	if creatures[j].GetEnergy().Void() {
		return true
	}

	return creatures[i].GetEnergy().GetFloat64() >
		creatures[j].GetEnergy().GetFloat64()
}

type ByAge []Creature

func (creatures ByAge) Len() int {
	return len(creatures)
}

func (creatures ByAge) Swap(i, j int) {
	creatures[i], creatures[j] = creatures[j], creatures[i]
}

func (creatures ByAge) Less(i, j int) bool {
	return creatures[i].GetAge() < creatures[j].GetAge()
}

type ByCurrentError []Creature

func (creatures ByCurrentError) Len() int {
	return len(creatures)
}

func (creatures ByCurrentError) Swap(i, j int) {
	creatures[i], creatures[j] = creatures[j], creatures[i]
}

func (creatures ByCurrentError) Less(i, j int) bool {
	if creatures[i].GetEnergy().Void() {
		return true
	}

	if creatures[j].GetEnergy().Void() {
		return false
	}

	a := creatures[i].GetEnergy().(ErrorBasedEnergy).GetCurrentError()
	b := creatures[j].GetEnergy().(ErrorBasedEnergy).GetCurrentError()

	return a < b
}

type ByEloScore []Creature

func (creatures ByEloScore) Len() int {
	return len(creatures)
}

func (creatures ByEloScore) Swap(i, j int) {
	creatures[i], creatures[j] = creatures[j], creatures[i]
}

func (creatures ByEloScore) Less(i, j int) bool {
	if creatures[i].GetEnergy().Void() {
		return true
	}

	if creatures[j].GetEnergy().Void() {
		return false
	}

	a := creatures[i].GetEnergy().(EloBasedEnergy).GetEloScore()
	b := creatures[j].GetEnergy().(EloBasedEnergy).GetEloScore()

	return a > b
}

type ByWinsInRow []*EloPlayer

func (players ByWinsInRow) Len() int {
	return len(players)
}

func (players ByWinsInRow) Swap(i, j int) {
	players[i], players[j] = players[j], players[i]
}

func (players ByWinsInRow) Less(i, j int) bool {
	energyA := players[i].Player.(Creature).GetEnergy()
	energyB := players[j].Player.(Creature).GetEnergy()
	if energyA.Void() {
		return true
	}

	if energyB.Void() {
		return false
	}

	a := energyA.(EloBasedEnergy).GetWinsInRow()
	b := energyB.(EloBasedEnergy).GetWinsInRow()

	return a < b
}
