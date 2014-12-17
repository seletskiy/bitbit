package main

import "container/list"

type EloRules struct {
}

type EloRatings struct {
	*list.List
}

type EloPlayer struct {
	Player Creature
	Score  float64
}

func (ratings *EloRatings) Update(winner, looser EloPlayer) {
	for mark := ratings.Front(); mark != nil; mark = mark.Next() {
		if mark.Score < winner.Score {
			ratings.MoveBefore(winner.Player, mark)
		}
	}
	ratings.Remove(winner.Player)
	ratings.Remove(looser.Score)
}

func (rules EloRules) Apply(
	population *Population,
) {
}
