package main

import (
	"bufio"
	"diagonalmagiccube/cube"
	"diagonalmagiccube/localsearch"
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	// Get maxIteration from user input
	var option int
	fmt.Print("1. Stochastic Hill Climbing\n2. Simulated Annealing\n3. Genetic Algorithm\nChoose algorithm:")
	_, err := fmt.Scanln(&option)
	if err != nil {
		fmt.Println("Invalid input:", err)
		return
	}

	var maxIteration int
	fmt.Print("Enter the number of iterations: ")
	_, err = fmt.Scanln(&maxIteration)
	if err != nil {
		fmt.Println("Invalid input:", err)
		return
	}

	if option == 1 {
		// Stochastic Hill-Climbing
		cubeA := cube.NewCube()
		fmt.Println("INITIAL STATE:")
		printState(cubeA)
		fmt.Printf("FINAL STATE (Stochastic Hill-Climbing with %d iterations):\n", maxIteration)
		executeSearch(localsearch.StochasticHillClimbing, cubeA, maxIteration)
		// SAVING FILE FOR UNITY3D
		cube.SaveCubeToFile(cubeA, "./cubevisual/Cube_Data/StreamingAssets/cube.json")
		StartVisualization()
	} else if option == 2 {
		// Simulated Annealing
		cubeB := cube.NewCube()
		fmt.Println("INITIAL STATE:")
		printState(cubeB)
		fmt.Printf("FINAL STATE (Simulated Annealing with %d iterations):\n", maxIteration)
		executeSearch(localsearch.SimulatedAnnealing, cubeB, maxIteration)
		// SAVING FILE FOR UNITY3D
		cube.SaveCubeToFile(cubeB, "./cubevisual/Cube_Data/StreamingAssets/cube.json")
		StartVisualization()
	} else if option == 3 {
		// Genetic Algorithm
		generation := cube.NewGeneration()
		cube.GenerationDetail(generation)
		printState(cube.BestIndividual(generation).GetCube())
		start := time.Now()
		final := localsearch.GeneticAlgorithm(generation, maxIteration)
		end := time.Since(start)
		// plotGeneticAlgorithm(generation, maxIteration/1000)
		cube.GenerationDetail(final)
		printState(cube.BestIndividual(final).GetCube())
		fmt.Printf("Function took %s\n\n", end)
	}

	pressToContinue()
}

func printState(c *cube.Cube) {
	fmt.Println("Cube Sequence:", c.GetSequence())
	fmt.Println("Objective Function Score:", c.GetScore())
	c.CountMagicSums()
	fmt.Println()
}

func executeSearch(searchFunc func(*cube.Cube, int, *int, []float64) *cube.Cube, c *cube.Cube, maxIteration int) {
	sliceDeltaE := make([]float64, maxIteration)
	stuckCount := 0
	start := time.Now()
	final := searchFunc(c, maxIteration, &stuckCount, sliceDeltaE)
	end := time.Since(start)
	printState(final)
	fmt.Println("Stuck Freq: ", stuckCount)
	fmt.Printf("Function took %s\n\n", end)
	// plotObjectiveFunction(c, maxIteration/1000)
	// plotProbabilitySA(sliceDeltaE, maxIteration/1000)
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

func plotProbabilitySA(sliceDeltaE []float64, interval int) {
	// Create or open the file for writing
	file, err := os.Create("outputDeltaE.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	// Ensure the file gets closed after writing
	defer file.Close()

	// Create a buffered writer to write to the file
	writer := bufio.NewWriter(file)

	for i := 0; i < len(sliceDeltaE); i++ {
		if ((i + 1) % interval) == 0 {
			// Write to the file instead of printing to console
			writer.WriteString(fmt.Sprintf("%d %f\n", i+1, sliceDeltaE[i]))
		}
	}

	// Flush the buffered writer to ensure everything is written to the file
	writer.Flush()
}

func plotGeneticAlgorithm(g *cube.Generation, interval int) {
	// Create or open the file for writing
	file, err := os.Create("outputGenetic.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	// Ensure the file gets closed after writing
	defer file.Close()

	// Create a buffered writer to write to the file
	writer := bufio.NewWriter(file)

	var i int = 0
	current := g
	for current != nil {
		if (i % interval) == 0 {
			// Write to the file instead of printing to console
			writer.WriteString(fmt.Sprintf("%d %d %d\n", i, current.GetAVGScore(), current.GetBestScore()))
		}
		current = current.GetNextGeneration()
		i++
	}

	// Flush the buffered writer to ensure everything is written to the file
	writer.Flush()
}

func StartVisualization() {
	// Path to the main.exe file
	executablePath := "./cubevisual/Cube.exe" // Use full path if it's not in the same directory

	// Create the command to execute main.exe
	cmd := exec.Command(executablePath)

	// Run the executable and wait for it to finish
	if err := cmd.Run(); err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("main.exe ran successfully.")
}
