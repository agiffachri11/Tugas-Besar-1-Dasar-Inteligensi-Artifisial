package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

const (
	CUBE_ORDER     = 5                                    // Ukuran kubus (5x5x5)
	SEQUENCE_SIZE  = CUBE_ORDER * CUBE_ORDER * CUBE_ORDER // Total elemen dalam urutan (5^3 = 125)
	MAGIC_NUMBER   = (CUBE_ORDER * (SEQUENCE_SIZE + 1)) / 2 // Nilai 'magic number' yang diharapkan
	MAX_ITERATIONS = 1000                                 // Jumlah iterasi maksimum untuk Simulated Annealing
)

// Struktur Cube yang merepresentasikan sebuah kubus
type Cube struct {
	sequence  [SEQUENCE_SIZE]int // Array yang menyimpan urutan elemen dalam kubus
	score     int                // Skor yang mewakili "kualitas" kubus
}

// Fungsi untuk membuat Cube baru dengan urutan acak
func NewCube() *Cube {
	source := rand.NewSource(time.Now().UnixNano()) // Inisialisasi sumber angka acak
	r := rand.New(source)

	cube := &Cube{}
	for i := 0; i < SEQUENCE_SIZE; i++ {
		cube.sequence[i] = i + 1 // Mengisi urutan dari 1 hingga SEQUENCE_SIZE
	}
	r.Shuffle(SEQUENCE_SIZE, func(i, j int) {
		cube.sequence[i], cube.sequence[j] = cube.sequence[j], cube.sequence[i] // Mengacak urutan
	})

	cube.score = cube.ObjectiveFunction() // Menghitung skor awal kubus
	return cube
}

// ObjectiveFunction menghitung skor untuk Cube
func (c *Cube) ObjectiveFunction() int {
	// Fungsi sederhana untuk menghitung skor (placeholder)
	total := 0
	for _, val := range c.sequence {
		total += val // Menjumlahkan semua elemen dalam urutan sebagai skor
	}
	return total
}

// GenerateNeighbor membuat tetangga baru dengan menukar dua elemen dalam urutan
func (c *Cube) GenerateNeighbor() *Cube {
	neighbor := *c // Membuat salinan dari Cube saat ini
	i, j := rand.Intn(SEQUENCE_SIZE), rand.Intn(SEQUENCE_SIZE) // Memilih dua indeks acak
	neighbor.sequence[i], neighbor.sequence[j] = neighbor.sequence[j], neighbor.sequence[i] // Menukar elemen
	neighbor.score = neighbor.ObjectiveFunction() // Menghitung skor untuk tetangga baru
	return &neighbor
}

// Schedule menentukan suhu berdasarkan iterasi
func Schedule(t int) float64 {
	return float64(1000) / float64(t+1) // Jadwal pendinginan (cooling schedule)
}

// SimulatedAnnealing menjalankan algoritma simulated annealing
func SimulatedAnnealing(initial *Cube, maxIterations int) *Cube {
	current := initial // Mengatur Cube awal sebagai current
	for t := 1; t < maxIterations; t++ {
		T := Schedule(t) // Menentukan suhu berdasarkan iterasi saat ini
		if T <= 0 {
			return current // Menghentikan algoritma jika suhu mencapai 0
		}

		next := current.GenerateNeighbor() // Menghasilkan tetangga acak dari Cube saat ini
		deltaE := float64(next.score - current.score) // Menghitung perubahan skor (energi)

		if deltaE < 0 {
			current = next // Jika tetangga lebih baik, pindah ke tetangga tersebut
		} else {
			// Jika tetangga lebih buruk, terima dengan probabilitas tertentu
			probability := math.Exp(-deltaE / T)
			if rand.Float64() < probability {
				current = next // Terima tetangga lebih buruk dengan probabilitas e^(-deltaE/T)
			}
		}
	}
	return current // Mengembalikan Cube terbaik setelah semua iterasi selesai
}

func main() {
	initialCube := NewCube() // Membuat Cube awal dengan urutan acak
	fmt.Println("Initial Score:", initialCube.score) // Menampilkan skor awal
	result := SimulatedAnnealing(initialCube, MAX_ITERATIONS) // Menjalankan algoritma Simulated Annealing
	fmt.Println("Final Score:", result.score) // Menampilkan skor akhir setelah optimasi
}
