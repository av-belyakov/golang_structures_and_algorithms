// - шаблон "Декоратор" (Decorator)
package structuraldesignpatterns

import (
	"fmt"
	"strings"
)

//В сценарии, где обязанности класса удаляются или добавляются, применяется шаблон декоратора.
//Шаблон декоратора помогает при создании подклассов при изменении функциональности
//вместо статического наследования. Объект может иметь несколько декораторов и
//декораторов времени выполнения. Принцип единой ответственности может быть реализован с помощью
//декоратора. Декоратор может быть применен к компонентам окна и графическому моделированию
//объектов. Шаблон декоратора помогает изменять существующие атрибуты экземпляра и добавлять
//новые методы во время выполнения.

// IProcess Interface
type IDProcess interface {
	process()
}

// ProcessClass struct
type ProcessClass struct{}

// ProcessClass method process
func (process *ProcessClass) process(str string, repeat int) string {
	fmt.Println("ProcessClass process")

	return strings.Repeat(str, repeat)
}

// ProcessDecorator struct
type ProcessDecorator struct {
	processInstance *ProcessClass
}

// ProcessDecorator class method process
func (decorator *ProcessDecorator) process(str string, repeat int) string {
	if decorator.processInstance == nil {
		fmt.Println("ProcessDecorator process")

		return str
	}

	fmt.Printf("ProcessDecorator process and ")

	return decorator.processInstance.process(str, repeat)
}

// DecoratorExample method
func DecoratorExample() {
	var process = &ProcessClass{}
	var decorator = &ProcessDecorator{}
	decorator.process("hotdog", 3)
	decorator.processInstance = process
	decorator.process("hotdog", 5)
}
