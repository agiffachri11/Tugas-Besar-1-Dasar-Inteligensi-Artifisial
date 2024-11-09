package localsearch

import (
	"diagonalmagiccube/cube"
	"math"
	"math/rand"
)

func schedule(t int) float64 {
	return float64(100) / float64(t+1) // Jadwal pendinginan (cooling schedule)
}

// Algoritma Simulated Annealing untuk mencari solusi
func SimulatedAnnealing(c *cube.Cube, maxIterations int) *cube.Cube {
	current := c // inisialisasi current sebagai kubus awal

	for t := 0; t < maxIterations; t++ {
		temperature := schedule(t) // Menghitung suhu sesuai iterasi

		// Membuat konfigurasi neighbor (tetangga) baru dari state sekarang
		next := current.RandomNeighbor()
		// Menghitung perubahan skor atau deltaE antara current dan neighbor
		deltaE := float64(next.GetScore() - current.GetScore())

		// Memperbarui current berdasarkan deltaE dan probabilita
		if deltaE < 0 { // Jika next lebih baik, (deltaE negatif) maka pindah ke next
			current.SetSuccessor(next)
			current = next
		} else if math.Exp(-deltaE/temperature) > rand.Float64() { // Jika lebih buruk, terima dengan probabilitas tertentu
			current.SetSuccessor(next)
			current = next
		}
	}
	return current
}
