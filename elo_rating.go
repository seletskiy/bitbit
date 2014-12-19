package main

import (
	"container/list"
	"log"
	"math"
	"math/rand"
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
			//log.Printf("%#v", ptr)
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
	index := 0
	playerIndex := 0
	opponents := make([]*list.Element, ratings.Len())
	if len(opponents) <= 1 {
		return nil
	}

	for ptr := ratings.Front(); ptr != nil; ptr = ptr.Next() {
		opponents[index] = ptr
		if ptr.Value.(*EloPlayer) == player {
			playerIndex = index
		}
		index++
	}

	variance := ratings.OpponentVariance
	offset := ratings.StrongProbability
	if playerIndex < int(variance*ratings.StrongProbability) {
		offset = float64(playerIndex) / variance
	}

	if len(opponents)-playerIndex <= int(variance*(1-ratings.StrongProbability)) {
		offset = 1
	}

	variance -= 1
	randValue := (rand.Float64() - offset) * float64(variance)
	if randValue > 0 {
		randValue += 1
	} else {
		randValue -= 1
	}

	chosenIndex := int(randValue) + playerIndex

	if chosenIndex < 0 {
		chosenIndex = playerIndex + 1
	}

	if chosenIndex >= len(opponents) {
		chosenIndex = playerIndex - 1
	}

	log.Printf("XXX %d vs %d", playerIndex, chosenIndex)

	return opponents[chosenIndex].Value.(*EloPlayer)
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
