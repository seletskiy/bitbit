package main

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
)

type AgentId int64

func (id AgentId) String() string {
	return fmt.Sprintf("%d", int(id))
}

type Transaction struct {
	Price     float64
	Volume    float64
	Timestamp int64
}

type MarketConnection struct {
	AgentId AgentId
	Market  *Market
}

type MarketActives struct {
	Funds float64
	Goods float64
}

type MarketPosition struct {
	AgentId AgentId
	TTL     int
	Amount  *MarketActives
	Bought  bool
}

type Market struct {
	Actives   *MarketActives
	Agents    map[AgentId]*MarketAgent
	Periods   []*Period
	Grain     int64
	Positions []*MarketPosition
	BasePrice float64
}

type MarketAgent struct {
	Id          AgentId
	Actives     *MarketActives
	BaseActives *MarketActives
}

type Period struct {
	Open      float64
	Close     float64
	High      float64
	Low       float64
	Timestamp int64
	Volume    float64
}

func (market *Market) GetLastPeriod() *Period {
	var last *Period
	if len(market.Periods) > 0 {
		last = market.Periods[len(market.Periods)-1]
	}

	return last
}

func (market *Market) MarkBasePrice() {
	last := market.GetLastPeriod()
	if last == nil {
		return
	}

	market.BasePrice = last.Close
}

func (market *Market) GetBasePrice() float64 {
	return market.BasePrice
}

func (market *Market) GetMarketPrice() float64 {
	last := market.GetLastPeriod()
	if last == nil {
		return -1
	}

	return last.Close
}

func (market *Market) AddTransaction(t Transaction) {
	Log(Debug,
		"MARKET: new transaction %fF %fV",
		t.Price,
		t.Volume,
	)

	last := market.GetLastPeriod()

	if last != nil && last.Timestamp+market.Grain >= t.Timestamp {
		last.Close = t.Price
		last.Volume += t.Volume
		last.High = math.Max(last.High, t.Price)
		last.Low = math.Min(last.Low, t.Price)
	} else {
		market.Periods = append(market.Periods, &Period{
			Open:      t.Price,
			Close:     t.Price,
			High:      t.Price,
			Low:       t.Price,
			Timestamp: t.Timestamp,
			Volume:    t.Volume,
		})
	}

	if len(market.Periods) > 100 {
		market.Periods = market.Periods[1:]
	}
}

func (market *Market) GetAgent(id AgentId) (*MarketAgent, error) {
	if agent, ok := market.Agents[id]; !ok {
		return nil, fmt.Errorf(`agent does not exist`)
	} else {
		return agent, nil
	}
}

func (market *Market) RegisterFrom(
	id AgentId, funds, goods float64,
) (*MarketConnection, error) {
	agent, err := market.GetAgent(id)
	if err != nil {
		return nil, err
	}

	return market.registerAndTransfer(agent.Actives, funds, goods)
}

func (market *Market) Register(
	funds, goods float64,
) (*MarketConnection, error) {
	return market.registerAndTransfer(market.Actives, funds, goods)
}

func (market *Market) registerAndTransfer(
	sourceActives *MarketActives,
	funds, goods float64,
) (*MarketConnection, error) {
	agent := &MarketAgent{
		Id:      AgentId(rand.Int63()),
		Actives: &MarketActives{},
	}

	if funds+goods > 0 {
		err := sourceActives.Transfer(agent.Actives, funds, goods)
		if err != nil {
			return nil, err
		}
	}

	agent.BaseActives = &MarketActives{
		Funds: funds,
		Goods: goods,
	}

	Log(Debug,
		`AGENT<%s> ACTIVES<%p> register on market`,
		agent.Id,
		agent.Actives,
	)

	market.Agents[agent.Id] = agent

	return &MarketConnection{
		AgentId: agent.Id,
		Market:  market,
	}, nil

}

func (market *Market) Unregister(id AgentId) error {
	agent, err := market.GetAgent(id)
	if err != nil {
		return err
	}

	agent.Actives.Transfer(
		market.Actives, agent.Actives.Funds, agent.Actives.Goods,
	)

	openedPositions := []*MarketPosition{}
	for _, position := range market.Positions {
		if position.AgentId != agent.Id {
			openedPositions = append(openedPositions, position)
			continue
		}

		position.Amount.Transfer(
			market.Actives,
			position.Amount.Funds,
			position.Amount.Goods,
		)
	}

	market.Positions = openedPositions

	Log(Debug, `AGENT<%s> unregister on market`, id)

	delete(market.Agents, id)

	return nil
}

func (market *Market) PlacePosition(
	id AgentId, ttl int, amount float64,
) error {
	agent, err := market.GetAgent(id)
	if err != nil {
		return err
	}

	if amount < 1e-6 {
		return errors.New("price and amount can't be below zero")
	}

	position := &MarketPosition{
		AgentId: id,
		Amount:  &MarketActives{},
	}

	err = agent.Actives.Transfer(position.Amount, amount, 0)
	if err != nil {
		return err
	}

	Log(Debug,
		"MARKET: AGENT<%s> open position: %fF TTL %d",
		id, amount, ttl,
	)

	market.Positions = append(market.Positions, position)

	return nil
}

func (market *Market) Simulate() error {
	last := market.GetLastPeriod()

	if last == nil {
		return errors.New("no transactions made")
	}

	marketPrice := market.GetMarketPrice()

	openPositions := []*MarketPosition{}
	for _, position := range market.Positions {
		if !market.evaluatePosition(position, marketPrice) {
			openPositions = append(openPositions, position)
			Log(Debug,
				"MARKET: open %fF",
				position.Amount.Funds,
			)
		}
	}

	//market.Positions = openPositions

	return nil
}

