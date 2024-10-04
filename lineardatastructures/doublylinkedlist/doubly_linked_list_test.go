package lineardatastructures

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoubleLinkedListExample(t *testing.T) {
	var linkedList DoubleLinkedList = DoubleLinkedList{}

	linkedList.AddToHead(1)
	linkedList.AddToHead(3)
	linkedList.AddToEnd(5)
	linkedList.AddAfter(1, 7)

	fmt.Println("list property = ", linkedList.GetPropertyList())

	node := linkedList.NodeBetweenValues(1, 5)
	assert.Equal(t, 7, node.property)
}
