// Линейная структура данных типа "Связанный список" (Linked List)
package lineardatastructures

import "fmt"

//LinkedList - это последовательность узлов, которые имеют свойства и ссылку на следующий узел в
//последовательности. Это линейная структура данных, которая используется для хранения данных. Структура
//данных позволяет добавлять и удалять компоненты из любого узла, расположенного рядом с другим узлом. Они
//не хранятся непрерывно в памяти, что создает различные массивы.

// Node class
type Node struct {
	property int
	nextNode *Node
}

// LinkedList class
type LinkedList struct {
	headNode *Node
}

// AddToHead method of LinkedList class
func (linkedList *LinkedList) AddToHead(property int) {
	var node = Node{}
	node.property = property
	if node.nextNode != nil {
		node.nextNode = linkedList.headNode
	}
	linkedList.headNode = &node
}

func LinkedListExample() {
	ll := LinkedList{}

	ll.AddToHead(1)
	ll.AddToHead(3)

	fmt.Println(ll.headNode.property)
}
