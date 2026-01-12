package flagpackage_test

import (
	"flag"
	"fmt"
	"os"
	"testing"

	"github.com/av-belyakov/golang_structures_and_algorithms/flagpackage"
	"github.com/stretchr/testify/assert"
)

var (
	flagName  string
	flagCount int
)

func TestMain(m *testing.M) {
	flag.StringVar(&flagName, "name", "xcccc", "name flag for test")
	flag.IntVar(&flagCount, "count", 12, "count flag for test")

	flag.Parse()

	os.Exit(m.Run())
}

func TestFlagsHandler(t *testing.T) {
	flagpackage.FlagsHandler()

	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}

	fmt.Println("\nALL ARGUMENTS COMMAND STRING:")
	fmt.Println(s)

	t.Run("Тест 1. Количество аргументов", func(t *testing.T) {
		assert.Greater(t, len(os.Args), 1)
	})

	t.Run("Тест 2. Наличие аргумента '-name' и '-count'", func(t *testing.T) {

		//
		// При выполнении go test пакет flag не работает корректно. Несмотря
		// на использование параметра -- перед пользовательскими флагами,
		// не могу получить доступ к '-name' и '-count'. Однако все эти аргументы
		// присутствуют в os.Args.
		//

		fmt.Println("flag name:", flagName, " --- ", flagName)
		fmt.Println("flag count:", flagCount, " --- ", flagCount)
		fmt.Println("flag args:", flag.Args())

		assert.Equal(t, flagName, "example-name")
	})
}
