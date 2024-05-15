package htmx

// RangeLoop is a loop control structure.
type RangeLoop interface {
	// Filter loops and filters the content.
	Filter(f func(int) bool) RangeLoop
	// Map loops and maps the content.
	Map(f func(int) Node) RangeLoop
	// Group returns the nodes as a group.
	Group() Node
}

type rangeLoop struct {
	nodes []Node
	src   []Node
}

// Range loops over the content.
func Range(nodes ...Node) RangeLoop {
	return &rangeLoop{nodes: []Node{}, src: nodes}
}

// Group returns the nodes as a group.
func (r *rangeLoop) Group() Node {
	return Group(r.nodes...)
}

// Filter loops and slices the content.
func (r *rangeLoop) Filter(f func(int) bool) RangeLoop {
	for i := range r.src {
		if f(i) {
			r.nodes = append(r.nodes, r.src[i])
		}
	}

	return r
}

// Map loops and maps the content.
func (r *rangeLoop) Map(f func(int) Node) RangeLoop {
	for i := range r.src {
		r.nodes[i] = f(i)
	}

	return r
}

// Filter loops and filters the content.
func Filter(f func(n Node) bool, children ...Node) []Node {
	var nodes []Node
	for i := 0; i < len(children); i++ {
		if f(children[i]) {
			nodes = append(nodes, children[i])
		}
	}

	return nodes
}

// Map loops and maps the content.
func Map(f func(i int) Node, children ...Node) []Node {
	nodes := make([]Node, len(children))

	for i := 0; i < len(children); i++ {
		nodes[i] = f(i)
	}

	return nodes
}

// Reduce reduces the content to a single node.
func Reduce(f func(prev Node, next Node) Node, children ...Node) Node {
	if len(children) == 0 {
		return nil
	}

	node := children[0]
	for _, n := range children[1:] {
		node = f(node, n)
	}

	return node
}
