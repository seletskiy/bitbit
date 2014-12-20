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
	//sort.Sort(ByEnergyReverse(*population))
	sort.Sort(ByEloScore(*population))

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

	for index, _ := range eloPlayers {
		player := eloPlayers[len(eloPlayers)-index-1]
		versus := eloRatings.ChooseOpponent(player)

		playerEnergy := player.Player.(Creature).GetEnergy()
		versusEnergy := versus.Player.(Creature).GetEnergy()

		playerEnergyValue := playerEnergy.GetFloat64()
		versusEnergyValue := versusEnergy.GetFloat64()

		energyValueDiff := math.Abs(1.0 - playerEnergyValue/versusEnergyValue)

		newPlayerScore := 0
		newVersusScore := 0
		if energyValueDiff > rules.DrawThreshold {
			if playerEnergyValue > versusEnergyValue {
				newPlayerScore = eloRatings.Compute(player, versus, EloMatchWin)
				newVersusScore = eloRatings.Compute(versus, player, EloMatchLoss)

				Log(Debug,
					"ELO: <%p> %4d->%4d ^%d (%12.5g) WIN  <%p> %4d->%4d ^%d (%12.5g)",
					player.Player,
					player.Score,
					newPlayerScore,
					playerEnergy.(EloBasedEnergy).GetWinsInRow(),
					playerEnergy.GetFloat64(),
					versus.Player,
					versus.Score,
					newVersusScore,
					versusEnergy.(EloBasedEnergy).GetWinsInRow(),
					versusEnergy.GetFloat64(),
				)
			} else {
				newPlayerScore = eloRatings.Compute(player, versus, EloMatchLoss)
				newVersusScore = eloRatings.Compute(versus, player, EloMatchWin)

				Log(Debug,
					"ELO: <%p> %4d->%4d ^%d (%12.5g) LOSS <%p> %4d->%4d ^%d (%12.5g)",
					player.Player,
					player.Score,
					newPlayerScore,
					playerEnergy.(EloBasedEnergy).GetWinsInRow(),
					playerEnergy.GetFloat64(),
					versus.Player,
					versus.Score,
					newVersusScore,
					versusEnergy.(EloBasedEnergy).GetWinsInRow(),
					versusEnergy.GetFloat64(),
				)
			}
		} else {
			newPlayerScore = eloRatings.Compute(player, versus, EloMatchDraw)
			newVersusScore = eloRatings.Compute(versus, player, EloMatchDraw)

			Log(Debug,
				"ELO: <%p> %4d->%4d ^%d (%12.5g) DRAW <%p> %4d->%4d ^%d (%12.5g)",
				player.Player,
				player.Score,
				newPlayerScore,
				playerEnergy.(EloBasedEnergy).GetWinsInRow(),
				playerEnergy.GetFloat64(),
				versus.Player,
				versus.Score,
				newVersusScore,
				versusEnergy.(EloBasedEnergy).GetWinsInRow(),
				versusEnergy.GetFloat64(),
			)
		}

		player.Score = newPlayerScore
		versus.Score = newVersusScore

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

		index++

		Log(Debug,
			"ELO: TOP %4d) %4d  %15.5g  CREATURE<%p> ^%d",
			index,
			int(ptr.Value.(*EloPlayer).Score),
			ptr.Value.(*EloPlayer).Player.(Creature).GetEnergy().GetFloat64(),
			ptr.Value.(*EloPlayer).Player,
			ptr.Value.(*EloPlayer).Player.(Creature).GetEnergy().(EloBasedEnergy).GetWinsInRow(),
		)
	}

	sort.Sort(ByEloScore(*population))
}
