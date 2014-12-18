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
