package localsearch

import (
	"diagonalmagiccube/cube"
)

func StochasticHillClimbing(c *cube.Cube, maxIteration int) *cube.Cube {
	current := c
	var neighbor *cube.Cube
	for i := 0; i < maxIteration; i++ {
		neighbor = current.RandomNeighbor()
		if neighbor.GetScore() < current.GetScore() {
			current.SetSuccessor(neighbor)
			current = neighbor
		}
	}
	return current
}