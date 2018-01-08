package euler

import (
	"strconv"
	"strings"
)

// Vertex id
type Vertex = int

// Treap with implicit key,
// also contain link to Edge
type Treap struct {
	priority            int
	size                int
	vertex              Vertex
	parent, left, right *Treap
	edge                *Edge
}

// TreapPair simple pair of treaps
type TreapPair struct {
	First, Second *Treap
}

func (p TreapPair) Destruct() (*Treap, *Treap) {
	return p.First, p.Second
}

// Edge it's only needed for Cut function
type Edge TreapPair

// Split k entries from left side of the treap
func (t *Treap) Split(k int) TreapPair {
	if t == nil {
		return TreapPair{}
	}
	if k == 0 {
		return TreapPair{Second: t}
	}
	l := t.left.getSize()
	if l >= k {
		pair := t.left.Split(k)
		t.left = pair.Second
		t.left.setParent(t)
		pair.First.setParent(nil)
		t.updateSize()
		pair.First.updateSize()
		return TreapPair{First: pair.First, Second: t}
	}
	pair := t.right.Split(k - l - 1)
	t.right = pair.First
	t.right.setParent(t)
	pair.Second.setParent(nil)
	t.updateSize()
	pair.Second.updateSize()
	return TreapPair{First: t, Second: pair.Second}
}

// Merge makes one treap from two
func Merge(first, second *Treap) *Treap {
	if second == nil {
		return first
	}
	if first == nil {
		return second
	}
	if first.priority > second.priority {
		first.right = Merge(first.right, second)
		first.right.setParent(first)
		first.updateSize()
		return first
	}
	second.left = Merge(first, second.left)
	second.left.setParent(second)
	second.updateSize()
	return second
}

// Root returns root of t or nil
func (t *Treap) Root() *Treap {
	if t == nil {
		return nil
	}
	current := t
	for current.parent != nil {
		current = current.parent
	}
	return current
}

// leftmost return leftmost (first) entry of t
func (t *Treap) leftmost() *Treap {
	current := t
	for current.left != nil {
		current = current.left
	}
	return current
}

// rightmost return rightmost (last) entry of t
func (t *Treap) rightmost() *Treap {
	current := t
	for current.right != nil {
		current = current.right
	}
	return current
}

// Stringify return euler tour tree string
func (t *Treap) Stringify() string {
	return strings.Join(t.stringify(), "-")
}

func (t *Treap) stringify() []string {
	if t == nil {
		return nil
	}
	str := strconv.Itoa(t.vertex)

	return append(append(t.left.stringify(), str), t.right.stringify()...)
}

func (t *Treap) isOnlyOne() bool {
	return t.left == nil && t.right == nil && t.parent == nil
}

func (t *Treap) getSize() int {
	if t == nil {
		return 0
	}
	return t.size
}

func (t *Treap) setParent(parent *Treap) {
	if t != nil {
		t.parent = parent
	}
}

func (t *Treap) updateSize() {
	if t != nil {
		t.size = 1 + t.left.getSize() + t.right.getSize()
	}
}
