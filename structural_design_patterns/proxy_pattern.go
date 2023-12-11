package structuraldesignpatterns

import "fmt"

//	Шаблон Прокси (Proxy)
//
//Шаблон прокси перенаправляет к реальному объекту и действует как интерфейс для других. Шаблон прокси
//управляет доступом к объекту и предоставляет дополнительные функциональные возможности. Дополнительные
//функциональные возможности могут быть связаны с аутентификацией, авторизацией и предоставлением прав доступа
//к объекту, чувствительному к ресурсам. Реальный объект не нужно изменять, предоставляя
//дополнительную логику. Удаленные, интеллектуальные, виртуальные прокси-серверы и прокси-серверы защиты - это сценарии, в которых применяется этот
//шаблон. Он также используется для предоставления альтернативы расширению функциональности с помощью
//наследование и композиция объектов. Прокси-объект также называется суррогатом, дескриптором
//или оболочкой.

//В данном примере шаблон прокси используется для выполнения дополнительной проверки входящих данных

// IRealObject interface
type IRealObject interface {
	performAction(int, int) (int, error)
}

// RealObject struct
type RealObject struct{}

// RealObject class method performAction
func (realObject *RealObject) performAction(a, b int) (int, error) {
	return a / b, nil
}

// VirtualProxy struct
type VirtualProxy struct {
	realObject *RealObject
}

func NewVirtualProxy() *VirtualProxy {
	return &VirtualProxy{}
}

// VirtualProxy class method performAction
func (virtualProxy *VirtualProxy) performAction(a, b int) (int, error) {
	if virtualProxy.realObject == nil {
		virtualProxy.realObject = &RealObject{}
	}

	if a == 0 || b == 0 {
		return 0, fmt.Errorf("it cannot be divided by 0")
	}

	if a < b {
		a, b = b, a
	}

	return virtualProxy.realObject.performAction(a, b)
}

func ProxyExample() {
	var object *VirtualProxy = NewVirtualProxy()
	result, err := object.performAction(10, 2)

	fmt.Println("Result: ", result, "Error: ", err)
}
