package graph

import "slices"

// BFS traverses the graph starting from node `start` in breadth-first order
func (g *Graph) BFS(start int) []int {
	visited := make(map[int]bool)
	var order []int

	queue := []int{start}
	visited[start] = true

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		order = append(order, node)

		for _, neighbor := range g.adj[node] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}

	return order
}

func (g *Graph) ShortestPath(start, target int) []int {
	// Base case: start and target are the same node
	if start == target {
		return []int{start}
	}

	// Map to keep track of visited nodes
	visited := make(map[int]bool)

	parent := make(map[int]int)

	queue := []int{start}
	visited[start] = true
	found := false

	// BFS traversal
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if curr == target {
			found = true
			break
		}

		for _, neighbor := range g.Neighbors(curr) {
			if !visited[neighbor] {
				visited[neighbor] = true
				parent[neighbor] = curr
				queue = append(queue, neighbor)
			}
		}
	}

	if !found {
		return nil
	}

	var path []int

	// Backtrack from target to start
	for curr := target; curr != start; curr = parent[curr] {
		path = append(path, curr)
	}
	path = append(path, start)

	slices.Reverse(path)
	return path
}

