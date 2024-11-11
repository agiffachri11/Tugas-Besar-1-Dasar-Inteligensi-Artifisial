package localsearch

import (
	"diagonalmagiccube/cube"
)

func GeneticAlgorithm(generation *cube.Generation, maxIterations int) *cube.Generation {
	currentGeneration := generation
	cube.SearchBestAVGScore(currentGeneration)
	for i := 0; i < maxIterations; i++ {
		currentGeneration = cube.Evolution(currentGeneration)
		cube.SearchBestAVGScore(currentGeneration)
	}
	return currentGeneration
}
