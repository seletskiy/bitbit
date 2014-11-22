package main

import (
	"fmt"
	"math"
)

var _ = fmt.Print

type Ticks []*Tick

type Tick struct {
	Ts    int
	Price float64
	Bid   float64
	Ask   float64
}

type Candles []*Candle

type Candle struct {
	Ts         int
	Period     int
	OpenPrice  float64
	ClosePrice float64
	Ticks      Ticks
}

func Aggregate(ticks Ticks, period int) Candles {
	var current = 0
	var result = make(Candles, 0)
	var candle *Candle = nil
	for _, t := range ticks {
		if t.Ts/period > current {
			current = t.Ts / period

			if candle != nil {
				result = append(result, candle)
			}

			candle = &Candle{
				Ts:         current * period,
				Period:     period,
				OpenPrice:  t.Price,
				ClosePrice: t.Price,
				Ticks:      make(Ticks, 0),
			}
		}

		candle.Ticks = append(candle.Ticks, t)
		candle.ClosePrice = t.Price
	}

	if candle != nil {
		return append(result, candle)
	} else {
		return result
	}
}

func RSI(candles Candles, MA func([]float64) float64) []float64 {
	u_seq := make([]float64, 0)
	d_seq := make([]float64, 0)

	rsi := make([]float64, 0)

	close_prev := candles[0].ClosePrice

	for _, c := range candles {
		diff := c.ClosePrice - close_prev

		u_seq = append(u_seq, math.Max(0, diff))
		d_seq = append(d_seq, math.Max(0, -diff))

		close_prev = c.ClosePrice
	}

	for _ = range candles {
		rs := MA(u_seq) / MA(d_seq)
		rsi = append(rsi, 100-100/(1+rs))
	}

	return rsi
}

func EMA(seq []float64, period int, alpha float64) float64 {
	if len(seq) == 1 {
		return alpha * seq[0]
	}

	if period <= 0 || len(seq) == 0 {
		return 0
	} else {
		last := len(seq) - 1
		return alpha*seq[last] +
			(1-alpha)*EMA(seq[:last-1], period-1, alpha)
	}
}

func PreSample(seq []float64, length int) []float64 {
	mean := 0.0

	for i, v := range seq {
		if i > length-1 {
			break
		}

		mean += v
	}

	fmt.Println(seq, mean)

	mean /= float64(length)

	prepend := make([]float64, length)
	for i := 0; i < length; i++ {
		prepend[i] = mean
	}

	return append(seq, prepend...)
}
