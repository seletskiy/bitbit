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

type ByAvgError []Creature

func (creatures ByAvgError) Len() int {
	return len(creatures)
}

func (creatures ByAvgError) Swap(i, j int) {
	creatures[i], creatures[j] = creatures[j], creatures[i]
}

func (creatures ByAvgError) Less(i, j int) bool {
	if creatures[i].GetEnergy().Void() {
		return true
	}

	if creatures[j].GetEnergy().Void() {
		return false
	}

	a := creatures[i].GetEnergy().(ErrorBasedEnergy).GetAvgError()
	b := creatures[j].GetEnergy().(ErrorBasedEnergy).GetAvgError()

	return a < b
}

type ByMaxError []Creature

func (creatures ByMaxError) Len() int {
	return len(creatures)
}

func (creatures ByMaxError) Swap(i, j int) {
	creatures[i], creatures[j] = creatures[j], creatures[i]
}

func (creatures ByMaxError) Less(i, j int) bool {
	if creatures[i].GetEnergy().Void() {
		return true
	}

	if creatures[j].GetEnergy().Void() {
		return false
	}

	a := creatures[i].GetEnergy().(ErrorBasedEnergy).GetMaxError()
	b := creatures[j].GetEnergy().(ErrorBasedEnergy).GetMaxError()

	return a < b
}
