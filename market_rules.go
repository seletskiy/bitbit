package main

type MarketRules struct {
	Market         *Market
	BootstrapFunds float64
	BootstrapGoods float64
}

func (rules MarketRules) Apply(
	population *Population,
) {
	alive := 0
	for _, creature := range *population {
		if creature.Died() {
			rules.Market.Unregister(
				// @TODO: craete interface method called Free() on Energy
				creature.GetEnergy().(*MarketEnergy).MarketConnection.AgentId,
			)
		} else {
			alive++
		}
	}

	//freeActives := rules.Market.Actives
	//bonusFunds := freeActives.Funds / float64(alive)
	//bonusGoods := freeActives.Goods / float64(alive)
	for _, creature := range *population {
		if creature.Died() {
			continue
		}

		//target, _ := rules.Market.GetAgent(
		//    creature.GetEnergy().(*MarketEnergy).MarketConnection.AgentId,
		//)

		//err := rules.Market.Bootstrap(
		//    target,
		//    bonusFunds,
		//    bonusGoods,
		//)

		//if err != nil {
		//    break
		//    logger.Log(Debug, "BONE EATING ERROR: %s", err)
		//}
	}
}
