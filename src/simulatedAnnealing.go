package main

import (
	"fmt"
	"math"
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
	
	// Menghitung diagonal pada berbagai bidang
	return score - rows - columns - pillars - planeDiagonals - spaceDiagonals
}

// Fungsi untuk membuat neighbor baru dengan menukar dua angka secara acak
func (c *Cube) GenerateNeighbor() *Cube {
	neighbor := *c // Menyalin state cube saat ini
	rand.Seed(time.Now().UnixNano())
	i, j := rand.Intn(SEQUENCE_SIZE), rand.Intn(SEQUENCE_SIZE)
	neighbor.sequence[i], neighbor.sequence[j] = neighbor.sequence[j], neighbor.sequence[i]
	neighbor.score = neighbor.ObjectiveFuntion()
	return &neighbor
}

// Fungsi untuk mengatur penurunan suhu dalam setiap iterasi 
func Schedule(t int) float64 {
	return 100.0 / float64(t+1)
}

// Fungsi MakeNode untuk membuat node dari intial state
func MakeNode(initialCube *Cube) *Cube {
	return initialCube
}

// Algoritma Simulated Annealing untuk mencari solusi 
func SimulatedAnnealing(problem *Cube, schedule func(int) float64) *Cube {
	current := MakeNode(problem) // inisialisasi current sebagai kubus awal
	rand.Seed(time.Now().UnixNano())

	for t := 1; ; t++ {
		temperature := schedule(t) // Menghitung suhu sesuai iterasi
		// Jika T sudah mencapai 0, proses pencarian berhenti
		if temperature == 0 {
			return current
		}

		// Membuat konfigurasi baru
		next := current.GenerateNeighbor() 
		// Menghitung perubahan skor atau deltaE antara current dan neighbor
		deltaE := float64(next.score - current.score)
		
		// Memperbarui current berdasarkan deltaE dan probabilitas
		// Jika deltaE > 0, maka neighbor adalah solusi yang lebih baik, next di assign ke current
		if deltaE > 0 {
			current = next
		} else if math.Exp(deltaE/temperature) > rand.Float64() {
			current = next
		}
	}
}

func main() {
	// inisialiasi cube awal 
	initialCube := NewCube()

	// Menjalankan algoritma Simulated Annealing
	solution := SimulatedAnnealing(initialCube, Schedule)

	// Menampilkan solusi akhir
	fmt.Println("State akhir:", solution.sequence)
	fmt.Println("Nilai objective akhir:", solution.score)
}