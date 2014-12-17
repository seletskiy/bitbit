package main

import (
	"container/list"
	"fmt"
	"math/rand"
	"strings"
)

type EloRules struct {
}

type EloRatings struct {
	*list.List
	TopCount          int
	StrongProbability float64
	OpponentVariance  float64
}

type EloPlayer struct {
	Player Creature
	Score  float64
}

func (ratings *EloRatings) Update(winner, looser *EloPlayer) {
	for mark := ratings.Front(); mark != nil; mark = mark.Next() {
		if mark.Score < winner.Score {
			ratings.MoveBefore(winner.Player, mark)
		}
	}

	for mark := ratings.Front(); mark != nil; mark = mark.Next() {
		if mark.Score > looser.Score {
			ratings.MoveAfter(looser.Player, mark)
		}
	}
}

func (ratings *EloRatings) String() string {
	result := make([]string, 0)

	for rating := ratings.Front(); rating != nil; rating = rating.Next() {
		if len(result) > ratings.TopCount {
			break
		}

		concrete := rating.Value.(*EloPlayer)
		result = append(
			result,
			fmt.Sprintf(
				"%4d %7.3f",
				concrete.Score,
				concrete.Player.GetEnergy().GetFloat64(),
			),
		)
	}

	return strings.Join(result, "\n")
}

func (ratings *EloRatings) ChooseOpponent(player *EloPlayer) *EloPlayer {
	index := 0
	opponents := make([]*list.Element, ratings.Len())
	for mark := ratings.Front(); mark != nil; mark = mark.Next() {
		opponents[index] = mark
		if mark == player {
			break
		} else {
			index++
		}
	}

	randValue := 2 * (rand.Float64() - ratings.StrongProbability) *
		ratings.OpponentVariance
	if randValue > 0 {
		randValue += 1
	} else {
		randValue -= 1
	}

	chosenIndex := int(randValue) + index

	if chosenIndex < 0 {
		chosenIndex = 0
	}

	if chosenIndex >= len(opponents) {
		chosenIndex = len(opponents) - 1
	}

	return opponents[chosenIndex]
}

func (rules EloRules) Apply(
	population *Population,
) {
}
