package cube

import "fmt"

// CountMagicSums prints the count of rows, columns, pillars, plane diagonals, and space diagonals with sum equal to MAGIC_NUMBER
func (c *Cube) CountMagicSums() {
	rows := c.CountMagicOnRow()
	columns := c.CountMagicOnColumn()
	pillars := c.CountMagicOnPillar()
	planeDiagonals := c.CountMagicOnPlaneDiagonal()
	spaceDiagonals := c.CountMagicOnSpaceDiagonal()

	// Print the results directly
	fmt.Printf("Rows: %d\n", rows)
	fmt.Printf("Columns: %d\n", columns)
	fmt.Printf("Pillars: %d\n", pillars)
	fmt.Printf("Plane Diagonals: %d\n", planeDiagonals)
	fmt.Printf("Space Diagonals: %d\n", spaceDiagonals)
}

func (c *Cube) CountMagicOnRow() int {
	var rows int = 0
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
	return rows
}

func (c *Cube) CountMagicOnColumn() int {
	var columns int = 0
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
	return columns
}

func (c *Cube) CountMagicOnPillar() int {
	var pillars int = 0
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
	return pillars
}

func (c *Cube) CountMagicOnPlaneDiagonal() int {
	var planeDiagonals int = 0
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
	return planeDiagonals
}

func (c *Cube) CountMagicOnSpaceDiagonal() int {
	var spaceDiagonals int = 0
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
	return spaceDiagonals
}

func GenerationDetail(generation *Generation) {
	fmt.Println("Generation Details:")
	fmt.Printf("Total Fitness: %.2f\n", generation.totalFitness)
	fmt.Printf("Population Size: %d\n", len(generation.population))
	for i := 0; i < POPULATION_SIZE; i++ {
		fmt.Printf("Individuals %d Score: %d\n", i+1, generation.population[i].cube.score)
	}
	fmt.Println()
}
