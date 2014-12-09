package main

func main() {
	////rand.Seed(time.Now().UnixNano())

	//log.SetOutput(os.Stdout)

	//csvFile, err := os.Open("data-inf.csv")
	//if err != nil {
	//    panic(err)
	//}

	//market := &Market{
	//    Grain: 1,
	//    Actives: &MarketActives{
	//        Funds: initialFunds,
	//        Goods: initialGoods,
	//    },
	//    Agents: map[AgentId]*MarketAgent{},
	//}

	//programLayout := RandProgramLayout(programLength)

	//programInstructionVariants := []RandInstructionVariant{
	//    {&ProgramInstructionAdd{}, 1.0},
	//    {&ProgramInstructionMov{}, 1.0},
	//    {&ProgramInstructionDiv{}, 0.1},
	//    {&ProgramInstructionNop{}, 1.0},

	//    //{&ProgramInstructionSell{}, 0.5},
	//    //{&ProgramInstructionBuy{}, 0.5},

	//    {&ProgramInstructionPosition{}, 0.5},

	//    {&ProgramInstructionFree{}, 0.5},
	//    {&ProgramInstructionMarket{}, 0.5},

	//    {&ProgramInstructionJumpGreaterThan{}, 0.5},
	//}

	//population := make(Population, initialPopulationSize)
	//for i := 0; i < initialPopulationSize; i++ {
	//    program := RandProgram(
	//        programLayout,
	//        programReferenceProbability,
	//        programValueVariance,
	//        programMemorySize,
	//        programInstructionVariants,
	//    )

	//    connection, err := market.Register(
	//        bootstrapFunds,
	//        bootstrapGoods,
	//    )

	//    if err != nil {
	//        panic(err)
	//    }

	//    energy := &MarketEnergy{
	//        MarketConnection: connection,
	//        MinSplitFunds:    bootstrapFunds * 2,
	//    }

	//    population[i] = RandProgoBact(
	//        programMemorySize,
	//        program,
	//        energy,
	//    )
	//}

	//environment := SimpleEnvironment{
	//    Rules: []Rules{
	//        ReproduceRules{
	//            ReproduceProbability: 0.05,
	//            MinReproduceAge:      20,
	//        },

	//        MutateRules{
	//            DNAMutationProbability: 0.5,
	//            MutationMaxSize:     2,
	//            GeneGenerator: func(amount int) []Gene {
	//                instructions := RandProgramInstructionSet(
	//                    amount,
	//                    programReferenceProbability,
	//                    programValueVariance,
	//                    programLength,
	//                    programInstructionVariants,
	//                )

	//                result := make([]Gene, amount)
	//                for index, instruction := range instructions {
	//                    result[index] = Codepoint{
	//                        Instruction: instruction,
	//                    }
	//                }

	//                return result
	//            },
	//        },

	//        BacterialGeneTransferRules{
	//            BirthTransferProbability:      0.5,
	//            ReproduceLossProbability:      0.5,
	//            TransferLossProbability:       0.05,
	//            ApplyLossProbability:          0.2,
	//            MaxPlasmidsNumber:             4,
	//            PlasmidPerAge:                 15,
	//            ExchangePlasmidsProbability:   0.01,
	//            MinAgeForExchange:             30,
	//            PlasmidPrefixLengthProportion: 0.3,
	//        },

	//        MarketRules{
	//            Market:         market,
	//            BootstrapFunds: 10.0,
	//        },

	//        NaturalSelectionRules{
	//            KillPossibility: 0.055,
	//        },

	//        //AgingRules{
	//        //    DieAge: 100,
	//        //},

	//        ReapRules{},
	//    },
	//}

	//tick := int64(0)
	//for len(population) > 0 {
	//    csvReader := csv.NewReader(csvFile)
	//    for len(population) > 0 {
	//        data, err := csvReader.Read()
	//        if err != nil {
	//            break
	//        }

	//        openPrice, _ := strconv.ParseFloat(data[1], 64)
	//        closePrice, _ := strconv.ParseFloat(data[4], 64)
	//        volume, _ := strconv.ParseFloat(data[5], 64)

	//        prices := []float64{
	//            openPrice,
	//            closePrice,
	//        }

	//        for subtick := 0; subtick < 2; subtick++ {
	//            logger.Log(Debug, "SIMULATION: tick start %d", tick)

	//            market.AddTransaction(Transaction{
	//                Price:     prices[subtick],
	//                Volume:    volume / 2,
	//                Timestamp: tick,
	//            })

	//            environment.Simulate(&population)

	//            market.Simulate()

	//            tick++
	//        }

	//    }
	//}
}
