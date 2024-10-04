// Шаблон Легчайший (Flyweight)
package structuraldesignpatterns

import "fmt"

//Flyweight используется для управления состоянием объекта с высокой вариативностью. Шаблон
//позволяет нам распределять общие части состояния объекта между несколькими объектами, вместо
//того чтобы каждый объект сохранял их. Переменные данные объекта называются внешним состоянием,
// а остальная часть состояния объекта является внутренней. Внешние данные передаются в методы
//flyweight и никогда не будут сохранены в нем.
//Шаблон Flyweight помогает уменьшить общее использование памяти и накладные расходы на
//инициализацию объекта. Шаблон помогает создавать межклассовые связи и сокращать объем памяти
//до приемлемого уровня.
//
//Объекты Flyweight неизменяемы. Объекты Value являются хорошим примером шаблона flyweight.
//Объекты Flyweight могут создаваться в режиме одного потока, обеспечивая один экземпляр
//для каждого значения. В сценарии параллельного потока создается несколько экземпляров. Это
//основано на критерии равенства объектов с минимальным весом.

// DataTransferObject interface
type DataTransferObject interface {
	getId() string
}

type DataTransferObjectFactory struct {
	pool map[string]DataTransferObject
}

// DataTransferObjectFactory class method getDataTransferObject
func (factory DataTransferObjectFactory) getDataTransferObject(dtoType string) DataTransferObject {
	var dto = factory.pool[dtoType]
	if dto == nil {
		fmt.Println("new DTO of dtoType: " + dtoType)
		switch dtoType {
		case "customer":
			factory.pool[dtoType] = FlyweightCustomer{id: "1"}
		case "employee":
			factory.pool[dtoType] = Employee{id: "2"}
		case "manager":
			factory.pool[dtoType] = Manager{id: "3"}
		case "address":
			factory.pool[dtoType] = Address{id: "4"}
		}
		dto = factory.pool[dtoType]
	}
	return dto
}

// Customer struct
type FlyweightCustomer struct {
	id   string //sequence generator
	name string
	ssn  string
}

// Customer class method getId
func (customer FlyweightCustomer) getId() string {
	//fmt.Println("getting customer Id")
	return customer.id
}

// Employee struct
type Employee struct {
	id   string
	name string
}

// Employee class method getId
func (employee Employee) getId() string {
	return employee.id
}

// Manager struct
type Manager struct {
	id   string
	name string
	dept string
}

// Manager class method getId
func (manager Manager) getId() string {
	return manager.id
}

// Address struct
type Address struct {
	id          string
	streetLine1 string
	streetLine2 string
	state       string
	city        string
}

// Address class method getId
func (address Address) getId() string {
	return address.id
}

func FlyweightExample() {
	var factory = DataTransferObjectFactory{make(map[string]DataTransferObject)}
	var customer DataTransferObject = factory.getDataTransferObject("customer")
	fmt.Println("Customer ", customer.getId())
	var employee DataTransferObject = factory.getDataTransferObject("employee")
	fmt.Println("Employee ", employee.getId())
	var manager DataTransferObject = factory.getDataTransferObject("manager")
	fmt.Println("Manager", manager.getId())
	var address DataTransferObject = factory.getDataTransferObject("address")
	fmt.Println("Address", address.getId())
}
