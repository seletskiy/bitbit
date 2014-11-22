package main

type Transaction struct {
	Price     float64
	Volume    float64
	Timestamp int64
}

type Market struct {
	Periods []Period
	Grain   int
}

type Period struct {
	Open      float64
	Close     float64
	High      float64
	Low       float64
	Timestamp int64
	Volume    int
}

func (m *Market) AddTransaction(t Transaction) {
	last := m.Periods[len(m.Periods)-1]

	if last.Timestamp+m.Grain >= t.Timestamp {
		// continue period
	} else {
		// new period
	}
}
