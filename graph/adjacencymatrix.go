package graph

import "fmt"

type GraphAM struct {
	adjMatrix   [][]bool
	numVertices int
}

func NewGraphAM(numVertices int) *GraphAM {
	adjMatrix := make([][]bool, numVertices)

	for i := range adjMatrix {
		adjMatrix[i] = make([]bool, numVertices)
	}

	return &GraphAM{adjMatrix: adjMatrix, numVertices: numVertices}
}

func (g *GraphAM) AddEdge(i int, j int) {
	g.adjMatrix[i][j] = true
	g.adjMatrix[j][i] = true
}

func (g *GraphAM) RemoveEdge(i int, j int) {
	g.adjMatrix[i][j] = false
	g.adjMatrix[j][i] = false
}

func (g *GraphAM) PrintGraphAM() {
	for i := 0; i < g.numVertices; i++ {
		fmt.Print(i, ": ")

		for _, j := range g.adjMatrix[i] {
			fmt.Print(map[bool]int{false: 0, true: 1}[j], " ")
		}

		fmt.Println()
	}
}
