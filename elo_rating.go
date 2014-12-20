package main

import (
	"container/list"
	"math"
	"math/rand"
	"sort"
)

type EloMatchResult float64

const EloCoefficientHigh = 15
const EloCoefficientNormal = 10
const EloCoefficientThreshold = 2400

const (
	EloMatchWin  EloMatchResult = 1.0
	EloMatchDraw EloMatchResult = 0.5
	EloMatchLoss EloMatchResult = 0.0
)

type EloRatings struct {
	*list.List
	StrongProbability float64
	OpponentVariance  float64
}

type EloPlayer struct {
	Player interface{}
	Score  int
}

// players should be sorted! best comes first!
func NewEloRatings(players []*EloPlayer) *EloRatings {
	ratings := &EloRatings{
		List: list.New(),
	}

	for _, player := range players {
		ratings.PushBack(player)
	}

	return ratings
}

func (ratings *EloRatings) Update(player *EloPlayer) {
	ratings.replaceElementSorted(player)
}

func (ratings *EloRatings) replaceElementSorted(player *EloPlayer) {
	element := ratings.findElement(player)
	ratings.Remove(element)

	for ptr := ratings.Front(); ptr != nil; ptr = ptr.Next() {
		if ptr.Value.(*EloPlayer).Score <= player.Score {
			ratings.InsertBefore(player, ptr)
			return
		}
	}

	ratings.PushBack(player)
}

func (ratings *EloRatings) findElement(player *EloPlayer) *list.Element {
	for ptr := ratings.Front(); ptr != nil; ptr = ptr.Next() {
		if ptr.Value.(*EloPlayer) == player {
			return ptr
		}
	}

	panic(123)

	return nil
}

func (ratings *EloRatings) ChooseOpponent(player *EloPlayer) *EloPlayer {
	players := make([]*EloPlayer, ratings.Len())

	index := 0
	for ptr := ratings.Front(); ptr != nil; ptr = ptr.Next() {
		players[index] = ptr.Value.(*EloPlayer)
		index++
	}

	sort.Sort(ByWinsInRow(players))

	playerIndex := 0
	for i, ptr := range players {
		if player == ptr {
			playerIndex = i
		}
	}

	//for _, player := range players {
	//    log.Printf("XXX %p %d",
	//        player.Player,
	//        player.Player.(Creature).GetEnergy().(EloBasedEnergy).GetWinsInRow(),
	//    )
	//}

	rangeStart := playerIndex - int(ratings.OpponentVariance/2)
	rangeEnd := playerIndex + int(ratings.OpponentVariance/2)

	if rangeStart < 0 {
		rangeEnd += -rangeStart
		rangeStart = 0
	}

	if rangeEnd >= len(players) {
		rangeStart -= rangeEnd - len(players)
		rangeEnd = len(players) - 1
	}

	versusIndex := 0
	for {
		versusIndex = rand.Intn(rangeEnd-rangeStart) + rangeStart
		if versusIndex != playerIndex {
			break
		}
	}

	//log.Printf("XXX %d vs %d", playerIndex, versusIndex)

	return players[versusIndex]
}

func (ratings *EloRatings) Compute(
	player, versus *EloPlayer, result EloMatchResult,
) int {
	diff := player.Score - versus.Score
	delta := 1 / ((math.Pow(10, -float64(diff)/400.0)) + 1)

	newPlayerScore := float64(player.Score) +
		ratings.GetEloCoefficient(player.Score, versus.Score)*
			(float64(result)-delta)

	return int(math.Ceil(newPlayerScore))
}

func (ratings *EloRatings) GetEloCoefficient(scoreA, scoreB int) float64 {
	if scoreA > EloCoefficientThreshold {
		return EloCoefficientHigh
	}

	if scoreB > EloCoefficientThreshold {
		return EloCoefficientHigh
	}

	return EloCoefficientNormal
}
