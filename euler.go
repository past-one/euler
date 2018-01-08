package euler

import (
	"math/rand"
	"strings"
	"sort"
)

// Euler structure that allows operations
//  IsConnected
//  Link
//  Cut
// with O(log(N)) complexity for any forest (acyclic graph) with int vertices
type Euler struct {
	treaps map[Vertex]*Treap
	edges  map[Vertex]map[Vertex]*Edge
}

// CreateEuler making empty tree
func CreateEuler() *Euler {
	return &Euler{
		treaps: make(map[Vertex]*Treap),
		edges:  make(map[Vertex]map[Vertex]*Edge),
	}
}

// IsConnected return true if vertices are in one treap
func (tree *Euler) IsConnected(first, second Vertex) bool {
	return tree.isConnected(tree.getTreap(first), tree.getTreap(second))
}

// Link creates edge in forest
//
// returns false if vertices are already linked
func (tree *Euler) Link(first, second Vertex) bool {
	firstTreap := tree.getTreap(first)
	secondTreap := tree.getTreap(second)
	if tree.isConnected(firstTreap, secondTreap) {
		return false
	}

	// in tree
	//  {3, 1-2-1}
	//  link(3, 2)

	// split with duplicates
	//  3  1-2-1 -> {3, 3}  {1-2, 2-1}
	part1, part4 := tree.splitByEntry(firstTreap, false, true).Destruct()
	// we need to split second treap to left if it's only one entry in treap (root == leaf)
	part3, part2 := tree.splitByEntry(secondTreap, !secondTreap.isOnlyOne(), true).Destruct()

	// remove duplicated entry
	//  {3, 3}  {[1-]2, 2-1} -> {3, 3}  {2, 2-1}
	var removing *Treap
	removing, part3 = part3.Split(1).Destruct()

	// relink vertex if needed
	if tree.getTreap(removing.vertex) == removing {
		tree.treaps[removing.vertex] = part2.rightmost()
	}

	// save edge for fast cutting
	firstEdgePart := part2.leftmost()
	//if firstEdgePart == nil {
	//	firstEdgePart = secondSplit.Second.rightmost()
	//}
	secondEdgePart := part4.leftmost()
	edge := &Edge{
		firstEdgePart,
		secondEdgePart,
	}
	firstEdgePart.edge = edge
	secondEdgePart.edge = edge
	tree.setEdge(first, second, edge)

	//  {3, 3}  {2, 2-1} -> 3-2-1-2-3
	Merge(Merge(part1, part2), Merge(part3, part4))

	return true
}

// Cut removes given edge
// return false if edge is not exist
func (tree *Euler) Cut(first, second Vertex) bool {
	// in tree
	//  {3-2-1-2-3}
	// cut(2, 1)

	//  edge(1, 2)
	//     | |
	//  3-2-1-2-3
	edge := tree.getEdge(first, second)
	if edge == nil {
		return false
	}

	// split left side
	//   edge(1, 2)
	//      | |
	//  3-2  1-2-3
	//
	// split right side
	//    edge(1, 2)
	//      | |
	//  3-2  1  2-3
	left, check := tree.splitByEntry(edge.First, true, false).Destruct()
	middle, right := tree.splitByEntry(edge.Second, true, false).Destruct()
	if check != middle && check != right {
		// wrong sides, swap
		left = middle
		right = check
	}

	// remove last entry of left treap
	//      edge(1, 2)
	//        | |
	//  3-(2)  1  2-3
	var removing *Treap
	left, removing = tree.splitByEntry(left.rightmost(), true, false).Destruct()

	// relink vertex if needed
	if tree.getTreap(removing.vertex) == removing {
		tree.treaps[removing.vertex] = right.leftmost()
	}

	// remove edge and unlink entries from it
	//  3-(2)  1  2-3
	tree.removeEdge(first, second)
	edge.First.edge = nil
	edge.Second.edge = nil

	// change link for edge of removing entry
	//    edge(2, 3)
	//   |          |
	//  3-(2)  {1, 2-3}
	//
	//    edge(2, 3)
	//        | |
	//  3  {1, 2-3}
	changeEdgeLink(removing, right.leftmost())

	// merge left and right sides
	//  edge(2, 3)
	//       | |
	//  3  1  2-3
	//
	//  edge(2, 3)
	//       | |
	//    1 3-2-3
	Merge(left, right)

	return true
}

