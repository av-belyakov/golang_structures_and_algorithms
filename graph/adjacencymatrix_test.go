package graph_test

import (
	"testing"

	"github.com/av-belyakov/golang_structures_and_algorithms/graph"
)

func TestAdjacencyMatrix(t *testing.T) {
	g := graph.NewGraphAM(4)

	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 2)
	g.AddEdge(2, 0)
	g.AddEdge(2, 3)

	g.PrintGraphAM()
}
