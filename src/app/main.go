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

	// Get maxIteration from user input
	var maxIteration int
	fmt.Print("Enter the number of iterations: ")
	_, err := fmt.Scanln(&maxIteration)
	if err != nil {
		fmt.Println("Invalid input:", err)
		return
	}

	// Stochastic Hill-Climbing
	fmt.Printf("FINAL STATE (Stochastic Hill-Climbing with %d iterations):\n", maxIteration)
	executeSearch(localsearch.StochasticHillClimbing, c, maxIteration)

	// Simulated Annealing
	fmt.Printf("FINAL STATE (Simulated Annealing with %d iterations):\n", maxIteration)
	executeSearch(localsearch.SimulatedAnnealing, c, maxIteration)

	pressToContinue()
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
	end := time.Since(start)
	printState(final)
	fmt.Println("Stuck freq: ", stuckFrequency(c))
	fmt.Printf("Function took %s\n\n", end)
}

func stuckFrequency(c *cube.Cube) int {
	var current *cube.Cube = c
	var stuckCount int = 0
	for current.GetSuccessor() != nil {
		if current.GetScore() <= current.GetSuccessor().GetScore() {
			stuckCount++
		}
		current = current.GetSuccessor()
	}
	return stuckCount
}

func pressToContinue() {
	fmt.Print("Press Enter to continue...")
	fmt.Scanln() // waits for the user to press Enter
}
