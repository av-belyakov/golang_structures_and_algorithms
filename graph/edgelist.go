package graph

import "fmt"

type Edge struct {
	destVertex *Node
	weight     int
}

type Node struct {
	edgeList []*Edge
	name     string
}

type GraphEL struct {
	nodes map[*Node]bool
}

func NewEdge(dest *Node, w int) *Edge {
	return &Edge{destVertex: dest, weight: w}
}

func NewNode(name string) *Node {
	return &Node{name: name, edgeList: []*Edge{}}
}

func (n *Node) getName() string {
	return n.name
}

func (n *Node) getEdges() []*Edge {
	return n.edgeList
}

func NewGraphEL() *GraphEL {
	return &GraphEL{nodes: make(map[*Node]bool)}
}

func (g *GraphEL) Len() int {
	return len(g.nodes)
}

func (g *GraphEL) AddEdge(v1 *Node, v2 *Node, weight int) bool {
	edgesV1 := v1.getEdges()
	edgesV2 := v2.getEdges()

	edgesV1 = append(edgesV1, NewEdge(v2, weight))
	edgesV2 = append(edgesV2, NewEdge(v1, weight))

	return true
}

func (g *GraphEL) AddVertex(v *Node) bool {
	g.nodes[v] = true

	return true
}

func (g *GraphEL) PrintGraph() {
	for v := range g.nodes {
		fmt.Print("vertex name: " + v.getName() + ":\n")

		for _, e := range v.getEdges() {
			fmt.Print("destVertex: " + e.destVertex.getName() + ", weight: " + fmt.Sprint(e.weight) + "\n")
		}

		fmt.Print("\n")
	}
}
