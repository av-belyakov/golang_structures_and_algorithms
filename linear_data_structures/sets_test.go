package lineardatastructures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetsExample(t *testing.T) {
	var set *Set = &Set{}
	set.New()

	set.AddElement(1)
	set.AddElement(2)
	set.AddElement(10)
	assert.Equal(t, true, set.ContainsElement(1))

	set.DeleteElement(1)
	assert.Equal(t, false, set.ContainsElement(1))

	var anotherSet *Set = &Set{}
	anotherSet.New()

	anotherSet.AddElement(2)
	anotherSet.AddElement(3)
	anotherSet.AddElement(4)

	var interSet *Set = set.InterSect(anotherSet)
	assert.Equal(t, true, interSet.ContainsElement(2))
	assert.Equal(t, false, interSet.ContainsElement(3))

	var unitedSet *Set = set.Union(anotherSet)
	assert.Equal(t, 4, len(unitedSet.integerMap))
}
