package dijkstrasalgorithm_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/av-belyakov/golang_structures_and_algorithms/algorithms/dijkstrasalgorithm"
)

func TestDijkstrasAlgorithm(t *testing.T) {
	// Пример графа в виде списка смежности
	// Вершины: 0, 1, 2, 3, 4
	graph := [][]dijkstrasalgorithm.Edge{
		0: {{1, 4}, {2, 1}}, // из 0 в 1 (вес 4) и в 2 (вес 1)
		1: {{3, 1}},         // из 1 в 3 (вес 1)
		2: {{1, 2}, {3, 5}}, // из 2 в 1 (вес 2) и в 3 (вес 5)
		3: {{4, 3}},         // из 3 в 4 (вес 3)
		4: {},               // вершина 4 не имеет исходящих ребер
	}

	start := 0
	end := 4

	distances, previous := dijkstrasalgorithm.DijkstrasAlgorithm(graph, start)

	fmt.Println("Кратчайшие расстояния от вершины", start, ":")
	for i, dist := range distances {
		if dist == math.MaxInt32 {
			fmt.Printf("  до %d: недостижима\n", i)
		} else {
			fmt.Printf("  до %d: %d\n", i, dist)
		}
	}

	fmt.Printf("\nПуть из %d в %d:\n", start, end)
	path := dijkstrasalgorithm.ReconstructPath(previous, start, end)

	if path == nil {
		fmt.Println("Путь не существует")
	} else {
		fmt.Print("  ")
		for i, node := range path {
			if i > 0 {
				fmt.Print(" → ")
			}
			fmt.Print(node)
		}
		fmt.Printf(" (длина: %d)\n", distances[end])
	}

	// Пример вывода:
	// 0 → 2 → 1 → 3 → 4 (длина: 7)
}
