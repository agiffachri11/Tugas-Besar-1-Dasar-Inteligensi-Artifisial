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

		// Membuat konfigurasi baru
		next := current.RandomNeighbor()
		// Menghitung perubahan skor atau deltaE antara current dan neighbor
		deltaE := float64(current.GetScore() - next.GetScore())

		// Memperbarui current berdasarkan deltaE dan probabilitas
		// Jika deltaE > 0, maka neighbor adalah solusi yang lebih baik, next di assign ke current
		if deltaE > 0 {
			current.SetSuccessor(next)
			current = next
		} else if math.Exp(deltaE/temperature) > rand.Float64() {
			current.SetSuccessor(next)
			current = next
		}
	}
	return current
}

// SimulatedAnnealing menjalankan algoritma simulated annealing
// func SimulatedAnnealing(initial *cube.Cube, maxIterations int) *cube.Cube {
// 	current := initial // Mengatur Cube awal sebagai current
// 	for t := 1; t < maxIterations; t++ {
// 		T := Schedule(t) // Menentukan suhu berdasarkan iterasi saat ini
// 		if T <= 0 {
// 			return current // Menghentikan algoritma jika suhu mencapai 0
// 		}

// 		next := current.GenerateNeighbor()            // Menghasilkan tetangga acak dari Cube saat ini
// 		deltaE := float64(next.score - current.score) // Menghitung perubahan skor (energi)

// 		if deltaE < 0 {
// 			current = next // Jika tetangga lebih baik, pindah ke tetangga tersebut
// 		} else {
// 			// Jika tetangga lebih buruk, terima dengan probabilitas tertentu
// 			probability := math.Exp(-deltaE / T)
// 			if rand.Float64() < probability {
// 				current = next // Terima tetangga lebih buruk dengan probabilitas e^(-deltaE/T)
// 			}
// 		}
// 	}
// 	return current // Mengembalikan Cube terbaik setelah semua iterasi selesai
// }