// Strings O(N*log(N)) complexity
func (tree *Euler) Strings() (result []string) {
	uniqueTreapMap := make(map[int]*Treap, len(tree.treaps))
	for _, treap := range tree.treaps {
		root := treap.Root()
		uniqueTreapMap[root.vertex] = root
	}

	keys := make([]int, 0, len(uniqueTreapMap))
	for key := range uniqueTreapMap {
		keys = append(keys, key)
	}
	sort.Ints(keys)

	for _, key := range keys {
		result = append(result, uniqueTreapMap[key].Stringify())
	}

	return
}

// String representation
func (tree *Euler) String() string {
	return strings.Join(tree.Strings(), "\n")
}

func (tree *Euler) isConnected(first, second *Treap) bool {
	return first.Root() == second.Root()
}

func (tree *Euler) getTreap(v Vertex) *Treap {
	result, ok := tree.treaps[v]
	if !ok {
		result = &Treap{priority: rand.Int(), size: 1, vertex: v}
		tree.treaps[v] = result
	}
	return result
}

func (tree *Euler) setEdge(first, second Vertex, edge *Edge) {
	edgesMap, key := tree.getEdgesMap(first, second)
	edgesMap[key] = edge
}

func (tree *Euler) getEdge(first, second Vertex) *Edge {
	edgesMap, key := tree.getEdgesMap(first, second)
	return edgesMap[key]
}

func (tree *Euler) removeEdge(first, second Vertex) {
	edgesMap, key := tree.getEdgesMap(first, second)
	delete(edgesMap, key)
}

func (tree *Euler) getEdgesMap(first, second Vertex) (map[Vertex]*Edge, int) {
	// first should be smaller
	if first > second {
		first, second = second, first
	}

	edgesMap, ok := tree.edges[first]
	// init if needed
	if !ok {
		edgesMap = make(map[Vertex]*Edge)
		tree.edges[first] = edgesMap
	}

	return edgesMap, second
}

// splitByEntry split treap by two parts
//    2                              2
//   / \   >> split by 2 right >>     \
//  1   1                          1   1
//
//    2                              2
//   / \   >> split by 2 left >>    /
//  1   1                          1   1
//
// if make duplicate, it will be create in opposite side
//    2                                dup -> 2  2
//   / \   >> split by 2 right with dup >>   /    \
//  1   1                                   1      1
//
//    2                                      2  2 <- dup
//   / \   >> split by 2 left with dup >>   /    \
//  1   1                                  1      1
func (tree *Euler) splitByEntry(
	entry *Treap,
	splitToRight,
	makeDuplicate bool,
) TreapPair {
	// k - number of entries in left side from entry
	k := entry.left.getSize()
	current := entry
	for current.parent != nil {
		if current.parent.right == current {
			k += current.parent.left.getSize() + 1
		}
		current = current.parent
	}

	if !splitToRight {
		// also take entry to left
		k++
	}
	result := current.Split(k)

	if makeDuplicate {
		// add duplicate to opposite side
		if splitToRight {
			result.First = Merge(result.First, tree.duplicateTreap(entry, true))
		} else {
			// we don't need to relink edge if split moves left and duplicate moves to right
			// because it will have another edge
			result.Second = Merge(tree.duplicateTreap(entry, false), result.Second)
		}
	}

	return result
}

// duplication relink vertex [and edge] in Euler struct
func (tree *Euler) duplicateTreap(t *Treap, relinkEdge bool) *Treap {
	result := &Treap{priority: rand.Int(), size: 1, vertex: t.vertex}
	tree.treaps[t.vertex] = result
	if relinkEdge {
		changeEdgeLink(t, result)
	}
	return result
}

func changeEdgeLink(from *Treap, to *Treap) {
	if from.edge != nil {
		if from.edge.First == from {
			from.edge.First = to
		} else {
			from.edge.Second = to
		}
		to.edge = from.edge
		from.edge = nil
	}
}
