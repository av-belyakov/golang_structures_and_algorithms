package lineardatastructures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedListExample(t *testing.T) {
	ll := LinkedList{}

	ll.AddToHead(1)
	ll.AddToHead(3)

	assert.Equal(t, 3, ll.headNode.property)
}
