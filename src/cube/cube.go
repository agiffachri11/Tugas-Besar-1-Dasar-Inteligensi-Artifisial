package cube

import (
	"math/rand"
	"time"
)

// Cube struct definition
type Cube struct {
	sequence  [SEQUENCE_SIZE]int
	score     int
	successor *Cube
}

func (c *Cube) GetSequence() [SEQUENCE_SIZE]int {
	return c.sequence
}

func (c *Cube) GetScore() int {
	return c.score
}

func (c *Cube) GetSuccessor() *Cube {
	return c.successor
}

func (c *Cube) SetSuccessor(successor *Cube) {
	c.successor = successor
}

// NewCube constructor initializes a Cube with a random sequence and score
func NewCube() *Cube {
	// Create a random source and random generator
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	// Create a new Cube instance with initialized score and nil successor
	cube := &Cube{successor: nil}

	// Populate and shuffle the sequence array in place
	for i := 0; i < SEQUENCE_SIZE; i++ {
		cube.sequence[i] = i + 1
	}

	r.Shuffle(SEQUENCE_SIZE, func(i, j int) {
		cube.sequence[i], cube.sequence[j] = cube.sequence[j], cube.sequence[i]
	})

	cube.score = cube.ObjectiveFunction()

	return cube
}

// Copy cube without copying the successor (successor = NIL)
func CopyCube(c *Cube) *Cube {
	// Create a new Cube instance with initialized score and nil successor
	cube := &Cube{successor: nil}

	// Copy the sequence array in place
	for i := 0; i < SEQUENCE_SIZE; i++ {
		cube.sequence[i] = c.sequence[i]
	}

	cube.score = c.score

	return cube
}

// Helper function to get an element at x, y, z
func (c *Cube) get(x, y, z int) int {
	return c.sequence[x*CUBE_ORDER*CUBE_ORDER+y*CUBE_ORDER+z]
}

func (c *Cube) RandomNeighbor() *Cube {
	neighbor := CopyCube(c)
	var idx1 int = rand.Intn(SEQUENCE_SIZE)
	var idx2 int = rand.Intn(SEQUENCE_SIZE)
	neighbor.sequence[idx1] = c.sequence[idx2]
	neighbor.sequence[idx2] = c.sequence[idx1]
	neighbor.score = neighbor.ObjectiveFunction()
	return neighbor
}
