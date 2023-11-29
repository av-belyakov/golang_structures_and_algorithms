package structuraldesignpatterns

import "fmt"

//	Шаблон "Декоратор" (Decorator)
//
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
func (process *ProcessClass) process() {
	fmt.Println("ProcessClass process")
}

// ProcessDecorator struct
type ProcessDecorator struct {
	processInstance *ProcessClass
}

// ProcessDecorator class method process
func (decorator *ProcessDecorator) process() {
	if decorator.processInstance == nil {
		fmt.Println("ProcessDecorator process")
	} else {
		fmt.Printf("ProcessDecorator process and ")
		decorator.processInstance.process()
	}
}

// DecoratorExample method
func DecoratorExample() {
	var process = &ProcessClass{}
	var decorator = &ProcessDecorator{}
	decorator.process()
	decorator.processInstance = process
	decorator.process()
}
