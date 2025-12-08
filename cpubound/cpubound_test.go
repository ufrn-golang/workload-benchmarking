package cpubound

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func generateRandomNumbers(n int) []int {
	numbers := make([]int, n)
	seed := rand.NewSource(time.Now().UnixNano())
	rand := rand.New(seed)

	for i := 0; i < n; i++ {
		numbers[i] = rand.Intn(n)
	}
	return numbers
}

// Simple unit test functions
func TestMergeSortSequential(t *testing.T) {
	unsorted := []int{10, 6, 2, 1, 5, 8, 3, 4, 7, 9}
	sorted := mergeSortSequential(unsorted)
	fmt.Println(sorted)
}

func TestMergeSortConcurrent(t *testing.T) {
	unsorted := []int{10, 6, 2, 1, 5, 8, 3, 4, 7, 9}
	sorted := mergeSortConcurrent(unsorted)
	fmt.Println(sorted)
	fmt.Printf("Goroutines created: %d\n", goroutines)
}

// Benchmarking functions
func BenchmarkMergeSortSequential(b *testing.B) {
	numbers := generateRandomNumbers(1 << 20)
	b.StartTimer()
	for b.Loop() {
		mergeSortSequential(numbers)
	}
	b.StopTimer()
	b.ReportMetric(b.Elapsed().Seconds()/float64(b.N), "s/op")
}

func BenchmarkMergeSortConcurrent(b *testing.B) {
	numbers := generateRandomNumbers(1 << 20)
	b.StartTimer()
	for b.Loop() {
		mergeSortConcurrent(numbers)
	}
	b.StopTimer()
	b.ReportMetric(b.Elapsed().Seconds()/float64(b.N), "s/op")
}
