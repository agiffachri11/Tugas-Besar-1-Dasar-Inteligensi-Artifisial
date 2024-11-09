package main

import (
	"diagonalmagiccube/cube"
	"diagonalmagiccube/localsearch"
	"fmt"
)

func main() {
	c := cube.NewCube()

	// Initial State
	fmt.Println("INITIAL STATE:")
	fmt.Println("Cube Sequence:", c.GetSequence())
	fmt.Println("Objective Function Score:", c.GetScore())
	c.CountMagicSums()
	fmt.Println("")

	var maxIteration int = 100000

	// Final State - Stochastic Hill-Climbing
	fmt.Printf("FINAL STATE (Stochastic Hill-Climbing with %d iterations):\n", maxIteration)
	final1 := localsearch.StochasticHillClimbing(c, maxIteration)
	fmt.Println("Cube Sequence:", final1.GetSequence())
	fmt.Println("Objective Function Score:", final1.GetScore())
	final1.CountMagicSums()
	fmt.Println("")

	// Final State - Simulated Annealing
	fmt.Printf("FINAL STATE (Simulated Annealing with %d iterations):\n", maxIteration)
	final2 := localsearch.SimulatedAnnealing(c, maxIteration)
	fmt.Println("Cube Sequence:", final2.GetSequence())
	fmt.Println("Objective Function Score:", final2.GetScore())
	final2.CountMagicSums()
	fmt.Println("")
}
