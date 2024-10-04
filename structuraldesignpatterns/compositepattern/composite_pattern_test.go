package structuraldesignpatterns

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompositeExample(t *testing.T) {
	var (
		branch  = &Branch{name: "branch 1"}
		leaf1   = Leaflet{name: "leaf 1"}
		leaf2   = Leaflet{name: "leaf 2"}
		branch2 = Branch{name: "branch 2"}
		leaf3   = Leaflet{name: "leaf 3"}
		leaf4   = Leaflet{name: "leaf 4"}
	)

	branch2.addLeaf(leaf3)
	branch2.addLeaf(leaf4)

	branch.addLeaf(leaf1)
	branch.addLeaf(leaf2)
	branch.addBranch(branch2)

	result := branch.perform()
	a := []string{"Branch: branch 1", "Leaflet leaf 1", "Leaflet leaf 2", "Branch: branch 2", "Leaflet leaf 3", "Leaflet leaf 4"}

	assert.Equal(t, a, result)
}
