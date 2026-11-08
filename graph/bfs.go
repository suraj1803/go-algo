package graph

import "slices"

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

// shortest path undirected graph
func (g *Graph) ShortestPath(start, target int) []int {
	return nil
}

