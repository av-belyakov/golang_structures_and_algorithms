package structuraldesignpatterns

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingletonExample(t *testing.T) {
	mcOne := NewSingletonMemoryCache()
	mcOne.Alloc = 624624
	assert.Equal(t, mcOne.Alloc, uint64(624624))

	mcTwo := NewSingletonMemoryCache()

	assert.Equal(t, mcTwo.Alloc, uint64(624624))
	mcTwo.Alloc = 6742845
	assert.Equal(t, mcTwo.Alloc, uint64(6742845))
}
