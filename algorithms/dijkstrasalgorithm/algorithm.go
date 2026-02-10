package dijkstrasalgorithm

import (
	"container/heap"
	"math"
)

func DijkstrasAlgorithm(graph [][]Edge, start int) ([]int, []int) {
	n := len(graph)

	// Инициализация расстояний
	distances := make([]int, n)
	previous := make([]int, n)

	for i := 0; i < n; i++ {
		distances[i] = math.MaxInt32
		previous[i] = -1
	}
	distances[start] = 0

	// Создаем приоритетную очередь
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	// Помещаем начальную вершину в очередь
	heap.Push(&pq, &Item{
		Node:     start,
		Distance: 0,
	})

	// Основной цикл алгоритма
	for pq.Len() > 0 {
		// Извлекаем вершину с минимальным расстоянием
		currentItem := heap.Pop(&pq).(*Item)
		currentNode := currentItem.Node
		currentDist := currentItem.Distance

		// Если извлеченное расстояние больше текущего, пропускаем
		if currentDist > distances[currentNode] {
			continue
		}

		// Проверяем всех соседей текущей вершины
		for _, edge := range graph[currentNode] {
			neighbor := edge.Node
			newDist := currentDist + edge.Weight

			// Если нашли более короткий путь
			if newDist < distances[neighbor] {
				distances[neighbor] = newDist
				previous[neighbor] = currentNode

				// Добавляем соседа в очередь
				heap.Push(&pq, &Item{
					Node:     neighbor,
					Distance: newDist,
				})
			}
		}
	}

	return distances, previous
}

// Восстановление пути из start в end
func ReconstructPath(previous []int, start, end int) []int {
	if previous[end] == -1 && start != end {
		return nil // Пути не существует
	}

	path := make([]int, 0)
	for at := end; at != -1; at = previous[at] {
		path = append([]int{at}, path...)
	}

	// Проверяем, что путь начинается с start
	if path[0] == start {
		return path
	}

	return nil
}
