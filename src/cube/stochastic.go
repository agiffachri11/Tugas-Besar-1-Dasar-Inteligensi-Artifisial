package main

import (
	"fmt"
	"math/rand"
	"time"
)

const SEQUENCE_SIZE = 10

type Cube struct {
	sequence []int
	score    int
}

func NewCube() *Cube {
    cube := &Cube{}
    rand.Seed(time.Now().UnixNano())

    for i := 0; i < SEQUENCE_SIZE; i++ {
        cube.sequence = append(cube.sequence, i+1)
    }

    // Acak urutan
    rand.Shuffle(len(cube.sequence), func(i, j int) {
        cube.sequence[i], cube.sequence[j] = cube.sequence[j], cube.sequence[i]
    })

    // Hitung skor awal
    cube.score = cube.ObjectiveFunction()

    return cube
}

func (c *Cube) ObjectiveFunction() int {
	total := 0
	for _, val := range c.sequence {
		total += val
	}
	return total
}

func (c *Cube) GenerateNeighbor() *Cube {
    neighbor := *c 
    rand.Seed(time.Now().UnixNano())
    i, j := rand.Intn(SEQUENCE_SIZE), rand.Intn(SEQUENCE_SIZE)
    fmt.Printf("Swapping indices: %d and %d\n", i, j) // Debugging statement
    neighbor.sequence[i], neighbor.sequence[j] = neighbor.sequence[j], neighbor.sequence[i]
    neighbor.score = neighbor.ObjectiveFunction()
    fmt.Println("Neighbor sequence:", neighbor.sequence) // Debugging statement
    return &neighbor
}

func StochasticHillClimbing(c *Cube, maxIteration int) *Cube {
    current := c
    var neighbor *Cube
    for i := 0; i < maxIteration; i++ {
        fmt.Printf("Iteration %d: Current Score: %d\n", i, current.score)
        neighbor = c.GenerateNeighbor()
        fmt.Println("Neighbor Score:", neighbor.score)

        if neighbor.score < current.score {
            current = neighbor
            fmt.Println("Better neighbor found, updating current cube")
        }
    }
    return current
}

func main() {
	initialCube := NewCube() 
	maxIterations := 10
	result := StochasticHillClimbing(initialCube, maxIterations)

	fmt.Println("Final score:", result.score)
}

