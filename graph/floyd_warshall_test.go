package graph

import (
	"math"
	"math/rand"
	"reflect"
	"testing"
)

func TestFloydWarshall(t *testing.T) {
	cases := []struct {
		n     int
		edges [][]int
		want  [][]int
	}{
		{0, nil, [][]int{}},
		{3, [][]int{{0, 1, 5}}, [][]int{{0, 5, math.MaxInt}, {math.MaxInt, 0, math.MaxInt}, {math.MaxInt, math.MaxInt, 0}}},
		{3, [][]int{{0, 1, 2}, {0, 1, 7}, {1, 2, -3}, {0, 0, 9}}, [][]int{{0, 2, -1}, {math.MaxInt, 0, -3}, {math.MaxInt, math.MaxInt, 0}}},
		{2, [][]int{{0, 1, -2}, {1, 0, 1}}, nil},
		{1, [][]int{{0, 0, -1}}, nil},
		{3, [][]int{{1, 2, -2}, {2, 1, 1}}, nil},
	}
	for _, tc := range cases {
		if got := floydWarshall(tc.n, tc.edges); !reflect.DeepEqual(got, tc.want) {
			t.Fatalf("edges=%v got %v want %v", tc.edges, got, tc.want)
		}
	}
}

func TestFloydAgainstBellmanFord(t *testing.T) {
	rng := rand.New(rand.NewSource(9))
	for trial := 0; trial < 200; trial++ {
		n := rng.Intn(10) + 1
		var edges [][]int
		// Potential differences allow negative edges without negative cycles.
		potential := make([]int, n)
		for i := range potential {
			potential[i] = rng.Intn(21) - 10
		}
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if rng.Intn(4) == 0 {
					edges = append(edges, []int{i, j, rng.Intn(10) + potential[j] - potential[i]})
				}
			}
		}
		got := floydWarshall(n, edges)
		if got == nil {
			t.Fatal("unexpected negative cycle")
		}
		for src := 0; src < n; src++ {
			want := BellmanFord(n, edges, src)
			if !reflect.DeepEqual(got[src], want) {
				t.Fatalf("src=%d edges=%v got=%v want=%v", src, edges, got[src], want)
			}
		}
	}
}

func TestFloydOverflow(t *testing.T) {
	for _, edges := range [][][]int{
		{{0, 1, math.MaxInt}},
		{{0, 1, math.MaxInt - 1}, {1, 2, 2}},
		{{0, 1, math.MinInt}, {1, 2, -1}},
	} {
		func() {
			defer func() {
				if recover() == nil {
					t.Error("expected overflow panic", edges)
				}
			}()
			floydWarshall(3, edges)
		}()
	}
}