func (market *Market) evaluatePosition(
	position *MarketPosition, price float64,
) bool {
	if position.TTL > 0 {
		position.TTL--
		return false
	}

	agent, _ := market.GetAgent(position.AgentId)
	if position.Bought {
		position.Amount.Funds = position.Amount.Goods * price
		position.Amount.Transfer(
			agent.Actives,
			position.Amount.Funds,
			position.Amount.Goods,
		)

		Log(Debug, "MARKET POSITION: closed %fG -> %fF (for %fF)",
			position.Amount.Funds,
			position.Amount.Goods,
			price,
		)

		return true
	} else {
		position.Amount.Goods = position.Amount.Funds / price
		position.Bought = true

		Log(Debug, "MARKET POSITION: entered %fF -> %fG (for %fF)",
			position.Amount.Funds,
			position.Amount.Goods,
			price,
		)

		return false
	}
}

func (market *Market) GetHigh() float64 {
	last := market.GetLastPeriod()
	if last == nil {
		return 0
	} else {
		return last.High
	}
}

func (market *Market) GetLow() float64 {
	last := market.GetLastPeriod()
	if last == nil {
		return 0
	} else {
		return last.Low
	}
}

func (actives *MarketActives) Transfer(
	target *MarketActives, funds, goods float64,
) error {
	if actives.Funds < funds {
		return fmt.Errorf(
			`not enough funds to transfer: %f < %f`,
			actives.Funds, funds,
		)
	}

	if actives.Goods < goods {
		return fmt.Errorf(
			`not enough goods to transfer: %f < %f`,
			actives.Goods, goods,
		)
	}

	if funds+goods <= 0 {
		return fmt.Errorf(
			`can't transfer actives <= 0: %fF %fG`,
			funds, goods,
		)
	}

	Log(Debug,
		"ACTIVES<%p> -> ACTIVES<%p> market transfer: %fF %fG",
		actives,
		target, funds, goods,
	)

	actives.Funds -= funds
	actives.Goods -= goods

	target.Funds += funds
	target.Goods += goods

	return nil
}

func (market *Market) GetGoods(id AgentId) (float64, error) {
	agent, err := market.GetAgent(id)
	if err != nil {
		return 0, err
	}

	return agent.Actives.Goods, nil
}

func (market *Market) GetFunds(id AgentId) (float64, error) {
	agent, err := market.GetAgent(id)
	if err != nil {
		return 0, err
	}

	return agent.Actives.Funds, nil
}

func (market *Market) Revoke(id AgentId, funds, goods float64) error {
	agent, err := market.GetAgent(id)
	if err != nil {
		return err
	}

	return agent.Actives.Transfer(market.Actives, funds, goods)
}

func (market *Market) GetPositionsFor(id AgentId) []*MarketPosition {
	positions := []*MarketPosition{}

	for _, position := range market.Positions {
		if position.AgentId != id {
			continue
		}

		positions = append(positions, position)
	}

	return positions
}

func (market *Market) GetTotalActives() (float64, float64) {
	totalFunds := market.Actives.Funds
	totalGoods := market.Actives.Goods
	for _, position := range market.Positions {
		totalFunds += position.Amount.Funds
		totalGoods += position.Amount.Goods
	}

	for _, agent := range market.Agents {
		totalFunds += agent.Actives.Funds
		totalGoods += agent.Actives.Goods
	}

	return totalFunds, totalGoods
}

func (market *Market) String() string {
	var last *Period
	if len(market.Periods) > 0 {
		last = market.Periods[len(market.Periods)-1]
	}

	totalFunds, totalGoods := market.GetTotalActives()
	if last == nil {
		return "MARKET: nil"
	} else {
		return fmt.Sprintf(
			"MARKET:\nLAST %fF -- %fF\nWALLET %fF %fG ~ %fF (free %fF %fG)",
			last.Low,
			last.High,
			totalFunds,
			totalGoods,
			totalFunds+totalGoods*market.BasePrice,
			market.Actives.Funds,
			market.Actives.Goods,
		)
	}
}

func (connection *MarketConnection) Close() {
	connection.Market.Unregister(connection.AgentId)
}

func (connection *MarketConnection) GetFreeGoods() float64 {
	amount, _ := connection.Market.GetGoods(connection.AgentId)

	return amount
}

func (connection *MarketConnection) GetFreeFunds() float64 {
	amount, _ := connection.Market.GetFunds(connection.AgentId)

	return amount
}

func (connection *MarketConnection) GetActives() (funds, goods float64) {
	funds = connection.GetFreeFunds()
	goods = connection.GetFreeGoods()

	positions := connection.Market.GetPositionsFor(connection.AgentId)
	for _, position := range positions {
		goods += position.Amount.Goods
		funds += position.Amount.Funds
	}

	return funds, goods
}

func (connection *MarketConnection) Transfer(
	id AgentId, funds, goods float64,
) error {
	target, err := connection.Market.GetAgent(id)
	if err != nil {
		return err
	}

	source, err := connection.Market.GetAgent(connection.AgentId)
	if err != nil {
		return err
	}

	return source.Actives.Transfer(target.Actives, funds, goods)
}

func (connection *MarketConnection) GetMarketPrice() float64 {
	return connection.Market.GetMarketPrice()
}

func (connection *MarketConnection) GetBasePrice() float64 {
	return connection.Market.GetBasePrice()
}

func (connection *MarketConnection) PlacePosition(
	amount float64, ttl int,
) error {
	return connection.Market.PlacePosition(
		connection.AgentId,
		ttl,
		amount,
	)
}
