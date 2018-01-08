package euler

import (
	"testing"
	"math/rand"
)

func init() {
	rand.Seed(0)
}

type query struct {
	a, b int
}

func benchmarkEulerRandom(b *testing.B, numbers int) {
	tree := CreateEuler()
	queries := numbers / 2
	for j := 0; j < queries; j++ {
		tree.Link(rand.Intn(numbers), rand.Intn(numbers))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		q := query{rand.Intn(numbers), rand.Intn(numbers)}

		f := tree.IsConnected
		choice := rand.Intn(100)
		if choice < 33 {
			f = tree.Link
		} else if choice < 66 {
			f = tree.Cut
		}
		b.StartTimer()
		f(q.a, q.b)
	}
}

func BenchmarkEulerRandom1000(b *testing.B) {
	benchmarkEulerRandom(b, 1000)
}

func BenchmarkEulerRandom10000(b *testing.B) {
	benchmarkEulerRandom(b, 10000)
}

func BenchmarkEulerRandom100000(b *testing.B) {
	benchmarkEulerRandom(b, 100000)
}

func BenchmarkEulerRandom1000000(b *testing.B) {
	benchmarkEulerRandom(b, 1000000)
}

func BenchmarkEulerRandom10000000(b *testing.B) {
	benchmarkEulerRandom(b, 10000000)
}
