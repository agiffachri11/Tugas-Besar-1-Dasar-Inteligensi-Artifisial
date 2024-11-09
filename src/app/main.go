package main

import (
	"diagonalmagiccube/cube"
	"diagonalmagiccube/localsearch"
	"fmt"
	"time"
)

func main() {
	c := cube.NewCube()
	fmt.Println("INITIAL STATE:")
	printState(c)

	maxIteration := 100000

	// Stochastic Hill-Climbing
	fmt.Printf("FINAL STATE (Stochastic Hill-Climbing with %d iterations):\n", maxIteration)
	executeSearch(localsearch.StochasticHillClimbing, c, maxIteration)

	// Simulated Annealing
	fmt.Printf("FINAL STATE (Simulated Annealing with %d iterations):\n", maxIteration)
	executeSearch(localsearch.SimulatedAnnealing, c, maxIteration)
}

func printState(c *cube.Cube) {
	fmt.Println("Cube Sequence:", c.GetSequence())
	fmt.Println("Objective Function Score:", c.GetScore())
	c.CountMagicSums()
	fmt.Println()
}

func executeSearch(searchFunc func(*cube.Cube, int) *cube.Cube, c *cube.Cube, maxIteration int) {
	start := time.Now()
	final := searchFunc(c, maxIteration)
	printState(final)
	fmt.Printf("Function took %s\n\n", time.Since(start))
}
