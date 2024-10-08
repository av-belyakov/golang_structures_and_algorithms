// - линейная структура данных типа "Двойной связанный список" (Doubly Linked List)
package lineardatastructures

import "fmt"

//В списке с двойной связью все узлы имеют указатель на узел, к которому они подключены, по обе
//стороны от них в списке. Это означает, что каждый узел подключен к двум узлам, и мы можем
//переходить вперед к следующему узлу или назад к предыдущему узлу.
//Двусвязные списки допускают операции вставки, удаления и, очевидно, обхода.

// DoubleLinkedList class
type DoubleLinkedList struct {
	headNode *DoubleLinkedNode
}

// DoubleLinkedList class
type DoubleLinkedNode struct {
	property     int
	nextNode     *DoubleLinkedNode
	previousNode *DoubleLinkedNode
}

// NodeBetweenValues поиск ноды содержащей значение попадающее между минимальным и
// максимальным значением
func (linkedList *DoubleLinkedList) NodeBetweenValues(firstProperty int, secondProperty int) *DoubleLinkedNode {
	var nodeWith *DoubleLinkedNode

	for node := linkedList.headNode; node != nil; node = node.nextNode {
		if node.previousNode != nil && node.nextNode != nil {
			if node.previousNode.property == firstProperty && node.nextNode.property == secondProperty {
				nodeWith = node

				break
			}
		}
	}

	return nodeWith
}

// NodeWithValue поиск ноды с заданым значение
func (linkedList *DoubleLinkedList) NodeWithValue(property int) *DoubleLinkedNode {
	var nodeWith *DoubleLinkedNode

	for node := linkedList.headNode; node != nil; node = node.nextNode {
		if node.property == property {
			nodeWith = node

			break
		}
	}

	return nodeWith
}

// DoubleLinkedList возвращает последнюю ноду из списка
func (linkedList *DoubleLinkedList) LastNode() *DoubleLinkedNode {
	var lastNode *DoubleLinkedNode

	for node := linkedList.headNode; node != nil; node = node.nextNode {
		if node.nextNode == nil {
			lastNode = node

			break
		}
	}

	return lastNode
}

// AddToHead добавляет свойство в заголовок ноды
func (linkedList *DoubleLinkedList) AddToHead(property int) {
	var node = &DoubleLinkedNode{
		property: property,
		nextNode: nil,
	}

	if linkedList.headNode != nil {
		node.nextNode = linkedList.headNode
		linkedList.headNode.previousNode = node
	}

	linkedList.headNode = node
}

// AddAfter добавляет новую ноду после ноды с указанным значением
func (linkedList *DoubleLinkedList) AddAfter(nodeProperty int, property int) {
	var node = &DoubleLinkedNode{
		property: property,
		nextNode: nil,
	}

	nodeWith := linkedList.NodeWithValue(nodeProperty)

	if nodeWith != nil {
		node.nextNode = nodeWith.nextNode
		node.previousNode = nodeWith
		nodeWith.nextNode = node
	}
}

// AddToEnd добавляет ноду в конец списка
func (linkedList *DoubleLinkedList) AddToEnd(property int) {
	var node = &DoubleLinkedNode{
		property: property,
		nextNode: nil,
	}

	lastNode := linkedList.LastNode()

	if lastNode != nil {
		lastNode.nextNode = node
		node.previousNode = lastNode
	}
}

// GetPropertyList возвращает список свойств из всех нод
func (linkedList *DoubleLinkedList) GetPropertyList() []int {
	list := []int{}

	for node := linkedList.headNode; node != nil; node = node.nextNode {
		list = append(list, node.property)
	}

	return list
}

func DoubleLinkedListExample() {
	var linkedList DoubleLinkedList = DoubleLinkedList{}

	linkedList.AddToHead(1)
	linkedList.AddToHead(3)
	linkedList.AddToEnd(5)
	linkedList.AddAfter(1, 7)

	fmt.Println(linkedList.headNode.property)

	node := linkedList.NodeBetweenValues(1, 5)
	fmt.Println(node.property)
}
