package cube

import (
	"math/rand"
)

type Generation struct {
	population     [POPULATION_SIZE]*Individual
	nextGeneration *Generation
	totalFitness   float64
	bestScore      int
	avgScore       int
}

type Individual struct {
	cube     *Cube
	parentX  *Individual
	parentY  *Individual
	mutation bool
}

func (g *Generation) GetPopulation() [POPULATION_SIZE]*Individual {
	return g.population
}

func (g *Generation) GetNextGeneration() *Generation {
	return g.nextGeneration
}

func (g *Generation) GetTotalFitness() float64 {
	return g.totalFitness
}

func (g *Generation) GetBestScore() int {
	return g.bestScore
}
func (g *Generation) GetAVGScore() int {
	return g.avgScore
}

func (g *Generation) SetPopulation(population [POPULATION_SIZE]*Individual) {
	g.population = population
}

func (g *Generation) SetNextGeneration(nextGeneration *Generation) {
	g.nextGeneration = nextGeneration
}

func (i *Individual) GetCube() *Cube {
	return i.cube
}

func (i *Individual) GetParentX() *Individual {
	return i.parentX
}

func (i *Individual) GetParentY() *Individual {
	return i.parentY
}

func (i *Individual) SetCube(cube *Cube) {
	i.cube = cube
}

func (i *Individual) SetParentX(parentX *Individual) {
	i.parentX = parentX
}

func (i *Individual) SetParentY(parentY *Individual) {
	i.parentY = parentY
}

// NewGeneration constructor initializes a generation with a random filled population of POPULATION_SIZE individual
func NewGeneration() *Generation {
	// Create a new Generation instance nil nextGeneration
	generation := &Generation{nextGeneration: nil}
	// Populate with individuals
	for i := 0; i < POPULATION_SIZE; i++ {
		generation.population[i] = NewIndividual()
		generation.totalFitness += fitness(generation.population[i])
	}
	return generation
}

// NewIndividual constructor initializes an individual with a random filled cube and nil parents
func NewIndividual() *Individual {
	// Create a new Generation instance nil nextGeneration
	individual := &Individual{cube: NewCube(), parentX: nil, parentY: nil, mutation: false}
	return individual
}

func fitness(individual *Individual) float64 {
	return 10000 / (float64(individual.cube.score) + 1)
}

func selection(generation *Generation) *Individual {
	// // Ensure total fitness is not zero
	// if generation.totalFitness == 0 {
	// 	return generation.population[rand.Intn(POPULATION_SIZE)] // Randomly pick an individual
	// }

	// Select a random point on the cumulative wheel
	randValue := rand.Float64()
	var cumulativeFitness float64 = 0

	// Find the individual whose cumulative fitness exceeds the random value
	for i := 0; i < POPULATION_SIZE; i++ {
		cumulativeFitness += fitness(generation.population[i]) / generation.totalFitness
		if randValue < cumulativeFitness {
			return generation.population[i]
		}
	}

	// Fallback in case something goes wrong (shouldn't happen)
	return generation.population[POPULATION_SIZE-1]
}

func crossOver(parentX *Individual, parentY *Individual) *Individual {
	offspring := &Individual{cube: &Cube{
		sequence:  [SEQUENCE_SIZE]int{}, // default values of int are zeros
		successor: nil,
	}, parentX: parentX, parentY: parentY, mutation: false}

	// Random crossover point
	crossoverPoint1 := rand.Intn(SEQUENCE_SIZE / 2)
	crossoverPoint2 := rand.Intn(SEQUENCE_SIZE/2) + SEQUENCE_SIZE/2
	if crossoverPoint1 > crossoverPoint2 {
		crossoverPoint1, crossoverPoint2 = crossoverPoint2, crossoverPoint1
	}

	// Copy the middle part of the parent1 genes
	for i := crossoverPoint1; i < crossoverPoint2; i++ {
		offspring.cube.sequence[i] = parentX.cube.sequence[i]
	}

	// Fill in the remaining genes from parentY, ensuring uniqueness
	for i := 0; i < SEQUENCE_SIZE; i++ {
		if i < crossoverPoint1 || i >= crossoverPoint2 {
			searchIndex := i
			// Find a gene from parentY that is not already in offspring
			for contains(offspring.cube.sequence, parentY.cube.sequence[searchIndex]) {
				searchIndex++
				// Reset if searchIndex goes out of bounds
				if searchIndex >= SEQUENCE_SIZE {
					searchIndex = 0
				}
			}
			offspring.cube.sequence[i] = parentY.cube.sequence[searchIndex]
		}
	}
	offspring.cube.score = offspring.cube.ObjectiveFunction()
	return offspring
}

// Helper function to check if a value already exists in the genes
func contains(sequence [SEQUENCE_SIZE]int, value int) bool {
	for _, element := range sequence {
		if element == value {
			return true
		}
	}
	return false
}

func mutation(individual *Individual) *Individual {
	if rand.Float64() < MUTATION_RATE {
		individual.mutation = true
		individual.cube = individual.cube.RandomNeighbor()
	}
	return individual
}

func Evolution(generation *Generation) *Generation {
	// Initialize nextGeneration as a new Generation instance
	// GenerationDetail(generation)
	nextGeneration := &Generation{totalFitness: 0}
	var population [POPULATION_SIZE]*Individual
	for i := 0; i < POPULATION_SIZE; i++ {
		parentX := selection(generation)
		parentY := selection(generation)
		child := crossOver(parentX, parentY)
		mutation(child)
		var abortion int = 0
		for ((child.cube.score > parentX.cube.score) || (child.cube.score > parentY.cube.score)) && abortion < 1000 {
			if child.cube.score == parentX.cube.score {
				mutation(child)
			} else {
				child = crossOver(parentX, parentY)
			}
			abortion++
		}
		population[i] = child
		nextGeneration.totalFitness += fitness(child)
	}
	nextGeneration.SetPopulation(population)
	generation.SetNextGeneration(nextGeneration)
	// fmt.Println(generation.totalFitness)
	return nextGeneration
}

// BestIndividual finds and returns the individual with the highest fitness in the generation.
func BestIndividual(generation *Generation) *Individual {
	var best *Individual
	var bestFitness float64

	// Loop through each individual in the population
	for _, individual := range generation.population {
		// Calculate the fitness for the current individual
		individualFitness := fitness(individual)

		// Check if this individual has a higher fitness than the current best
		if best == nil || individualFitness > bestFitness {
			best = individual
			bestFitness = individualFitness
		}
	}

	return best
}

func SearchBestAVGScore(generation *Generation) {
	var total int = 0
	var best int = generation.population[0].cube.score
	for i := 0; i < POPULATION_SIZE; i++ {
		total += generation.population[i].cube.score
		if generation.population[i].cube.score < best {
			best = generation.population[i].cube.score
		}
	}
	generation.bestScore = best
	generation.avgScore = total / POPULATION_SIZE
}
