package lineardatastructures

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueuesExample(t *testing.T) {
	//****
	// Simple Queues
	//****
	queue := make(Queue, 0)

	order1 := &Order{}
	priority := 3
	quantity := 30
	product := "Mouse"
	customerName := "Hex Grey"
	order1.New(priority, quantity, product, customerName)

	order2 := &Order{}
	order2.New(2, 20, "Computer", "Greg White")

	order3 := &Order{}
	order3.New(1, 10, "Monitor", "John Smith")

	queue.Add(order3)
	queue.Add(order1)
	queue.Add(order2)

	var result *Order
	for i := 0; i < 1; i++ {
		result = queue[i]
	}

	fmt.Println(result)

	assert.Equal(t, priority, result.Priority)
	assert.Equal(t, quantity, result.Quantity)
	assert.Equal(t, product, result.Product)
	assert.Equal(t, customerName, result.CustomerName)

	//****
	// Sync Queues
	//****
	queueSync := &QueueSync{}
	queueSync.New()

	/*

		Здесь надо еще раз внимательно посмотреть алгоритм Sync Queues
		что бы написать тесты

	*/
}
