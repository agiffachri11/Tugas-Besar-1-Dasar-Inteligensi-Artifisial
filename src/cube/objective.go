package cube

func absoluteInt(n int) int {
	if n < 0 {
		return n * -1
	}
	return n
}

// IMPLEMENTATION #1
// Find a cube value, 0 if it's a diagonal magic cube
// func (c *Cube) ObjectiveFunction() int {
// 	score := MAGIC_NUMBER_AMOUNT
// 	rows := c.CountMagicOnRow()
// 	columns := c.CountMagicOnColumn()
// 	pillars := c.CountMagicOnPillar()
// 	planeDiagonals := c.CountMagicOnPlaneDiagonal()
// 	spaceDiagonals := c.CountMagicOnSpaceDiagonal()

// 	// Calculate the score by subtracting the counts
// 	return score - (rows + columns + pillars + planeDiagonals + spaceDiagonals)
// }

// IMPLEMENTATION #2
// Find a cube value, 0 if it's a diagonal magic cube
func (c *Cube) ObjectiveFunction() int {
	score := 0
	// Count rows in each XY layer (constant Z)
	for z := 0; z < CUBE_ORDER; z++ {
		for y := 0; y < CUBE_ORDER; y++ {
			sum := 0
			for x := 0; x < CUBE_ORDER; x++ {
				sum += c.get(x, y, z)
			}
			score += absoluteInt(MAGIC_NUMBER - sum)
		}
	}
	// Count columns in each XY layer (constant Z)
	for z := 0; z < CUBE_ORDER; z++ {
		for x := 0; x < CUBE_ORDER; x++ {
			sum := 0
			for y := 0; y < CUBE_ORDER; y++ {
				sum += c.get(x, y, z)
			}
			score += absoluteInt(MAGIC_NUMBER - sum)
		}
	}
	// Count pillars (constant X and Y)
	for x := 0; x < CUBE_ORDER; x++ {
		for y := 0; y < CUBE_ORDER; y++ {
			sum := 0
			for z := 0; z < CUBE_ORDER; z++ {
				sum += c.get(x, y, z)
			}
			score += absoluteInt(MAGIC_NUMBER - sum)
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
		score += absoluteInt(MAGIC_NUMBER - sum1)
		score += absoluteInt(MAGIC_NUMBER - sum2)
	}
	for x := 0; x < CUBE_ORDER; x++ {
		// YZ plane diagonals (constant X)
		sum1, sum2 := 0, 0
		for i := 0; i < CUBE_ORDER; i++ {
			sum1 += c.get(x, i, i)
			sum2 += c.get(x, i, CUBE_ORDER-1-i)
		}
		score += absoluteInt(MAGIC_NUMBER - sum1)
		score += absoluteInt(MAGIC_NUMBER - sum2)
	}
	for y := 0; y < CUBE_ORDER; y++ {
		// XZ plane diagonals (constant Y)
		sum1, sum2 := 0, 0
		for i := 0; i < CUBE_ORDER; i++ {
			sum1 += c.get(i, y, i)
			sum2 += c.get(CUBE_ORDER-1-i, y, i)
		}
		score += absoluteInt(MAGIC_NUMBER - sum1)
		score += absoluteInt(MAGIC_NUMBER - sum2)
	}
	// Count space diagonals (corner to corner)
	sum1, sum2, sum3, sum4 := 0, 0, 0, 0
	for i := 0; i < CUBE_ORDER; i++ {
		sum1 += c.get(i, i, i)
		sum2 += c.get(i, i, CUBE_ORDER-1-i)
		sum3 += c.get(i, CUBE_ORDER-1-i, i)
		sum4 += c.get(CUBE_ORDER-1-i, i, i)
	}
	score += absoluteInt(MAGIC_NUMBER - sum1)
	score += absoluteInt(MAGIC_NUMBER - sum2)
	score += absoluteInt(MAGIC_NUMBER - sum3)
	score += absoluteInt(MAGIC_NUMBER - sum4)
	return score
}
