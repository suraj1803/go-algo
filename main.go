package main

import (
	"fmt"
	"algorithms/graph"
)

func main() {
	g := graph.New()
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(3, 4)
	g.AddEdge(3, 5)

	fmt.Println("--- BFS ---")
	bfsOrder := g.BFS(1)
	fmt.Printf("BFS traversal from 1: %v\n", bfsOrder)
	shortestPath := g.ShortestPathBFS(1, 5)
	fmt.Printf("Shortest path from 1 to 5 (BFS): %v\n\n", shortestPath)

	fmt.Println("--- DFS ---")
	dfsOrder := g.DFS(2)
	fmt.Printf("DFS traversal from 1: %v\n", dfsOrder)

	path := g.FindPathDFS(2, 5)
	if path != nil {
		fmt.Printf("Path from 2 to 5 (DFS): %v\n", path)
	} else {
		fmt.Printf("No path exists between 2 and 5 :(\n")
	}

	fmt.Println("\n--- Topological Sort (DAG) ---")
	dag := graph.New()
	dag.AddDirectedEdge(5, 2)
	dag.AddDirectedEdge(5, 0)
	dag.AddDirectedEdge(4, 0)
	dag.AddDirectedEdge(4, 1)
	dag.AddDirectedEdge(2, 3)
	dag.AddDirectedEdge(3, 1)

	topoOrder := dag.TopologicalSort()
	fmt.Printf("Topological Sort: %v\n", topoOrder)
}
