package graph


// dfs of undirected graph
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

func (g *Graph) FindPath(start, target int) []int {
	// Base case: start and target are the same node
	if start == target {
		return []int{start}
	}
	visited := make(map[int]bool)
	path := []int{}

	var dfs func(u int)  bool
	dfs = func(u int ) bool {
		visited[u] = true
		path = append(path, u)

		if u == target {
			return true
		}

		for _, neighbor := range g.Neighbors(u) {
			if !visited[neighbor] {
				if dfs(neighbor) {
					return true
				}
			}
		}

		// dead end : backtrack by removing 'u'
		path = path[:len(path) - 1]
		return false
	}

	if dfs(start) {
		return path
	}

	return  nil
}
