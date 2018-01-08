package euler

import (
	"testing"
	"math/rand"
)

func init() {
	rand.Seed(0)
}

type testQuery struct {
	a, b int
}

func benchmarkEulerRandom(b *testing.B, choiceLevelLink, numbers int) {
	tree := CreateEuler()
	queries := numbers / 2
	for j := 0; j < queries; j++ {
		tree.Link(rand.Intn(numbers), rand.Intn(numbers))
	}
	choiceLevelCut := choiceLevelLink * 2
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		query := testQuery{rand.Intn(numbers), rand.Intn(numbers)}

		chosenFunction := tree.IsConnected
		choice := rand.Intn(100)
		if choice < choiceLevelLink {
			chosenFunction = tree.Link
		} else if choice < choiceLevelCut {
			chosenFunction = tree.Cut
		}
		b.StartTimer()
		chosenFunction(query.a, query.b)
	}
}

//
// 33% IsConnected, 33% Link, 33% Cut
//

func BenchmarkEulerRandomThirdRead1000(b *testing.B) {
	benchmarkEulerRandom(b, 33, 1000)
}

func BenchmarkEulerRandomThirdRead10000(b *testing.B) {
	benchmarkEulerRandom(b, 33, 10000)
}

func BenchmarkEulerRandomThirdRead100000(b *testing.B) {
	benchmarkEulerRandom(b, 33, 100000)
}

func BenchmarkEulerRandomThirdRead1000000(b *testing.B) {
	benchmarkEulerRandom(b, 33, 1000000)
}

func BenchmarkEulerRandomThirdRead10000000(b *testing.B) {
	benchmarkEulerRandom(b, 33, 10000000)
}

//
// 50% IsConnected, 25% Link, 25% Cut
//

func BenchmarkEulerRandomHalfRead1000(b *testing.B) {
	benchmarkEulerRandom(b, 25, 1000)
}

func BenchmarkEulerRandomHalfRead10000(b *testing.B) {
	benchmarkEulerRandom(b, 25, 10000)
}

func BenchmarkEulerRandomHalfRead100000(b *testing.B) {
	benchmarkEulerRandom(b, 25, 100000)
}

func BenchmarkEulerRandomHalfRead1000000(b *testing.B) {
	benchmarkEulerRandom(b, 25, 1000000)
}

func BenchmarkEulerRandomHalfRead10000000(b *testing.B) {
	benchmarkEulerRandom(b, 25, 10000000)
}

//
// 100% IsConnected
//

func BenchmarkEulerRandomOnlyRead1000(b *testing.B) {
	benchmarkEulerRandom(b, 0, 1000)
}

func BenchmarkEulerRandomOnlyRead10000(b *testing.B) {
	benchmarkEulerRandom(b, 0, 10000)
}

func BenchmarkEulerRandomOnlyRead100000(b *testing.B) {
	benchmarkEulerRandom(b, 0, 100000)
}

func BenchmarkEulerRandomOnlyRead1000000(b *testing.B) {
	benchmarkEulerRandom(b, 0, 1000000)
}

func BenchmarkEulerRandomOnlyRead10000000(b *testing.B) {
	benchmarkEulerRandom(b, 0, 10000000)
}
