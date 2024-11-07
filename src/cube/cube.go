package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	CUBE_ORDER          = 5
	MAGIC_NUMBER        = (CUBE_ORDER * (SEQUENCE_SIZE + 1)) / 2       // 315
	MAGIC_NUMBER_AMOUNT = 3*(CUBE_ORDER*CUBE_ORDER) + 6*CUBE_ORDER + 4 // 109
	SEQUENCE_SIZE       = CUBE_ORDER * CUBE_ORDER * CUBE_ORDER         // 125
)

// Cube struct definition
type Cube struct {
	sequence  [SEQUENCE_SIZE]int
	score     int
	successor *Cube
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

	cube.score = cube.ObjectiveFuntion()

	return cube
}

func CopyCube(c *Cube) *Cube {
	// Create a new Cube instance with initialized score and nil successor
	cube := &Cube{successor: nil}

	// Copy the sequence array in place
	for i := 0; i < SEQUENCE_SIZE; i++ {
		cube.sequence[i] = c.sequence[i]
	}

	cube.score = cube.ObjectiveFuntion()

	return cube
}

// Helper function to get an element at x, y, z
func (c *Cube) get(x, y, z int) int {
	return c.sequence[x*CUBE_ORDER*CUBE_ORDER+y*CUBE_ORDER+z]
}

// Find a cube value, 0 if it's a diagonal magic cube
func (c *Cube) ObjectiveFuntion() int {
	score := MAGIC_NUMBER_AMOUNT
	rows, columns, pillars, planeDiagonals, spaceDiagonals := 0, 0, 0, 0, 0
	// Count rows in each XY layer (constant Z)
	for z := 0; z < CUBE_ORDER; z++ {
		for y := 0; y < CUBE_ORDER; y++ {
			sum := 0
			for x := 0; x < CUBE_ORDER; x++ {
				sum += c.get(x, y, z)
			}
			if sum == MAGIC_NUMBER {
				rows++
			}
		}
	}
	// Count columns in each XY layer (constant Z)
	for z := 0; z < CUBE_ORDER; z++ {
		for x := 0; x < CUBE_ORDER; x++ {
			sum := 0
			for y := 0; y < CUBE_ORDER; y++ {
				sum += c.get(x, y, z)
			}
			if sum == MAGIC_NUMBER {
				columns++
			}
		}
	}
	// Count pillars (constant X and Y)
	for x := 0; x < CUBE_ORDER; x++ {
		for y := 0; y < CUBE_ORDER; y++ {
			sum := 0
			for z := 0; z < CUBE_ORDER; z++ {
				sum += c.get(x, y, z)
			}
			if sum == MAGIC_NUMBER {
				pillars++
			}
		}
	}
	// Count plane diagonals (each XY, YZ, XZ plane)
	for z := 0; z < CUBE_ORDER; z++ {
		// XY plane diagonals (constant Z)
		sum1, sum2 := 0, 0
		for i := 0; i < CUBE_ORDER; i++ {
			sum1 += c.get(i, i, z)
			sum2 += c.get(i, CUBE_ORDER-1-i, z)
		}
		if sum1 == MAGIC_NUMBER {
			planeDiagonals++
		}
		if sum2 == MAGIC_NUMBER {
			planeDiagonals++
		}
	}
	for x := 0; x < CUBE_ORDER; x++ {
		// YZ plane diagonals (constant X)
		sum1, sum2 := 0, 0
		for i := 0; i < CUBE_ORDER; i++ {
			sum1 += c.get(x, i, i)
			sum2 += c.get(x, i, CUBE_ORDER-1-i)
		}
		if sum1 == MAGIC_NUMBER {
			planeDiagonals++
		}
		if sum2 == MAGIC_NUMBER {
			planeDiagonals++
		}
	}
	for y := 0; y < CUBE_ORDER; y++ {
		// XZ plane diagonals (constant Y)
		sum1, sum2 := 0, 0
		for i := 0; i < CUBE_ORDER; i++ {
			sum1 += c.get(i, y, i)
			sum2 += c.get(CUBE_ORDER-1-i, y, i)
		}
		if sum1 == MAGIC_NUMBER {
			planeDiagonals++
		}
		if sum2 == MAGIC_NUMBER {
			planeDiagonals++
		}
	}
	// Count space diagonals (corner to corner)
	sum1, sum2, sum3, sum4 := 0, 0, 0, 0
	for i := 0; i < CUBE_ORDER; i++ {
		sum1 += c.get(i, i, i)
		sum2 += c.get(i, i, CUBE_ORDER-1-i)
		sum3 += c.get(i, CUBE_ORDER-1-i, i)
		sum4 += c.get(CUBE_ORDER-1-i, i, i)
	}
	if sum1 == MAGIC_NUMBER {
		spaceDiagonals++
	}
	if sum2 == MAGIC_NUMBER {
		spaceDiagonals++
	}
	if sum3 == MAGIC_NUMBER {
		spaceDiagonals++
	}
	if sum4 == MAGIC_NUMBER {
		spaceDiagonals++
	}
	return score - rows - columns - pillars - planeDiagonals - spaceDiagonals
}

func main() {
	// Create a new Cube
	cube := NewCube()
	// Display Cube details
	fmt.Println("Sequence:", cube.sequence)
	fmt.Println("Score:", cube.score)
	fmt.Println("Successor:", cube.successor)
	cube.CountMagicSums()
}
