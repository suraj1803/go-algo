package graph

// Graph represents an adjacency list graph with integer vertices.
type Graph struct {
	adj map[int][]int
}

// New creates and initializes an empty Graph.
func New() *Graph {
	return &Graph{
		adj: make(map[int][]int),
	}
}

// AddEdge adds an undirected edge between vertices u and v.
func (g *Graph) AddEdge(u, v int) {
	g.adj[u] = append(g.adj[u], v)
	g.adj[v] = append(g.adj[v], u)
}

// AddDirectedEdge adds a directed edge from vertex u to vertex v.
func (g *Graph) AddDirectedEdge(u, v int) {
	g.adj[u] = append(g.adj[u], v)
	if _, ok := g.adj[v]; !ok {
		g.adj[v] = []int{}
	}
}

// Neighbors returns the slice of neighbors adjacent to the given node.
func (g *Graph) Neighbors(node int) []int {
	return g.adj[node]
}

// Vertices returns a slice containing all vertices in the graph.
func (g *Graph) Vertices() []int {
	vertices := make([]int, 0, len(g.adj))
	for v := range g.adj {
		vertices = append(vertices, v)
	}
	return vertices
}
