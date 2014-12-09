package main

type Rules interface {
	Apply(population *Population)
}
