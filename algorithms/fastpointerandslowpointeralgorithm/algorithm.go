package fastpointerandslowpointeralgorithm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Node struct {
	data int
	next *Node
}

// алгоритм быстрого и медленного указателей подходит для обнаружение циклов
// в структуре данных linked list (связный список)
// подробнее https://algodaily.com/lessons/using-the-two-pointer-technique/go

func hasCycle(head *Node) bool {
	fast, slow := head, head

	for fast != nil && fast.next != nil {
		fast = fast.next.next
		slow = slow.next
		if fast == slow {
			return true
		}
	}

	return false
}

func TestHsCycle(t *testing.T) {
	parent := &Node{data: 1}
	child := &Node{data: 2, next: parent}

	parent.next = child

	assert.True(t, hasCycle(parent))
}
