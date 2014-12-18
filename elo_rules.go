package main

import (
	"math"
	"sort"
)

type EloRules struct {
	DrawThreshold     float64
	TopCount          int
	StrongProbability float64
	OpponentVariance  float64
}

func (rules EloRules) Apply(
	population *Population,
) {
	sort.Sort(ByEnergy(*population))

	eloPlayers := make([]*EloPlayer, len(*population))
	for i, creature := range *population {
		eloPlayers[i] = &EloPlayer{
			Player: creature,
			Score:  creature.GetEnergy().(EloBasedEnergy).GetEloScore(),
		}
	}

	eloRatings := NewEloRatings(eloPlayers)
	eloRatings.StrongProbability = rules.StrongProbability
	eloRatings.OpponentVariance = rules.OpponentVariance

	for _, player := range eloPlayers {
		versus := eloRatings.ChooseOpponent(player)

		playerEnergy := player.Player.(Creature).GetEnergy()
		versusEnergy := versus.Player.(Creature).GetEnergy()

		Log(Debug,
			"ELO: CREATURE<%p> %4d (%f) vs CREATURE<%p> %4d (%f)",
			player.Player,
			playerEnergy.(EloBasedEnergy).GetEloScore(),
			playerEnergy.GetFloat64(),
			versus.Player,
			versusEnergy.(EloBasedEnergy).GetEloScore(),
			versusEnergy.GetFloat64(),
		)

		playerEnergyValue := playerEnergy.GetFloat64()
		versusEnergyValue := versusEnergy.GetFloat64()

		energyValueDiff := math.Abs(1.0 - playerEnergyValue/versusEnergyValue)

		if energyValueDiff > rules.DrawThreshold {
			if playerEnergyValue > versusEnergyValue {
				player.Score = eloRatings.Compute(player, versus, EloMatchWin)
				versus.Score = eloRatings.Compute(versus, player, EloMatchLoss)

				Log(Debug,
					"ELO: CREATURE<%p> WIN -> %d",
					player.Player,
					player.Score,
				)
			} else {
				player.Score = eloRatings.Compute(player, versus, EloMatchLoss)
				versus.Score = eloRatings.Compute(versus, player, EloMatchWin)

				Log(Debug,
					"ELO: CREATURE<%p> LOSS -> %d",
					player.Player,
					player.Score,
				)
			}
		} else {
			player.Score = eloRatings.Compute(player, versus, EloMatchDraw)
			versus.Score = eloRatings.Compute(versus, player, EloMatchDraw)
			Log(Debug, "ELO: CREATURE<%p> DRAW -> %d", player, player.Score)
		}

		eloRatings.Update(player)
		eloRatings.Update(versus)

		player.Player.(Creature).GetEnergy().(EloBasedEnergy).SetEloScore(
			player.Score,
		)

		versus.Player.(Creature).GetEnergy().(EloBasedEnergy).SetEloScore(
			versus.Score,
		)
	}

	index := 0
	for ptr := eloRatings.Front(); ptr != nil; ptr = ptr.Next() {
		//if index >= rules.TopCount {
		//    break
		//}

		Log(Debug,
			"ELO: TOP | %4d  %f  CREATURE<%p>",
			int(ptr.Value.(*EloPlayer).Score),
			ptr.Value.(*EloPlayer).Player.(Creature).GetEnergy().GetFloat64(),
			ptr.Value.(*EloPlayer).Player,
		)
		index++
	}
}
