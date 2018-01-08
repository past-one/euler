package euler

import (
	"testing"
	"reflect"
)

// for testing, not benchmarking
func createTestTree(allIndices [][]int) *Euler {
	tree := CreateEuler()

	for _, indices := range allIndices {
		treaps := make([]*Treap, len(indices))
		var treap *Treap
		for i, index := range indices {
			treaps[i] = &Treap{priority: index, size: 1, vertex: index}
			tree.treaps[index] = treaps[i]
			treap = Merge(treap, treaps[i])
		}
	}

	return tree
}

func TestEuler_String(t *testing.T) {
	tests := []struct {
		tree     *Euler
		expected string
	}{
		{
			createTestTree([][]int{
				{1, 2, 1},
			}),
			"1-2-1",
		},
		{
			createTestTree([][]int{
				{1},
				{2},
			}),
			"1\n2",
		},
		{
			createTestTree([][]int{
				{1, 2, 1},
				{3, 4, 3},
			}),
			"1-2-1\n3-4-3",
		},
	}

	for _, test := range tests {
		got := test.tree.String()

		if got != test.expected {
			t.Errorf("Expected \n'%v'\nGot \n'%v'\n", test.expected, got)
		}
	}
}

func TestEuler_IsConnected(t *testing.T) {
	tests := []struct {
		tree          *Euler
		first, second Vertex
		expected      bool
	}{
		{
			createTestTree([][]int{
				{1, 2, 1},
			}),
			1,
			2,
			true,
		},
		{
			createTestTree([][]int{
				{1},
				{2},
			}),
			1,
			2,
			false,
		},
		{
			createTestTree([][]int{
				{1, 2, 1},
				{3, 4, 3},
			}),
			2,
			4,
			false,
		},
		{
			createTestTree([][]int{
				{1, 2, 1},
				{3, 4, 3},
			}),
			3,
			4,
			true,
		},
	}

	for _, test := range tests {
		got := test.tree.IsConnected(test.first, test.second)

		if got != test.expected {
			t.Errorf("In test %v\nexpected %v,\ngot %v", test, test.expected, got)
		}
	}
}

func TestEuler_Link(t *testing.T) {
	tests := []struct {
		tree           *Euler
		first, second  Vertex
		expectedTree   []string
		expectedResult bool
	}{
		{
			createTestTree([][]int{
				{1, 2, 1},
			}),
			1,
			2,
			[]string{
				"1-2-1",
			},
			false,
		},
		{
			createTestTree([][]int{
				{1},
				{2},
			}),
			1,
			2,
			[]string{
				"1-2-1",
			},
			true,
		},
		{
			createTestTree([][]int{
				{1, 3, 1},
				{2},
			}),
			1,
			2,
			[]string{
				"1-3-1-2-1",
			},
			true,
		},
		{
			createTestTree([][]int{
				{2},
				{1, 3, 1},
			}),
			2,
			1,
			[]string{
				"2-1-3-1-2",
			},
			true,
		},
		{
			createTestTree([][]int{
				{1, 2, 1},
				{3, 4, 3, 5, 3},
			}),
			1,
			4,
			[]string{
				"1-2-1-4-3-5-3-4-1",
			},
			true,
		},
		{
			createTestTree([][]int{
				{3, 4, 3, 5, 3},
				{1, 2, 1},
			}),
			4,
			1,
			[]string{
				"3-4-1-2-1-4-3-5-3",
			},
			true,
		},
	}

	for _, test := range tests {
		was := test.tree.Strings()
		gotResult := test.tree.Link(test.first, test.second)
		gotTree := test.tree.Strings()

		if gotResult != test.expectedResult ||
			!reflect.DeepEqual(gotTree, test.expectedTree) {
			t.Errorf(
				"%v.Link(%v, %v)\nExpected %v,\n'%v'\nGot %v,\n'%v'",
				was,
				test.first,
				test.second,
				test.expectedResult,
				test.expectedTree,
				gotResult,
				gotTree,
			)
		}
	}
}

func createTestTreeByLink(edges []struct{ a, b int }, vertices []int) *Euler {
	tree := CreateEuler()

	for _, edge := range edges {
		if !tree.Link(edge.a, edge.b) {
			panic("invalid arguments in tests")
		}
	}

	for _, vertex := range vertices {
		// init vertex
		tree.getTreap(vertex)
	}

	return tree
}

