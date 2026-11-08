package graph

// DFS traverses the graph starting from node `s` in depth-first order
// and returns the visited nodes in order.
func (g *Graph) DFS(s int) []int {
	visited := make(map[int]bool)
	var order []int

	var dfsHelper func(u int)
	dfsHelper = func(u int) {
		visited[u] = true
		order = append(order, u)

		for _, neighbor := range g.Neighbors(u) {
			if !visited[neighbor] {
				dfsHelper(neighbor)
			}
		}
	}

	dfsHelper(s)

	return order
}
