package main

import (
	"fmt"
	"math"
)

type MarketEnergy struct {
	*MarketConnection
	MinSplitFunds float64
}

func (energy *MarketEnergy) TransferTo(target Energy) {
	targetId := target.(*MarketEnergy).MarketConnection.AgentId
	energy.MarketConnection.Transfer(
		targetId,
		energy.MarketConnection.GetFreeFunds(),
		energy.MarketConnection.GetFreeGoods(),
	)
}

func (energy *MarketEnergy) Scatter(n int) []Energy {
	// @TODO: to be done
	return nil
}

func (energy *MarketEnergy) Split() Energy {
	if energy.GetTotalFunds() < energy.MinSplitFunds {
		return nil
	}

	Log(Debug, "ENERGY<%p> split: %s", energy, energy)

	connection := energy.MarketConnection

	newConnection, err := connection.Market.RegisterFrom(
		energy.AgentId,
		connection.GetFreeFunds()/2,
		connection.GetFreeGoods()/2,
	)

	if err != nil {
		panic(err)
		return nil
	}

	return &MarketEnergy{
		MarketConnection: newConnection,
		MinSplitFunds:    energy.MinSplitFunds,
	}
}

func (energy *MarketEnergy) Void() bool {
	funds, goods := energy.MarketConnection.GetActives()

	return math.Abs(funds+goods) < 1e-6
}

func (energy *MarketEnergy) Exceed(target Energy) bool {
	//opponent := target.(*MarketEnergy)
	//return energy.Profitability > opponent.Profitability
	return false
}

func (energy *MarketEnergy) GetTotalFunds() float64 {
	myFunds, myGoods := energy.MarketConnection.GetActives()
	return myFunds + myGoods*energy.MarketConnection.GetBasePrice()
}

func (energy *MarketEnergy) String() string {
	connection := energy.MarketConnection
	agent, _ := connection.Market.GetAgent(
		energy.MarketConnection.AgentId,
	)
	funds, goods := connection.GetActives()
	return fmt.Sprintf(
		`%fF %fG ~ %fF AGENT<%d> ACTIVES<%p>`,
		connection.GetFreeFunds(),
		connection.GetFreeGoods(),
		funds+goods*connection.GetBasePrice(),
		connection.AgentId,
		agent.Actives,
	)
}

func (energy *MarketEnergy) Simulate() {

}

func (energy *MarketEnergy) Preveal() float64 {
	return 0
}

func (energy *MarketEnergy) Free() {
	// noop
}

func (energy *MarketEnergy) GetFloat64() float64 {
	return 0
}
