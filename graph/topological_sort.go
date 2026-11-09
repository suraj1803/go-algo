package graph

// TopologicalSort performs a topological sort on a Directed Acyclic Graph (DAG)
// using DFS post-order labeling.
// It assigns decreasing labels (n, n-1, ..., 1) to vertices when all of their
// outgoing edges have been fully explored.
// Repeats this for all vertices to cover disconnected components.
func (g *Graph) TopologicalSort() []int {
	vertices := g.Vertices()
	n := len(vertices)
	if n == 0 {
		return nil
	}

	visited := make(map[int]bool)
	order := make([]int, n)
	curLabel := n 

	var dfs func(u int)
	dfs = func(u int) {
		visited[u] = true

		for _, neighbor := range g.Neighbors(u) {
			if !visited[neighbor] {
				dfs(neighbor)
			}
		}

		curLabel--
		order[curLabel] = u
	}

	for _, v := range vertices {
		if !visited[v] {
			dfs(v)
		}
	}

	return order
}
