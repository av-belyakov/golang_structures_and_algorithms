package lineardatastructures

import "fmt"

//Очередь состоит из элементов, подлежащих обработке в определенном порядке или на основе приоритета.
//Очередь заказов на основе приоритета показана в следующем коде, структурированном в виде кучи.
//Такие операции, как постановка в очередь, удаление из очереди и просмотр, могут выполняться в очереди. Очередь - это
//линейная структура данных и последовательная коллекция. Элементы добавляются в конец и удаляются из начала коллекции.
//Очереди обычно используются для хранения задач, которые необходимо выполнить, или входящих HTTP-запросов, которые
//должны быть обработаны сервером. В реальном времени жизнь, обработка прерываний в системах реального времени,
//обработка вызовов и планирование задач процессора являются хорошими примерами использования очередей.

// Queues—Array of Orders Type
type Queue []*Order

// Order class
type Order struct {
	Priority     int
	Quantity     int
	Product      string
	CustomerName string
}

func (order *Order) New(priority, quantity int, product, customerName string) {
	order.Priority = priority
	order.Quantity = quantity
	order.Product = product
	order.CustomerName = customerName
}

// Add метод добавляет новый список в очередь
func (queue *Queue) Add(order *Order) {
	if len(*queue) == 0 {
		*queue = append(*queue, order)

		return
	}

	var (
		i          int
		appended   bool
		addedOrder *Order
	)

	for i, addedOrder = range *queue {
		if order.Priority > addedOrder.Priority {
			*queue = append((*queue)[:i], append(Queue{order}, (*queue)[i:]...)...)
			appended = true

			break
		}
	}

	if !appended {
		*queue = append(*queue, order)
	}
}

func QueuesExample() {
	queue := make(Queue, 0)

	order1 := &Order{}
	order1.New(2, 20, "Computer", "Greg White")

	order2 := &Order{}
	order2.New(1, 10, "Monitor", "John Smith")

	queue.Add(order1)
	queue.Add(order2)

	for i := 0; i < len(queue); i++ {
		fmt.Println(queue[i])
	}
}
