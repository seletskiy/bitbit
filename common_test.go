package main

import (
	"fmt"
	"testing"
)

var _ = fmt.Print

func TestAggregateEmptyList(t *testing.T) {
	if len(Aggregate(Ticks{}, 5)) != 0 {
		t.Fatal("aggregate on empty list failed")
	}
}

func TestAggregateOneItem(t *testing.T) {
	items := Ticks{&Tick{100, 10, 0, 0}}
	result := Aggregate(items, 5)

	if len(result) != 1 {
		t.Fatal("aggregate on one item result have len() != 1")
	}

	if result[0].OpenPrice != result[0].ClosePrice {
		t.Fatal("aggregate on one item have inequal open and close prices")
	}
}

func TestAggregateTwoFarItems(t *testing.T) {
	items := Ticks{
		&Tick{100, 10, 0, 0},
		&Tick{106, 15, 0, 0},
	}

	result := Aggregate(items, 5)

	if len(result) != 2 {
		t.Fatal("aggregate on two far items have len() != 2")
	}

	if result[0].OpenPrice != 10 {
		t.Fatal("aggregate on two items have invalid 0/OpenPrice")
	}

	if result[1].OpenPrice != 15 {
		t.Fatal("aggregate on two items have invalid 1/OpenPrice")
	}
}

func TestAggregate10To2(t *testing.T) {
	items := Ticks{
		&Tick{100, 10, 0, 0},
		&Tick{101, 12, 0, 0},
		&Tick{102, 5, 0, 0},
		&Tick{103, 16, 0, 0},
		&Tick{104, 4, 0, 0},

		&Tick{105, 4, 0, 0},
		&Tick{106, 10, 0, 0},
		&Tick{107, 12, 0, 0},
		&Tick{108, 11, 0, 0},
		&Tick{109, 5, 0, 0},
	}

	result := Aggregate(items, 5)

	if len(result) != 2 {
		t.Fatal("aggregate 10 to 2 have len() != 2")
	}

	if result[0].OpenPrice != 10 || result[0].ClosePrice != 4 {
		t.Fatal("aggregate 10 to 2 have invalid 0/OpenClosePrice")
	}

	if result[1].OpenPrice != 4 || result[1].ClosePrice != 5 {
		t.Fatal("aggregate 10 to 2 have invalid 1/OpenClosePrice")
	}

	if len(result[0].Ticks) != 5 || len(result[1].Ticks) != 5 {
		t.Fatal("aggregate 10 to 2 have invalid ticks count in periods")
	}
}

func TestRSI(t *testing.T) {
	candles := Candles{
		&Candle{1383671569, 300, 242.62, 243.01, Ticks{}},
		&Candle{1383671322, 900, 243.0, 243.01, Ticks{}},
		&Candle{1383671269, 300, 240.25, 242.62, Ticks{}},
		&Candle{1383670969, 300, 246.29, 240.25, Ticks{}},
		&Candle{1383670925, 3600, 246.3, 243.01, Ticks{}},
		&Candle{1383670669, 300, 245.96, 246.29, Ticks{}},
		&Candle{1383670422, 900, 247.97, 243.0, Ticks{}},
		&Candle{1383670369, 300, 247.97, 245.96, Ticks{}},
	}

	rsi := RSI(candles, func(seq []float64) float64 {
		seq2 := PreSample(seq, 5)
		fmt.Println(seq2)
		return EMA(seq2, 5, 0.8)
	})

	fmt.Println(rsi)

	t.Fatal(1)
}
