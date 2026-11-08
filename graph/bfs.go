package graph

import (
	"slices"
)

// BFS traverses the graph starting from node `start` in breadth-first order
// and returns the visited nodes in order.
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

// ShortestPath finds the shortest path between start and target in an unweighted graph using BFS.
// Returns the sequence of nodes from start to target, or nil if no path exists.
func (g *Graph) ShortestPath(start, target int) []int {
	if start == target {
		return []int{start}
	}

	visited := make(map[int]bool)
	parent := make(map[int]int)

	queue := []int{start}
	visited[start] = true
	found := false

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

	// If target was unreachable
	if !found {
		return nil
	}

	var path []int

	for curr := target; curr != start; curr = parent[curr] {
		path = append(path, curr)
	}
	path = append(path, start)

	slices.Reverse(path)
	if len(path) > 0 {
		return path
	}
	return nil
}

