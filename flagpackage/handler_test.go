package flagpackage_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/av-belyakov/golang_structures_and_algorithms/flagpackage"
)

func TestFlagsHandler(t *testing.T) {
	flagpackage.FlagsHandler()

	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}

	fmt.Println("\nALL ARGUMENTS COMMAND STRING:")
	fmt.Println(s)
}
