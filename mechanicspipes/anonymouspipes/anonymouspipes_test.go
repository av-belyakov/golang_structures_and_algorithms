// Анонимные конвееры
package mechanicsanonymouspipes_test

import (
	"os/exec"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnonymousPipes(t *testing.T) {
	//создание и запуск команды в косноле
	echoCmd := exec.Command("echo", "Hello, world!")
	//захват вывода echoCmd
	echoOutput, err := echoCmd.Output()
	assert.NoError(t, err)

	// создаем grep команду
	grepCmd := exec.Command("grep", "Hello")
	// перехватываем вывод echo
	grepCmd.Stdin = strings.NewReader(string(echoOutput))
	//создаем вывод команды grep
	grepOutput, err := grepCmd.Output()
	assert.NoError(t, err)

	assert.Equal(t, string(grepOutput), "Hello, world!\n")
}
