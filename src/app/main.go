package main

import (
	"bufio"
	"diagonalmagiccube/cube"
	"diagonalmagiccube/localsearch"
	"fmt"
	"os"
	"time"
)

func main() {
	// Get maxIteration from user input
	var maxIteration int
	fmt.Print("Enter the number of iterations: ")
	_, err := fmt.Scanln(&maxIteration)
	if err != nil {
		fmt.Println("Invalid input:", err)
		return
	}

	// Stochastic Hill-Climbing
	cubeA := cube.NewCube()
	fmt.Println("INITIAL STATE:")
	printState(cubeA)
	fmt.Printf("FINAL STATE (Stochastic Hill-Climbing with %d iterations):\n", maxIteration)
	executeSearch(localsearch.StochasticHillClimbing, cubeA, maxIteration)
	plotObjectiveFunction(cubeA, maxIteration/10)

	// Simulated Annealing
	cubeB := cube.NewCube()
	fmt.Println("INITIAL STATE:")
	printState(cubeB)
	fmt.Printf("FINAL STATE (Simulated Annealing with %d iterations):\n", maxIteration)
	executeSearch(localsearch.SimulatedAnnealing, cubeB, maxIteration)
	plotObjectiveFunction(cubeB, maxIteration/10)

	// Genetic Algorithm
	// generation := cube.NewGeneration()
	// cube.GenerationDetail(generation)
	// printState(cube.BestIndividual(generation).GetCube())
	// start := time.Now()
	// final := localsearch.GeneticAlgorithm(generation, maxIteration)
	// end := time.Since(start)
	// cube.GenerationDetail(final)
	// printState(cube.BestIndividual(final).GetCube())
	// fmt.Printf("Function took %s\n\n", end)

	// SAVING FILE FOR UNITY3D
	// cube.SaveCubeToFile(cubeB, "cube.json")

	pressToContinue()
}

func printState(c *cube.Cube) {
	fmt.Println("Cube Sequence:", c.GetSequence())
	fmt.Println("Objective Function Score:", c.GetScore())
	c.CountMagicSums()
	fmt.Println()
}

func executeSearch(searchFunc func(*cube.Cube, int, *int) *cube.Cube, c *cube.Cube, maxIteration int) {
	stuckCount := 0
	start := time.Now()
	final := searchFunc(c, maxIteration, &stuckCount)
	end := time.Since(start)
	printState(final)
	fmt.Println("Stuck Freq: ", stuckCount)
	fmt.Printf("Function took %s\n\n", end)
}

func pressToContinue() {
	fmt.Print("Press Enter to continue...")
	fmt.Scanln() // waits for the user to press Enter
}

func plotObjectiveFunction(c *cube.Cube, interval int) {
	// Create or open the file for writing
	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	// Ensure the file gets closed after writing
	defer file.Close()

	// Create a buffered writer to write to the file
	writer := bufio.NewWriter(file)

	var i int = 0
	current := c
	for current != nil {
		if (i % interval) == 0 {
			// Write to the file instead of printing to console
			writer.WriteString(fmt.Sprintf("%d %d\n", i, current.GetScore()))
		}
		current = current.GetSuccessor()
		i++
	}

	// Flush the buffered writer to ensure everything is written to the file
	writer.Flush()
}
