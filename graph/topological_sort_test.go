package graph

import (
	"testing"
)

func TestTopologicalSort(t *testing.T) {
	t.Run("Connected DAG", func(t *testing.T) {
		g := New()
		edges := [][2]int{
			{5, 2},
			{5, 0},
			{4, 0},
			{4, 1},
			{2, 3},
			{3, 1},
		}

		for _, e := range edges {
			g.AddDirectedEdge(e[0], e[1])
		}

		order := g.TopologicalSort()
		if len(order) != len(g.Vertices()) {
			t.Fatalf("Expected %d vertices in order, got %d", len(g.Vertices()), len(order))
		}

		// Verify every edge u -> v has u before v in order
		pos := make(map[int]int)
		for i, v := range order {
			pos[v] = i
		}

		for _, e := range edges {
			u, v := e[0], e[1]
			if pos[u] >= pos[v] {
				t.Errorf("Invalid ordering for edge %d -> %d: %d appears at index %d, %d appears at index %d",
					u, v, u, pos[u], v, pos[v])
			}
		}
	})

	t.Run("Disconnected DAG", func(t *testing.T) {
		g := New()
		// Component 1: 1 -> 2
		g.AddDirectedEdge(1, 2)
		// Component 2: 3 -> 4
		g.AddDirectedEdge(3, 4)

		order := g.TopologicalSort()
		if len(order) != 4 {
			t.Fatalf("Expected 4 vertices in order, got %d", len(order))
		}

		pos := make(map[int]int)
		for i, v := range order {
			pos[v] = i
		}

		if pos[1] >= pos[2] {
			t.Errorf("Expected 1 before 2, got %d and %d", pos[1], pos[2])
		}
		if pos[3] >= pos[4] {
			t.Errorf("Expected 3 before 4, got %d and %d", pos[3], pos[4])
		}
	})
}
