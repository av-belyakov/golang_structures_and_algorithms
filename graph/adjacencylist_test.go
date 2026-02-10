package graph_test

import (
	"fmt"
	"testing"

	"github.com/av-belyakov/golang_structures_and_algorithms/graph"
)

func TestAdjacencyList(t *testing.T) {
	edges := []graph.EdgeAL{
		{0, 1, 6}, {1, 2, 7}, {2, 0, 5}, {2, 1, 4},
		{3, 2, 10}, {4, 5, 1}, {5, 4, 3},
	}

	graph := graph.NewGraphAL(edges)

	// print adjacency list representation of the graph
	PrintGraphAL(graph)
}

func PrintGraphAL(graph *graph.GraphAL) {
	src := 0
	n := len(graph.Adj)

	for src < n {
		for _, edge := range graph.Adj[src] {
			fmt.Printf("%d --> %d (%d)\t", src, edge.Value, edge.Weight)
		}

		fmt.Println()
		src++
	}
}
