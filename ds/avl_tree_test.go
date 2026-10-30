package ds

import (
	"math"
	"math/rand"
	"slices"
	"sort"
	"testing"
)

func TestAVLTreeBalance(t *testing.T) {
	tree := NewAVLTree()
	values := map[int]int{}
	rng := rand.New(rand.NewSource(42))
	var check func(*avlNode) (int, int)
	check = func(n *avlNode) (int, int) {
		if n == nil {
			return 0, 0
		}
		lh, lc := check(n.left)
		rh, rc := check(n.right)
		height := max(lh, rh) + 1
		if int(n.hght) != height || lh-rh > 1 || rh-lh > 1 {
			t.Fatalf("node %d: stored height %d, actual %d, child heights %d/%d", n.Value, n.hght, height, lh, rh)
		}
		if values[n.Value] != n.count || n.count <= 0 {
			t.Fatalf("unexpected value %d", n.Value)
		}
		if n.left != nil && n.left.Value >= n.Value || n.right != nil && n.right.Value <= n.Value {
			t.Fatal("invalid node order")
		}
		return height, lc + rc + n.count
	}
	for i := 0; i < 10000; i++ {
		v := rng.Intn(100)
		if rng.Intn(2) == 0 {
			tree.Insert(v)
			values[v]++
		} else {
			tree.Delete(v)
			if values[v] > 1 {
				values[v]--
			} else {
				delete(values, v)
			}
		}
		want := []int{}
		for value, count := range values {
			for j := 0; j < count; j++ {
				want = append(want, value)
			}
		}
		sort.Ints(want)
		if !slices.Equal(tree.Values(), want) {
			t.Fatal("incorrect sorted traversal")
		}
		query := rng.Intn(104) - 2
		pos := sort.SearchInts(want, query)
		lower, lowerOK := tree.Lower(query)
		if lowerOK != (pos > 0) || lowerOK && lower != want[pos-1] {
			t.Fatal("incorrect lower neighbor")
		}
		pos = sort.Search(len(want), func(i int) bool { return want[i] > query })
		higher, higherOK := tree.Higher(query)
		if higherOK != (pos < len(want)) || higherOK && higher != want[pos] {
			t.Fatal("incorrect higher neighbor")
		}
		if tree.Count(query) != values[query] {
			t.Fatal("incorrect multiplicity")
		}
		_, count := check(tree.root)
		wantCount := 0
		for _, c := range values {
			wantCount += c
		}
		if count != wantCount || tree.Len() != wantCount {
			t.Fatal("incorrect tree size")
		}
		for v := range values {
			if _, ok := tree.Get(v); !ok {
				t.Fatalf("missing value %d", v)
			}
		}
	}
}

func TestAVLTreeDuplicateSuccessor(t *testing.T) {
	tree := NewAVLTree().Insert(2).Insert(1).Insert(3).Insert(3).Insert(3)
	tree.Delete(2)
	if !slices.Equal(tree.Values(), []int{1, 3, 3, 3}) || tree.Len() != 4 || tree.Count(3) != 3 {
		t.Fatal("lost or duplicated successor values")
	}
	for _, want := range []int{3, 3, 3, 1} {
		got, ok := tree.Pop()
		if !ok || got != want {
			t.Fatalf("Pop()=%d,%v, want %d", got, ok, want)
		}
	}
	if _, ok := tree.Pop(); ok {
		t.Fatal("pop from empty tree succeeded")
	}
}

func TestAVLTreeTraversalAndBounds(t *testing.T) {
	var tree AVLTree
	tree.Insert(math.MinInt).Insert(5).Insert(5).Insert(math.MaxInt)
	if _, ok := tree.Lower(math.MinInt); ok {
		t.Fatal("unexpected lower neighbor")
	}
	if _, ok := tree.Higher(math.MaxInt); ok {
		t.Fatal("unexpected higher neighbor")
	}
	visited := []int{}
	tree.Ascend(func(v int) bool { visited = append(visited, v); return len(visited) < 2 })
	if !slices.Equal(visited, []int{math.MinInt, 5}) {
		t.Fatal("early stop failed")
	}
	copy := tree.Values()
	copy[0] = 42
	if v, _ := tree.Head(); v != math.MinInt {
		t.Fatal("Values aliases tree")
	}
	for _, want := range []int{math.MinInt, 5, 5, math.MaxInt} {
		got, ok := tree.Pull()
		if !ok || got != want {
			t.Fatalf("Pull()=%d,%v, want %d", got, ok, want)
		}
	}
	if _, ok := tree.Pull(); ok {
		t.Fatal("pull from empty tree succeeded")
	}
	tree.Delete(5)
	if tree.Len() != 0 {
		t.Fatal("delete missing value changed size")
	}
	if _, ok := tree.Get(5); ok {
		t.Fatal("found deleted value")
	}
	if _, ok := tree.Head(); ok {
		t.Fatal("empty tree has head")
	}
	if _, ok := tree.Tail(); ok {
		t.Fatal("empty tree has tail")
	}
	if len(tree.Values()) != 0 {
		t.Fatal("empty traversal")
	}
}
