package euler

import (
	"reflect"
	"testing"
)

func TestTreap_Split(t *testing.T) {
	leafTreap := &Treap{priority: 1, size: 1}

	tests := []struct {
		treap    *Treap
		k        int
		expected TreapPair
	}{
		{
			&Treap{
				priority: 3,
				size:     3,
				left: &Treap{
					priority: 2,
					size:     2,
					right:    leafTreap,
				},
			},
			1,
			TreapPair{
				First:  &Treap{priority: 2, size: 1},
				Second: &Treap{priority: 3, size: 2, left: leafTreap},
			},
		},
		{
			&Treap{
				priority: 3,
				size:     3,
				left: &Treap{
					priority: 2,
					size:     2,
					right:    leafTreap,
				},
			},
			2,
			TreapPair{
				First:  &Treap{priority: 2, size: 2, right: leafTreap},
				Second: &Treap{priority: 3, size: 1},
			},
		},
		{
			&Treap{priority: 2, size: 2, left: leafTreap},
			1,
			TreapPair{
				First:  leafTreap,
				Second: &Treap{priority: 2, size: 1},
			},
		},
		{
			&Treap{priority: 2, size: 2, right: leafTreap},
			1,
			TreapPair{
				First:  &Treap{priority: 2, size: 1},
				Second: leafTreap,
			},
		},
	}

	for _, test := range tests {
		got := test.treap.Split(test.k)

		if !reflect.DeepEqual(got, test.expected) {
			t.Errorf("Expected %v,\ngot %v", test.expected, got)
		}
	}
}

func TestTreap_Merge(t *testing.T) {
	first := &Treap{priority: 1, size: 1}
	firstParent := &Treap{priority: 2, size: 2, left: first}
	first.parent = firstParent

	second := &Treap{priority: 1, size: 1}
	secondParent := &Treap{priority: 2, size: 2, right: second}
	second.parent = secondParent

	thirdLeft := &Treap{priority: 3, size: 2, right: &Treap{priority: 2, size: 1}}
	thirdLeft.right.setParent(thirdLeft)
	thirdRight := &Treap{priority: 1, size: 1}
	thirdResult := &Treap{
		priority: 3,
		size:     3,
		right:    &Treap{priority: 2, size: 2, right: &Treap{priority: 1, size: 1}},
	}
	thirdResult.right.setParent(thirdResult)
	thirdResult.right.right.setParent(thirdResult.right)

	tests := []struct {
		first, second *Treap
		expected      *Treap
	}{
		{
			&Treap{priority: 1, size: 1},
			&Treap{priority: 2, size: 1},
			firstParent,
		},
		{
			&Treap{priority: 2, size: 1},
			&Treap{priority: 1, size: 1},
			secondParent,
		},
		{
			nil,
			&Treap{priority: 2, size: 1},
			&Treap{priority: 2, size: 1},
		},
		{
			thirdLeft,
			thirdRight,
			thirdResult,
		},
	}

	for _, test := range tests {
		got := Merge(test.first, test.second)

		if !reflect.DeepEqual(got, test.expected) {
			t.Errorf("Expected %v,\ngot %v", test.expected, got)
		}
	}
}

func TestTreap_FindRoot(t *testing.T) {
	child := &Treap{priority: 1, size: 1}
	parent := Merge(child, &Treap{priority: 2, size: 1})

	tests := []struct {
		in       *Treap
		expected *Treap
	}{
		{
			child,
			parent,
		},
		{
			parent,
			parent,
		},
		{
			nil,
			nil,
		},
	}

	for _, test := range tests {
		got := test.in.Root()

		if !reflect.DeepEqual(got, test.expected) {
			t.Errorf("Expected %v,\ngot %v", test.expected, got)
		}
	}
}

func TestTreap_Stringify(t *testing.T) {
	v0 := 0
	v1 := 1
	v2 := 2

	tests := []struct {
		in       *Treap
		expected string
	}{
		{
			nil,
			"",
		},
		{
			&Treap{vertex: v0},
			"0",
		},
		{
			&Treap{
				vertex: v2,
				right:  &Treap{vertex: v1},
			},
			"2-1",
		},
		{
			&Treap{
				vertex: v2,
				left:   &Treap{vertex: v1},
				right:  &Treap{vertex: v0},
			},
			"1-2-0",
		},
	}

	for _, test := range tests {
		got := test.in.Stringify()

		if got != test.expected {
			t.Errorf("Expected %v,\ngot %v", test.expected, got)
		}
	}
}

