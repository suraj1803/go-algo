package main

import (
	"fmt"

	"algorithms/graph"
)

func main() {
	fmt.Println("--- Graph: BFS Traversal ---")
	g := graph.New()
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(3, 4)
	g.AddEdge(3, 5)

	traversal := g.BFS(1)
	fmt.Printf("BFS traversal from 1: %v\n\n", traversal)

	path := g.ShortestPath(1, 5)
	fmt.Printf("Shortest path from 1 to 5: %v\n", path) // Should print: [1 3 5]
}
