package main

type Bacteria interface {
	Creature
	GetPlasmids() []*Plasmid
	SetPlasmids([]*Plasmid)
}
