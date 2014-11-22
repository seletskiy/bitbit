package main

//import (
//    "encoding/csv"
//    "io"
//    "os"
//)

//type Simulation struct {
//    File       *os.File
//    Source     *io.Reader
//    Population Population
//}

//func NewSimulation(string FileName) *Simulation {
//    file, err := os.OpenFile(FileName)
//    if err != nil {
//        panic(err)
//    }

//    return &Simulation{
//        File:   file,
//        Source: csv.NewReader(file),
//    }
//}

//func (Simulation *sim) Run() {

//}

//func (Simulation *sim) RunCycle() {
//    sim.Source.Reset()

//    market := NewMarket(sim.Grain)

//    for {
//        market.AddTransaction(sim.Source.GetNextTransaction())

//        for _, strategy := range sim.Population {
//            if strategy.Alive() {
//                strategy.Eval(market)
//            }
//        }

//        for _, strategy := range sim.Population {
//            sim.Reaper.Hunt(strategy)
//        }
//    }
//}
