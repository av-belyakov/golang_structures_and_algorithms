package graph

type EdgeAL struct {
	Src, Dest, Weight int
}

type NodeAL struct {
	Value, Weight int
}

type GraphAL struct {
	Adj [][]NodeAL
}

func NewGraphAL(edges []EdgeAL) *GraphAL {
	adj := make([][]NodeAL, len(edges))

	for _, e := range edges {
		adj[e.Src] = append(adj[e.Src], NodeAL{e.Dest, e.Weight})
	}

	return &GraphAL{Adj: adj}
}
