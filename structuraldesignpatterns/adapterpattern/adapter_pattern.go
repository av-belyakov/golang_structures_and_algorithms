// - шаблон "Адаптер" (Adapter)
package structuraldesignpatterns

import (
	"fmt"
	"strconv"
)

//Шаблон адаптера предоставляет оболочку с интерфейсом, требуемым клиентом API для связывания
//несовместимых типов и выполнения функций транслятора между двумя типами. Адаптер использует
//интерфейс класса, чтобы быть классом с другим совместимым интерфейсом. Когда требования
//меняются, возникают сценарии, в которых функциональность класса должна быть изменена из-за
//несовместимых интерфейсов.

// IProcess interface
type IAProcess interface {
	process() string
}

// Adapter struct
type Adapter struct {
	adaptee Adaptee
}

// Adapter class method process
func (adapter Adapter) process() string {
	fmt.Println("Adapter process")
	return adapter.adaptee.convert()
}

// Adaptee Struct
type Adaptee struct {
	adapterType int
}

// Adaptee class method convert
func (adaptee Adaptee) convert() string {
	fmt.Println("Adaptee convert")
	return strconv.Itoa(adaptee.adapterType)
}

// AdapterExample method
func AdapterExample(num int) string {
	var processor IAProcess = Adapter{adaptee: Adaptee{num}}

	return processor.process()
}

//За счет пользовательского типа Adapter можно выполнить метод convert
//другого вложенного пользовательского типа Adaptee, а интерфейс IProcess
//позволит абстрагироватся от конкретных типов в пользу определенных методов.