func TestEuler_Cut(t *testing.T) {
	tests := []struct {
		tree           *Euler
		first, second  Vertex
		expectedTree   []string
		expectedResult bool
	}{
		{
			createTestTreeByLink(
				[]struct{ a, b int }{},
				[]int{1, 2},
			),
			1,
			2,
			[]string{"1", "2"},
			false,
		},
		{
			createTestTreeByLink(
				[]struct{ a, b int }{
					{1, 2},
				},
				[]int{},
			),
			1,
			2,
			[]string{"1", "2"},
			true,
		},
		{
			createTestTreeByLink(
				[]struct{ a, b int }{
					{1, 2},
					{3, 4},
					{2, 3},
				},
				[]int{},
			),
			2,
			3,
			[]string{"1-2-1", "3-4-3"},
			true,
		},
		{
			createTestTreeByLink(
				[]struct{ a, b int }{
					{1, 2},
					{3, 4},
					{2, 3},
				},
				[]int{},
			),
			1,
			2,
			[]string{"1", "2-3-4-3-2"},
			true,
		},
		{
			createTestTreeByLink(
				[]struct{ a, b int }{
					{1, 2},
					{3, 4},
					{2, 3},
					{5, 2},
				},
				[]int{},
			),
			1,
			2,
			[]string{"1", "5-2-3-4-3-2-5"},
			true,
		},
		{
			createTestTreeByLink(
				[]struct{ a, b int }{
					{1, 2},
					{3, 4},
					{2, 3},
					{5, 3},
				},
				[]int{},
			),
			1,
			2,
			[]string{"1", "5-3-2-3-4-3-5"},
			true,
		},
	}

	for _, test := range tests {
		was := test.tree.Strings()
		gotResult := test.tree.Cut(test.first, test.second)
		gotTree := test.tree.Strings()

		if gotResult != test.expectedResult ||
			!reflect.DeepEqual(gotTree, test.expectedTree) {
			t.Errorf(
				"%v.Cut(%v, %v)\nExpected %v,\n'%v'\nGot %v,\n'%v'",
				was,
				test.first,
				test.second,
				test.expectedResult,
				test.expectedTree,
				gotResult,
				gotTree,
			)
		}
	}
}

func alarm(t *testing.T, name string, was []string, first, second Vertex, expected, got interface{}) {
	t.Errorf(
		"%v.%s(%v, %v)\nExpected:\n'%v'\nGot:\n'%v'",
		was,
		name,
		first,
		second,
		expected,
		got,
	)
}

func testIsConnected(t *testing.T, tree *Euler, first, second Vertex, expected bool) {
	got := tree.IsConnected(first, second)

	if !reflect.DeepEqual(got, expected) {
		alarm(t, "IsConnected", tree.Strings(), first, second, expected, got)
	}
}

func testLink(t *testing.T, tree *Euler, first, second Vertex, expected []string) {
	was := tree.Strings()
	if !tree.Link(first, second) {
		panic("invalid test")
	}
	got := tree.Strings()

	if !reflect.DeepEqual(got, expected) {
		alarm(t, "Link", was, first, second, expected, got)
	}
}

func testCut(t *testing.T, tree *Euler, first, second Vertex, expected []string) {
	was := tree.Strings()
	if !tree.Cut(first, second) {
		panic("invalid test")
	}
	got := tree.Strings()

	if !reflect.DeepEqual(got, expected) {
		alarm(t, "Cut", was, first, second, expected, got)
	}
}

func TestEuler(t *testing.T) {
	tree := CreateEuler()

	testLink(t, tree, 1, 2, []string{"1-2-1"})
	testLink(t, tree, 3, 4, []string{"1-2-1", "3-4-3"})

	testIsConnected(t, tree, 1, 3, false)

	testLink(t, tree, 2, 3, []string{"1-2-3-4-3-2-1"})
	testLink(t, tree, 5, 3, []string{"5-3-2-1-2-3-4-3-5"})

	testIsConnected(t, tree, 2, 5, true)

	testCut(t, tree, 5, 3, []string{"3-2-1-2-3-4-3", "5"})
	testCut(t, tree, 2, 3, []string{"2-1-2", "3-4-3", "5"})

	testIsConnected(t, tree, 4, 1, false)

	testLink(t, tree, 3, 2, []string{"3-4-3-2-1-2-3", "5"})

	testIsConnected(t, tree, 1, 4, true)
	testIsConnected(t, tree, 3, 5, false)
}
