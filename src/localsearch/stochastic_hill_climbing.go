package localsearch

import (
	"diagonalmagiccube/cube"
)

func StochasticHillClimbing(c *cube.Cube, maxIteration int, stuckCount *int, sliceDeltaE []float64) *cube.Cube {
	current := c
	var neighbor *cube.Cube
	for i := 0; i < maxIteration; i++ {
		neighbor = current.RandomNeighbor()
		current.SetSuccessor(cube.CopyCube(current))
		if neighbor.GetScore() < current.GetScore() {
			current.SetSuccessor(neighbor)
		}
		current = current.GetSuccessor()
	}
	return current
}
