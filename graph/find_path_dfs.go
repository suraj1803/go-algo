package graph

// FindPathDFS finds a path between start and target in a graph using DFS backtracking.
// Returns the sequence of nodes from start to target, or nil if no path exists.
func (g *Graph) FindPathDFS(start, target int) []int {
	// Base case: start and target are the same node
	if start == target {
		return []int{start}
	}
	visited := make(map[int]bool)
	path := []int{}

	var dfs func(u int) bool
	dfs = func(u int) bool {
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
		path = path[:len(path)-1]
		return false
	}

	if dfs(start) {
		return path
	}

	return nil
}
