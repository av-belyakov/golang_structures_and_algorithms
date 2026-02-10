package graph_test

import (
	"testing"

	"github.com/av-belyakov/golang_structures_and_algorithms/graph"
	"github.com/stretchr/testify/assert"
)

func TestEdgeList(t *testing.T) {
	ourGraph := graph.NewGraphEL()

	// nodes
	v0 := graph.NewNode("0")
	v1 := graph.NewNode("1")
	v2 := graph.NewNode("2")
	v3 := graph.NewNode("3")

	ourGraph.AddVertex(v0)
	ourGraph.AddVertex(v1)
	ourGraph.AddVertex(v2)
	ourGraph.AddVertex(v3)

	// edges
	ourGraph.AddEdge(v0, v1, 2)
	ourGraph.AddEdge(v1, v2, 3)
	ourGraph.AddEdge(v2, v0, 1)
	ourGraph.AddEdge(v2, v3, 1)
	ourGraph.AddEdge(v3, v2, 4)

	ourGraph.PrintGraph()

	// граф состоит из 4 вершин
	assert.Equal(t, ourGraph.Len(), 4)
}
