package dijkstrasalgorithm

// Структура для представления ребра графа
type Edge struct {
	Node   int
	Weight int
}

// Структура для приоритетной очереди (min-heap)
type Item struct {
	Node     int
	Distance int
	Index    int
}

type PriorityQueue []*Item
